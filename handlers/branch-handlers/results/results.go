package handlerResultsBranch

import (
	resultsBranch "go-test-sigma/controllers/branch-controllers/results"
	util "go-test-sigma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultsBranch.Service
}

func NewhandlerResultBranch(service resultsBranch.Service) *handler {
	return &handler{service: service}
}
func (h *handler) ResultsBranchHandler(ctx *gin.Context) {

	resultBranch, errResultBranch := h.service.ResultsBranchService()
	switch errResultBranch {
	case "RESULTS_BRANCH_NOT_FOUND_404":
		util.APIResponse(ctx, "Branch data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
	default:
		util.APIResponse(ctx, "result branch data successfully", http.StatusOK, http.MethodGet, resultBranch)
	}
}
