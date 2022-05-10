package handlerResultBranch

import (
	resultBranch "go-test-sigma/controllers/branch-controllers/result"
	util "go-test-sigma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service resultBranch.Service
}

func NewhandlerResultBranch(service resultBranch.Service) *handler {
	return &handler{service: service}
}
func (h *handler) ResultBranchHandler(ctx *gin.Context) {
	var input resultBranch.InputResultBranch
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
	resultBranch, errResultBranch := h.service.ResultBranchService(&input)
	switch errResultBranch {
	case "RESULT_BRANCH_NOT_FOUND_404":
		util.APIResponse(ctx, "Branch data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
	default:
		util.APIResponse(ctx, "result branch data successfully", http.StatusOK, http.MethodGet, resultBranch)
	}
}
