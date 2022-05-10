package resultsBranch

import (
	model "go-test-sigma/models"

	"gorm.io/gorm"
)

type Repository interface {
	ResultsBranchRepository() (*[]model.Branch, string)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultsBranchRepository() (*[]model.Branch, string) {
	var branches []model.Branch
	db := r.db.Model(&branches)
	errorCode := make(chan string, 1)

	resultBranchs := db.Debug().Select("*").Find(&branches)
	if resultBranchs.Error != nil {
		errorCode <- "RESULTS_BRANCH_NOT_FOUND_404"
		return &branches, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &branches, <-errorCode
}
