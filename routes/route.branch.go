package route

import (
	createBranch "go-test-sigma/controllers/branch-controllers/create"
	deleteBranch "go-test-sigma/controllers/branch-controllers/delete"
	resultBranch "go-test-sigma/controllers/branch-controllers/result"
	resultsBranch "go-test-sigma/controllers/branch-controllers/results"

	updateBranch "go-test-sigma/controllers/branch-controllers/update"
	handlerCreateBranch "go-test-sigma/handlers/branch-handlers/create"
	handlerResultBranch "go-test-sigma/handlers/branch-handlers/result"
	handlerResultsBranch "go-test-sigma/handlers/branch-handlers/results"

	handlerDeleteBranch "go-test-sigma/handlers/branch-handlers/delete"
	handlerUpdateBranch "go-test-sigma/handlers/branch-handlers/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBranchRoutes(db *gorm.DB, route *gin.Engine) {
	// index
	resultsBranchRepository := resultsBranch.NewRepositoryResults(db)
	resultsBranchService := resultsBranch.NewServiceResult(resultsBranchRepository)
	resultsBranchHandler := handlerResultsBranch.NewhandlerResultBranch(resultsBranchService)
	// Store
	createBranchRepository := createBranch.NewRepositoryCreate(db)
	createBranchService := createBranch.NewServiceCreate(createBranchRepository)
	createBranchHandler := handlerCreateBranch.NewHandlerCreateBranch(createBranchService)

	// show
	resultBranchRepository := resultBranch.NewRepositoryResult(db)
	resultBranchService := resultBranch.NewServiceResult(resultBranchRepository)
	resultBranchHandler := handlerResultBranch.NewhandlerResultBranch(resultBranchService)
	// Update
	updatestudentRepository := updateBranch.NewRepositoryUpdate(db)
	updateBranchService := updateBranch.NewServiceUpdate(updatestudentRepository)
	updateBranchHandler := handlerUpdateBranch.NewHandlerUpdateBranch(updateBranchService)

	// Delete
	deleteBranchRepository := deleteBranch.NewRepositoryDelete(db)
	deleteBranchService := deleteBranch.NewServiceDelete(deleteBranchRepository)
	deleteBranchHandler := handlerDeleteBranch.NewhandlerDeleteBranch(deleteBranchService)

	groupRoute := route.Group("/api/v1")

	groupRoute.POST("/branch", createBranchHandler.CreateBranchHandler)
	groupRoute.GET("/branch", resultsBranchHandler.ResultsBranchHandler)
	groupRoute.GET("/branch/:id", resultBranchHandler.ResultBranchHandler)
	groupRoute.PUT("/branch/:id", updateBranchHandler.UpdateBranchHandler)
	groupRoute.DELETE("/branch/:id", deleteBranchHandler.DeleteBranchHandler)
}
