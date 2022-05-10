package handlerCreateBranch

import (
	createBranch "go-test-sigma/controllers/branch-controllers/create"
	util "go-test-sigma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service createBranch.Service
}

func NewHandlerCreateBranch(service createBranch.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateBranchHandler(ctx *gin.Context) {
	var input createBranch.InputCreateBranch
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Name",
				Message: "name is requred on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Name",
				Message: "name must be using lowercase",
			},
			{
				Tag:     "required",
				Field:   "Code",
				Message: "code is requred on body",
			},
		},
	}
	errResponse, errCount := util.GoValidator(&input, config.Options)
	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errCreateBranch := h.service.CreateBranchService(&input)
	switch errCreateBranch {
	case "CREATE_BRANCH_CONFLICT_409":
		util.APIResponse(ctx, "Code Branch already exist", http.StatusConflict, http.MethodPost, nil)
		return
	case "CREATE_BRANCH_FAILED_403":
		util.APIResponse(ctx, "Create new Branch failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		util.APIResponse(ctx, "Create new Branch successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
