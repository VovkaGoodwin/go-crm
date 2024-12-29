package http_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) healthCheck(c *gin.Context) {
	r := h.hc.Handle()
	h.respondSuccess(c, http.StatusOK, r.Result)
}
