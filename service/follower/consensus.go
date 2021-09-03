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

package follower

import (
	"fmt"

	"github.com/dgraph-io/badger/v2"
	"github.com/optakt/flow-dps/models/dps"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/storage/badger/operation"
)

// Consensus is a wrapper around the database that the Flow consensus follower populates. It is used to
// expose the current height and block ID of the consensus follower's last finalized block.
type Consensus struct {
	log zerolog.Logger

	db        *badger.DB
	execution *Execution
	data      map[uint64]conItem
}

type conItem struct {
	header     *flow.Header
	guarantees []*flow.CollectionGuarantee
	seals      []*flow.Seal
}

// NewConsensus returns a new Consensus instance.
func NewConsensus(log zerolog.Logger, db *badger.DB, execution *Execution) *Consensus {
	f := Consensus{
		log:       log,
		db:        db,
		execution: execution,
		data:      make(map[uint64]conItem),
	}

	return &f
}

// Header returns the header for the given height, if available.
func (c *Consensus) Header(height uint64) (*flow.Header, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return it.header, nil
}

// Guarantees returns the collection guarantees for the given height, if available.
func (c *Consensus) Guarantees(height uint64) ([]*flow.CollectionGuarantee, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return it.guarantees, nil
}

// Seals returns the block seals for the given height, if available.
func (c *Consensus) Seals(height uint64) ([]*flow.Seal, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return it.seals, nil
}

// Commit returns the state commitment for the given height, if available.
func (c *Consensus) Commit(height uint64) (flow.StateCommitment, error) {
	it, ok := c.data[height]
	if !ok {
		return flow.DummyStateCommitment, dps.ErrUnavailable
	}
	commit, ok := c.execution.Commit(it.header.ID())
	if !ok {
		return flow.DummyStateCommitment, dps.ErrUnavailable
	}
	return commit, nil
}

func (c *Consensus) Collections(height uint64) ([]*flow.LightCollection, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	collections, ok := c.execution.Collections(it.header.ID())
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return collections, nil
}

func (c *Consensus) Transactions(height uint64) ([]*flow.TransactionBody, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	transactions, ok := c.execution.Transactions(it.header.ID())
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return transactions, nil
}

func (c *Consensus) Results(height uint64) ([]*flow.TransactionResult, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	results, ok := c.execution.Results(it.header.ID())
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return results, nil
}

func (c *Consensus) Events(height uint64) ([]flow.Event, error) {
	it, ok := c.data[height]
	if !ok {
		return nil, dps.ErrUnavailable
	}
	events, ok := c.execution.Events(it.header.ID())
	if !ok {
		return nil, dps.ErrUnavailable
	}
	return events, nil
}

// OnBlockFinalized is a callback that is used to update the state of the Consensus.
func (c *Consensus) OnBlockFinalized(finalID flow.Identifier) {

	var header flow.Header
	var guarantees []*flow.CollectionGuarantee
	var seals []*flow.Seal
	err := c.db.View(func(tx *badger.Txn) error {

		err := operation.RetrieveHeader(finalID, &header)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve header: %w", err)
		}

		var collIDs []flow.Identifier
		err = operation.LookupPayloadGuarantees(finalID, &collIDs)(tx)
		if err != nil {
			return fmt.Errorf("could not look up guarantees: %w", err)
		}

		guarantees = make([]*flow.CollectionGuarantee, 0, len(collIDs))
		for _, collID := range collIDs {
			var guarantee flow.CollectionGuarantee
			err = operation.RetrieveGuarantee(collID, &guarantee)(tx)
			if err != nil {
				return fmt.Errorf("could not retrieve guarantee (collection: %x): %w", collID, err)
			}
			guarantees = append(guarantees, &guarantee)
		}

		var sealIDs []flow.Identifier
		err = operation.LookupPayloadSeals(finalID, &sealIDs)(tx)
		if err != nil {
			return fmt.Errorf("could not look up seals: %w", err)
		}

		seals = make([]*flow.Seal, 0, len(sealIDs))
		for _, sealID := range sealIDs {
			var seal flow.Seal
			err = operation.RetrieveSeal(sealID, &seal)(tx)
			if err != nil {
				return fmt.Errorf("could not retrieve seal: %w", err)
			}
			seals = append(seals, &seal)
		}

		return nil
	})
	if err != nil {
		c.log.Error().Err(err).Msg("could not load consensus data")
		return
	}

	it := conItem{
		header:     &header,
		guarantees: guarantees,
		seals:      seals,
	}
	c.data[header.Height] = it
}