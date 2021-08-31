// Copyright 2021 Optakt Labs OÜ
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package execution

import (
	"errors"
	"fmt"
	"math"

	"github.com/dgraph-io/badger/v2"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
	"github.com/optakt/flow-dps/models/dps"
)

const cacheSize = 1024

type blockReader interface {
	Notify() <-chan string
	Read(filename string) ([]byte, error)
}

type Follower struct {
	log zerolog.Logger

	blocks   blockReader
	codec    dps.Codec
	db       *badger.DB
	validate *validator.Validate

	// FIXME: Index by block ID instead of height.
	data    map[uint64]*BlockData
	current *BlockData
	height  uint64
	index   int

	stop chan struct{}
}

func New(log zerolog.Logger, blocks blockReader, codec dps.Codec, db *badger.DB) *Follower {
	f := Follower{
		log: log,

		blocks:   blocks,
		codec:    codec,
		db:       db,
		validate: validator.New(),

		height: math.MaxUint64,
		data:   make(map[uint64]*BlockData, cacheSize),

		stop: make(chan struct{}),
	}

	return &f
}

// Run launches the execution follower.
func (f *Follower) Run() error {
	for {
		select {
		case <-f.stop:
			return nil
		default:
		}

		// If cache is already full, wait for items to get consumed.
		if len(f.data) == cacheSize {
			continue
		}

		// If cache has available space, wait for new block data to be available.
		select {
		case <-f.stop:
			return nil
		case file := <-f.blocks.Notify():
			b, err := f.blocks.Read(file)
			if err != nil {
				return fmt.Errorf("could not read block data: %w", err)
			}

			var data BlockData
			err = f.codec.Unmarshal(b, &data)
			if err != nil {
				return fmt.Errorf("could not decode block data: %w", err)
			}
			err = f.validate.Struct(data)
			if err != nil {
				return fmt.Errorf("invalid block data: %w", err)
			}

			f.data[data.Block.Header.Height] = &data

			// FIXME: Rewrite this in a more elegant way
			// This should only ever happen once, when we get the oldest available block.
			if data.Block.Header.Height < f.height {
				if f.height != math.MaxUint64 {
					// This should never happen, it would mean we received blocks in an
					// incorrect order!
					return fmt.Errorf("fatal discrepancy: received block data in incorrect order")
				}
				f.height = data.Block.Header.Height
			}
		}
	}
}

// Update returns the next trie update, in chronological order. It is also in charge of moving onto the next
func (f *Follower) Update() (*ledger.TrieUpdate, error) {
	// If we reached the end of the trie updates for the current block, it means it should have been indexed
	// successfully. Therefore, we move on to the next block.
	if f.current == nil || f.index == len(f.current.TrieUpdates) {
		err := f.next()
		if err != nil {
			return nil, fmt.Errorf("could not forward execution follower to the next block: %w", err)
		}
	}

	// Copy next update to be returned.
	update := f.current.TrieUpdates[f.index]
	f.index++

	return update, nil
}

func (f *Follower) next() error {
	if len(f.data) == 0 {
		return dps.ErrUnavailable
	}

	// Only increment height and reset index if we are moving from one block to the next,
	// not if this is the first block.
	if f.current != nil {
		f.height++
		f.index = 0
	}

	var exists bool
	f.current, exists = f.data[f.height]
	if !exists {
		return errors.New("fatal discrepancy: execution follower height does not match available block data")
	}

	// Sanity check: Verify that we can find a matching seal and execution result for that block ID.
	blockID := f.current.Block.ID()
	for _, blockData := range f.data {
		for _, seal := range blockData.Block.Payload.Seals {
			// Only look for the seal of the current block.
			if seal.BlockID != blockID {
				continue
			}

			for _, result := range blockData.Block.Payload.Results {
				// Only look for the result of the current block.
				if result.BlockID != blockID {
					continue
				}

				finalState, err := result.FinalStateCommitment()
				if err != nil {
					return fmt.Errorf("could not compute state commitment from execution result: %w", err)
				}

				if seal.FinalState != finalState {
					return errors.New("fatal discrepancy: mismatch between seal and execution result state commitments")
				}
				break
			}
			break
		}
	}

	return nil
}

// Header returns the header for the current block.
func (f *Follower) Header(height uint64) (*flow.Header, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	return f.current.Block.Header, nil
}

// Collections returns the collections for the current block.
func (f *Follower) Collections(height uint64) ([]*flow.LightCollection, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	var colls []*flow.LightCollection
	for _, coll := range f.current.Collections {
		lightColl := coll.Collection().Light()
		colls = append(colls, &lightColl)
	}

	return colls, nil
}

// Guarantees returns the guarantees for the current block.
func (f *Follower) Guarantees(height uint64) ([]*flow.CollectionGuarantee, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	var guars []*flow.CollectionGuarantee
	for _, coll := range f.current.Collections {
		guars = append(guars, coll.Guarantee)
	}

	return guars, nil
}

// Seals returns the seals for the current block.
func (f *Follower) Seals(height uint64) ([]*flow.Seal, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	return f.current.Block.Payload.Seals, nil
}

// Transactions returns the transactions for the current block.
func (f *Follower) Transactions(height uint64) ([]*flow.TransactionBody, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	var transactions []*flow.TransactionBody
	for _, coll := range f.current.Collections {
		transactions = append(transactions, coll.Transactions...)
	}

	return transactions, nil
}

// Results returns the results for the current block.
func (f *Follower) Results(height uint64) ([]*flow.TransactionResult, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	return f.current.TxResults, nil
}

// Events returns the events for the current block.
func (f *Follower) Events(height uint64) ([]flow.Event, error) {
	if f.current == nil {
		return nil, errors.New("no block data available")
	}
	if height != f.height {
		return nil, fmt.Errorf("block data requested for wrong block height (current: %d, requested %d)", f.height, height)
	}

	var events []flow.Event
	for _, e := range f.current.Events {
		e := e
		events = append(events, *e)
	}

	return events, nil
}

func (f *Follower) Stop() {
	close(f.stop)
}