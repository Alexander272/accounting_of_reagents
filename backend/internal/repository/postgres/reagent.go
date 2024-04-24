package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type ReagentRepo struct {
	db *sqlx.DB
}

func NewReagentRepo(db *sqlx.DB) *ReagentRepo {
	return &ReagentRepo{
		db: db,
	}
}

type Columns int

const (
	Id Columns = iota
	Type
	Name
	Uname
	Document
	Purity
	DateOfManufacture
	Consignment
	Manufacturer
	ShelfLife
	Closet
	Shelf
	//TODO дописать остальные поля
)

// String returns the string representation of the Subject
func (s Columns) String() string {
	//TODO дописать остальные поля
	return [...]string{"id", "type", "name", "uname", "document", "purity", "dateOfManufacture", "consignment", "manufacturer", "shelfLife", "closet", "shelf"}[s]
}

type Reagent interface{}

func (r *ReagentRepo) Get(ctx context.Context, req *models.Params) ([]*Reagent, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *ReagentRepo) GetById(ctx context.Context, id string) (*models.Reagent, error) {
	// columns := []string{
	// 	Id.String(), Type.String(), Name.String(), Uname.String(), Document.String(), Purity.String(), DateOfManufacture.String(),
	// 	Consignment.String(), Manufacturer.String(), ShelfLife.String(), Closet.String(), Shelf.String(),
	// }

	return nil, fmt.Errorf("not implemented")
}
