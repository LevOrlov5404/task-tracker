package handler

import (
	"github.com/LevOrlov5404/task-tracker/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createProject(c *gin.Context) {
	var project models.ProjectToCreate
	if err := c.BindJSON(&project); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Project.Create(c, project)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllProjects(c *gin.Context) {
	items, err := h.services.Project.GetAll(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getProjectByID(c *gin.Context) {

}

func (h *Handler) updateProject(c *gin.Context) {

}

func (h *Handler) deleteProject(c *gin.Context) {

}
