package postgres

import (
	"database/sql"
	"github.com/fatih/structs"
	"github.com/lann/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/nikchabanyk/blober/internal/data"
)

const (
	blobsTable = "blobs"
)

var (
	ErrBlobsConflict = errors.New("blobs primary key conflict")
	ErrNoWallet      = errors.New("blobs_wallets_fkey violated")
)

type Blobs struct {
	db  *pgdb.DB
	sel squirrel.SelectBuilder
}

func NewBlobs(db *pgdb.DB) *Blobs {
	return &Blobs{
		db, squirrel.Select("*").From(blobsTable),
	}
}

func (q *Blobs) New() data.Blobs {
	return NewBlobs(q.db.Clone())
}

func (q *Blobs) Transaction(fn func(data.Blobs) error) error {
	return q.db.Transaction(
		func() error {
			return fn(q)
		},
	)
}

func (q *Blobs) Create(blob *data.Blob) error {
	err := q.db.Exec(squirrel.Insert(blobsTable).SetMap(structs.Map(blob)))
	if err != nil {
		return errors.Wrap(
			err, "failed to exec statement",
		)
	}

	return nil
}

func (q *Blobs) Get(id string) (*data.Blob, error) {
	var result data.Blob
	stmt := q.sel.Where("id = ?", id)

	err := q.db.Get(&result, stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (q *Blobs) GetAll() ([]*data.Blob, error) {
	var results []*data.Blob
	stmt := q.sel // You can customize the statement here if needed.
	err := q.db.Select(
		&results, stmt,
	)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (q *Blobs) DeleteByID(id string) error {
	err := q.db.Exec(
		squirrel.Delete(blobsTable).Where(squirrel.Eq{"id": id}),
	)
	if err != nil {
		return err
	}
	return nil
}
