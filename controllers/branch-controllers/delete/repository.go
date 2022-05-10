package deleteBranch

import (
	model "go-test-sigma/models"

	"gorm.io/gorm"
)

type Repository interface {
	DeleteBranchRepository(input *model.Branch) (*model.Branch, string)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteBranchRepository(input *model.Branch) (*model.Branch, string) {
	var branches model.Branch
	db := r.db.Model(&branches)
	errorCode := make(chan string, 1)

	checkBranch := db.Debug().Select("*").Where("id = ?", input.ID).Find(&branches)
	if checkBranch.RowsAffected < 1 {
		errorCode <- "DELETE_BRANCH_NOT_FOUND_404"
		return &branches, <-errorCode
	}
	deleteBranch := db.Debug().Select("*").Where("id = ?", input.ID).Find(&branches).Delete(&branches)
	if deleteBranch.Error != nil {
		errorCode <- "DELETE_BRANCH_FAILED_403"
		return &branches, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &branches, <-errorCode
}
