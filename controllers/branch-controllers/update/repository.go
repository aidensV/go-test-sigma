package updateBranch

import (
	model "go-test-sigma/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateBranchRepository(input *model.Branch) (*model.Branch, string)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}
func (r *repository) UpdateBranchRepository(input *model.Branch) (*model.Branch, string) {
	var branches model.Branch
	db := r.db.Model(&branches)
	errorCode := make(chan string, 1)

	branches.ID = input.ID
	checkBranchId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&branches)
	if checkBranchId.RowsAffected < 1 {
		errorCode <- "UPDATE_BRANCH_NOT_FOUND_404"
		return &branches, <-errorCode
	}
	branches.Name = input.Name
	branches.Code = input.Code
	branches.Address = input.Address

	updateBranch := db.Debug().Select("name", "code", "address", "updated_at").Where("id = ?", input.ID).Updates(branches)
	if updateBranch.Error != nil {
		errorCode <- "UPDATE_BRANCH_FAILED_403"
		return &branches, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &branches, <-errorCode
}
