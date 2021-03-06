package models

type (
	TaskToCreate struct {
		Title              string `json:"title" binding:"required"`
		Description        string `json:"description"`
		AssigneeID         int64  `json:"assigneeId" binding:"required"`
		ImportanceStatusID int64  `json:"importanceStatusId" binding:"required"`
		ProgressStatusID   int64  `json:"progressStatusId" binding:"required"`
	}

	TaskToUpdate struct {
		TaskToCreate
		CreationDate string `json:"creationDate" binding:"required"`
	}

	Task struct {
		ID                 int64  `json:"id" db:"id"`
		Title              string `json:"title" db:"title"`
		Description        string `json:"description" db:"description"`
		CreationDate       string `json:"creationDate" db:"creation_date"`
		AssigneeID         int64  `json:"assigneeId" db:"assignee_id"`
		ImportanceStatusID int64  `json:"importanceStatusId" db:"importance_status_id"`
		ProgressStatusID   int64  `json:"progressStatusId" db:"progress_status_id"`
	}

	TaskParams struct {
		ID                 *int64  `json:"id"`
		Title              *string `json:"title"`
		Description        *string `json:"description"`
		CreationDate       *string `json:"creationDate"`
		AssigneeID         *int64  `json:"assigneeId"`
		ImportanceStatusID *int64  `json:"importanceStatusId"`
		ProgressStatusID   *int64  `json:"progressStatusId"`
	}

	TaskWithSubtasks struct {
		Task
		Subtasks []Subtask `json:"subtasks"`
	}

	TaskWithProjectID struct {
		ProjectID int64 `db:"project_id"`
		Task
	}
)

func (taskWithProjectID TaskWithProjectID) ToTaskWithSubtasks() TaskWithSubtasks {
	return TaskWithSubtasks{
		Task: taskWithProjectID.Task,
	}
}
