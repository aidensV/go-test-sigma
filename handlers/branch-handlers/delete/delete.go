package handlerDeleteBranch

import (
	deleteBranch "go-test-sigma/controllers/branch-controllers/delete"
	util "go-test-sigma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service deleteBranch.Service
}

func NewhandlerDeleteBranch(service deleteBranch.Service) *handler {
	return &handler{service: service}
}
func (h *handler) DeleteBranchHandler(ctx *gin.Context) {
	var input deleteBranch.InputDeleteBranch
	input.ID = ctx.Param("id")
	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
			{
				Tag:     "uuid",
				Field:   "ID",
				Message: "params must be uuid format",
			},
		},
	}
	errResponse, errCount := util.GoValidator(&input, config.Options)
	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}
	deleteBranch, errDeleteBranch := h.service.DeleteBranchService(&input)
	switch errDeleteBranch {
	case "DELETE_BRANCH_NOT_FOUND_404":
		util.APIResponse(ctx, "Branch data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
	default:
		util.APIResponse(ctx, "delete branch data successfully", http.StatusOK, http.MethodGet, deleteBranch)
	}
}
