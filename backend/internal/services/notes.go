package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type NoteService struct {
	repo repository.Notes
}

func NewNoteService(repo repository.Notes) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

type Note interface {
	GetByReagentId(context.Context, string) ([]*models.Note, error)
	Create(context.Context, *models.NoteDTO) (string, error)
	Update(context.Context, *models.NoteDTO) error
	Delete(context.Context, *models.DeleteNoteDTO) error
}

func (s *NoteService) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Note, error) {
	notes, err := s.repo.GetByReagentId(ctx, reagentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes by reagent id. error: %w", err)
	}
	return notes, nil
}

func (s *NoteService) Create(ctx context.Context, dto *models.NoteDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create note. error: %w", err)
	}
	return id, nil
}

func (s *NoteService) Update(ctx context.Context, dto *models.NoteDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update note. error: %w", err)
	}
	return nil
}

func (s *NoteService) Delete(ctx context.Context, dto *models.DeleteNoteDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete note. error: %w", err)
	}
	return nil
}
