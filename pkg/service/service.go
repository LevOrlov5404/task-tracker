package service

import (
	"context"
	"github.com/LevOrlov5404/task-tracker/models"
	"github.com/LevOrlov5404/task-tracker/pkg/repository"
)

type (
	User interface {
		CreateUser(ctx context.Context, user models.UserToCreate) (int64, error)
		GetUserByEmailPassword(ctx context.Context, email, password string) (models.UserToGet, error)
		GetUserByID(ctx context.Context, id int64) (models.UserToGet, error)
		UpdateUser(ctx context.Context, id int64, user models.UserToCreate) error
		GetAllUsers(ctx context.Context) ([]models.UserToGet, error)
		DeleteUser(ctx context.Context, id int64) error
		GenerateToken(ctx context.Context, email, password string) (string, error)
		ParseToken(token string) (int64, error)
	}

	ImportanceStatus interface {
		Create(ctx context.Context, status models.StatusToCreate) (int64, error)
	}

	ProgressStatus interface {
		Create(ctx context.Context, status models.StatusToCreate) (int64, error)
	}

	Project interface {
		CreateProject(ctx context.Context, project models.ProjectToCreate) (int64, error)
		GetProjectByID(ctx context.Context, id int64) (models.Project, error)
		UpdateProject(ctx context.Context, id int64, project models.ProjectToUpdate) error
		GetAllProjects(ctx context.Context) ([]models.Project, error)
		DeleteProject(ctx context.Context, id int64) error
	}

	Task interface {
		CreateTaskToProject(ctx context.Context, projectID int64, task models.TaskToCreate) (int64, error)
		GetTaskByID(ctx context.Context, id int64) (models.Task, error)
		UpdateTask(ctx context.Context, id int64, task models.TaskToUpdate) error
		GetAllTasksToProject(ctx context.Context, id int64) ([]models.Task, error)
		DeleteTask(ctx context.Context, id int64) error
	}

	Subtask interface {
	}

	Service struct {
		User
		ImportanceStatus
		ProgressStatus
		Project
		Task
		Subtask
	}
)

func NewService(repo *repository.Repository, salt, signingKey string) *Service {
	return &Service{
		User:             NewUserService(repo.User, salt, signingKey),
		ImportanceStatus: NewImportanceStatusService(repo.ImportanceStatus),
		ProgressStatus:   NewProgressStatusService(repo.ProgressStatus),
		Project:          NewProjectService(repo.Project),
		Task:             NewTaskService(repo.Task),
	}
}
