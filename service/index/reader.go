package index

import (
	"fmt"

	"github.com/rs/zerolog"

	"golang.org/x/sync/errgroup"

	"github.com/dgraph-io/badger/v2"
	"github.com/onflow/flow-go/model/flow"

	"github.com/onflow/flow-archive/models/archive"
)

// ConcurrentPathReadLimit sets the number of concurrent paths in the Values lookup
const ConcurrentPathReadLimit int = 1024

// Reader implements the `index.Reader` interface on top of the DPS server's
// Badger database index.
type Reader struct {
	log zerolog.Logger
	db  *badger.DB

	// Old badger-based index. Eventually all methods from here would migrate to lib2.
	lib archive.ReadLibrary
	// New pebble-based index.  Would likely consist of multiple separate pebble databases underneath.
	lib2 archive.ReadLibrary2
}

// NewReader creates a new index reader, using the given database as the
// underlying state repository. It is recommended to provide a read-only Badger
// database.
func NewReader(
	log zerolog.Logger,
	db *badger.DB,
	lib archive.ReadLibrary,
	lib2 archive.ReadLibrary2,
) *Reader {

	r := Reader{
		log: log.With().Str("component", "index_reader").Logger(),
		db:  db,

		lib:  lib,
		lib2: lib2,
	}

	return &r
}

// First returns the height of the first finalized block that was indexed.
func (r *Reader) First() (uint64, error) {
	var height uint64
	err := r.db.View(r.lib.RetrieveFirst(&height))
	return height, err
}

// Last returns the height of the last finalized block that was indexed.
func (r *Reader) Last() (uint64, error) {
	var height uint64
	err := r.db.View(r.lib.RetrieveLast(&height))
	return height, err
}

// LatestRegisterHeight returns the latest height for which all registers are indexed
func (r *Reader) LatestRegisterHeight() (uint64, error) {
	var height uint64
	err := r.db.View(r.lib.RetrieveLatestRegisterHeight(&height))
	return height, err
}

// HeightForBlock returns the height for the given block identifier.
func (r *Reader) HeightForBlock(blockID flow.Identifier) (uint64, error) {
	var height uint64
	err := r.db.View(r.lib.LookupHeightForBlock(blockID, &height))
	return height, err
}

// Commit returns the commitment of the execution state as it was after the
// execution of the finalized block at the given height.
func (r *Reader) Commit(height uint64) (flow.StateCommitment, error) {
	var commit flow.StateCommitment
	err := r.db.View(r.lib.RetrieveCommit(height, &commit))
	return commit, err
}

// Header returns the header for the finalized block at the given height.
func (r *Reader) Header(height uint64) (*flow.Header, error) {
	var header flow.Header
	err := r.db.View(r.lib.RetrieveHeader(height, &header))
	return &header, err
}

func (r *Reader) getValue(height uint64, reg flow.RegisterID) ([]byte, error) {
	payload, err := r.lib2.GetPayload(height, reg)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve payload (register: %x): %w", reg, err)
	}
	return payload, nil
}

func (r *Reader) getValues(height uint64, regs flow.RegisterIDs) ([][]byte, error) {
	values := make([]flow.RegisterValue, len(regs))
	g := errgroup.Group{}
	g.SetLimit(ConcurrentPathReadLimit)
	for i, reg := range regs {
		i, reg := i, reg
		g.Go(func() error {
			payload, err := r.getValue(height, reg)
			if err != nil {
				return fmt.Errorf("failed to getValue for %d/%d: %w", i, len(regs), err)
			}

			values[i] = payload
			return nil
		})
	}
	return values, g.Wait()
}

// Values returns the Ledger values of the execution state at the given paths
// as they were after the execution of the finalized block at the given height.
// For compatibility with existing Flow execution node code, a path that is not
// found within the indexed execution state returns a nil value without error.
func (r *Reader) Values(height uint64, regs flow.RegisterIDs) ([]flow.RegisterValue, error) {
	first, err := r.First()
	if err != nil {
		return nil, fmt.Errorf("could not check first height: %w", err)
	}
	last, err := r.Last()
	if err != nil {
		return nil, fmt.Errorf("could not check last height: %w", err)
	}

	// TODO(rbtz): on average we execute 3 queries for each register fetch:
	//   first height, last height, and the register value itself.
	// We should cache the first and last heights in memory to avoid the overhead.
	if height < first || height > last {
		return nil, fmt.Errorf("invalid height (given: %d, first: %d, last: %d)", height, first, last)
	}

	// fast-path for fetching a single register
	if len(regs) == 1 {
		value, err := r.getValue(height, regs[0])
		if err != nil {
			return nil, fmt.Errorf("could not retrieve value: %w", err)
		}
		return []flow.RegisterValue{value}, nil
	}

	return r.getValues(height, regs)
}

// Collection returns the collection with the given ID.
func (r *Reader) Collection(collID flow.Identifier) (*flow.LightCollection, error) {
	var collection flow.LightCollection
	err := r.db.View(r.lib.RetrieveCollection(collID, &collection))
	return &collection, err
}

// CollectionsByHeight returns the collection IDs at the given height.
func (r *Reader) CollectionsByHeight(height uint64) ([]flow.Identifier, error) {
	var collIDs []flow.Identifier
	err := r.db.View(r.lib.LookupCollectionsForHeight(height, &collIDs))
	return collIDs, err
}

// Guarantee returns the guarantee with the given collection ID.
func (r *Reader) Guarantee(collID flow.Identifier) (*flow.CollectionGuarantee, error) {
	var collection flow.CollectionGuarantee
	err := r.db.View(r.lib.RetrieveGuarantee(collID, &collection))
	return &collection, err
}

// Transaction returns the transaction with the given ID.
func (r *Reader) Transaction(txID flow.Identifier) (*flow.TransactionBody, error) {
	var transaction flow.TransactionBody
	err := r.db.View(r.lib.RetrieveTransaction(txID, &transaction))
	return &transaction, err
}

// HeightForTransaction returns the height of the block within which the given
// transaction identifier is.
func (r *Reader) HeightForTransaction(txID flow.Identifier) (uint64, error) {
	var height uint64
	err := r.db.View(r.lib.LookupHeightForTransaction(txID, &height))
	return height, err
}

// TransactionsByHeight returns the transaction IDs within the block with the given ID.
func (r *Reader) TransactionsByHeight(height uint64) ([]flow.Identifier, error) {
	var txIDs []flow.Identifier
	err := r.db.View(r.lib.LookupTransactionsForHeight(height, &txIDs))
	return txIDs, err
}

// Result returns the transaction result for the given transaction ID.
func (r *Reader) Result(txID flow.Identifier) (*flow.TransactionResult, error) {
	var result flow.TransactionResult
	err := r.db.View(r.lib.RetrieveResult(txID, &result))
	return &result, err
}

// Events returns the events of all transactions that were part of the
// finalized block at the given height. It can optionally filter them by event
// type; if no event types are given, all events are returned.
func (r *Reader) Events(height uint64, types ...flow.EventType) ([]flow.Event, error) {
	first, err := r.First()
	if err != nil {
		return nil, fmt.Errorf("could not check first height: %w", err)
	}
	last, err := r.Last()
	if err != nil {
		return nil, fmt.Errorf("could not check last height: %w", err)
	}
	if height < first || height > last {
		return nil, fmt.Errorf("invalid height (given: %d, first: %d, last: %d)", height, first, last)
	}

	var events []flow.Event
	err = r.db.View(r.lib.RetrieveEvents(height, types, &events))
	if err != nil {
		return nil, fmt.Errorf("could not retrieve events: %w", err)
	}

	return events, nil
}

// Seal returns the seal with the given ID.
func (r *Reader) Seal(sealID flow.Identifier) (*flow.Seal, error) {
	var seal flow.Seal
	err := r.db.View(r.lib.RetrieveSeal(sealID, &seal))
	return &seal, err
}

// SealsByHeight returns all of the seals that were part of the finalized block at the given height.
func (r *Reader) SealsByHeight(height uint64) ([]flow.Identifier, error) {
	var sealIDs []flow.Identifier
	err := r.db.View(r.lib.LookupSealsForHeight(height, &sealIDs))
	return sealIDs, err
}
