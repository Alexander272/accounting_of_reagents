package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type NotesRepo struct {
	db *sqlx.DB
}

func NewNotesRepo(db *sqlx.DB) *NotesRepo {
	return &NotesRepo{
		db: db,
	}
}

type Notes interface {
	Create(context.Context, *models.NoteDTO) (string, error)
	Update(context.Context, *models.NoteDTO) error
	Delete(context.Context, *models.DeleteNoteDTO) error
}

func (r *NotesRepo) Create(ctx context.Context, dto *models.NoteDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, reagent_id, comment, note) VALUES (:id, :reagent_id, :comment, :note)`, NotesTable)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *NotesRepo) Update(ctx context.Context, dto *models.NoteDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET comment=:comment, note=:note WHERE id=:id`, NotesTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *NotesRepo) Delete(ctx context.Context, dto *models.DeleteNoteDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, NotesTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
