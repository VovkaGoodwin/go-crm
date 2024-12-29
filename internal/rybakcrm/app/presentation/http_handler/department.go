package http_handler

import (
	"crm-backend/pkg/app_errors"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"crm-backend/internal/rybakcrm/app/application/dto"
)

type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *Handler) createDepartment(c *gin.Context) {
	var req CreateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Debug("create department", "error", err)
		h.respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	createDepartmentDto := &dto.CreateDepartmentRequest{
		Name:        req.Name,
		Description: req.Description,
	}

	ctx := c.Request.Context()

	response, err := h.dep.CreateDepartment(ctx, createDepartmentDto)
	h.logger.Debug("create department", "response", response)
	if err != nil {
		h.logger.Debug("create department", "error", err)
		h.respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(c, http.StatusCreated, *response.Department)
}

func (h *Handler) getAllDepartments(c *gin.Context) {
	ctx := c.Request.Context()

	response, err := h.dep.GetAllDepartments(ctx, &dto.GetAllDepartmentsRequest{})
	if err != nil {
		h.logger.Debug("get all departments", "error", err)
		h.respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(c, http.StatusOK, response.Departments)
}

func (h *Handler) getDepartment(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("get department", "error", err)
		h.respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	request := &dto.GetDepartmentByIdRequest{
		DepartmentId: id,
	}

	response, err := h.dep.GetDepartmentById(ctx, request)
	if err != nil {
		h.logger.Debug("get department", "error", err)
		h.respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(c, http.StatusOK, *response.Department)
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) updateDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("update department", "error", err)
		h.respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var req UpdateDepartmentRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		h.logger.Debug("update department binding", "error", err)
		h.respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	request := &dto.UpdateDepartmentRequest{
		DepartmentId: id,
		Name:         req.Name,
		Description:  req.Description,
	}

	response, err := h.dep.UpdateDepartment(ctx, request)

	if err != nil {
		h.logger.Debug("update department", "error", err)

		status := http.StatusInternalServerError

		if errors.Is(err, app_errors.ErrDepartmentWithThisNameExists) {
			status = http.StatusBadRequest
		}

		h.respondError(c, status, err.Error())
		return
	}

	h.respondSuccess(c, http.StatusOK, *response.Department)
}

func (h *Handler) deleteDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("delete department", "error", err)
		h.respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	request := &dto.DeleteDepartmentRequest{
		DepartmentId: id,
	}

	response, err := h.dep.DeleteDepartment(ctx, request)

	if err != nil {
		h.logger.Debug("delete department", "error", err)
		h.respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(c, http.StatusNoContent, response)
}
