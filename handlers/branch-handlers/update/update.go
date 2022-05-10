package handlerUpdateBranch

import (
	updateBranch "go-test-sigma/controllers/branch-controllers/update"
	util "go-test-sigma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service updateBranch.Service
}

func NewHandlerUpdateBranch(service updateBranch.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateBranchHandler(ctx *gin.Context) {
	var input updateBranch.InputUpdateBranch
	input.ID = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

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
			{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Name",
				Message: "name must be using lowercase",
			},
			{
				Tag:     "required",
				Field:   "Code",
				Message: "code is required on body",
			},
		},
	}
	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}
	_, errUpdateBranch := h.service.UpdateBranchService(&input)
	switch errUpdateBranch {
	case "UPDATE_BRANCH_NOT_FOUND_404":
		util.APIResponse(ctx, "branch data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)
	case "UPDATE_BRANCH_FAILED_403":
		util.APIResponse(ctx, "update branch data failed", http.StatusForbidden, http.MethodPost, nil)
	default:
		util.APIResponse(ctx, "update branch successfully", http.StatusOK, http.MethodPost, nil)

	}
}
