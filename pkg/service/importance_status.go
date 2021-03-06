package service

import (
	"context"
	"github.com/LevOrlov5404/task-tracker/models"
	"github.com/LevOrlov5404/task-tracker/pkg/repository"
)

type ImportanceStatusService struct {
	repo repository.ImportanceStatus
}

func NewImportanceStatusService(repo repository.ImportanceStatus) *ImportanceStatusService {
	return &ImportanceStatusService{repo: repo}
}

func (s *ImportanceStatusService) Create(ctx context.Context, status models.StatusToCreate) (int64, error) {
	return s.repo.Create(ctx, status)
}

func (s *ImportanceStatusService) GetByID(ctx context.Context, id int64) (models.Status, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ImportanceStatusService) Update(ctx context.Context, id int64, status models.StatusToCreate) error {
	return s.repo.Update(ctx, id, status)
}

func (s *ImportanceStatusService) GetAll(ctx context.Context) ([]models.Status, error) {
	return s.repo.GetAll(ctx)
}

func (s *ImportanceStatusService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
