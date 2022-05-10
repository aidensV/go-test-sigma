package createBranch

import (
	model "go-test-sigma/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBranchRepository(input *model.Branch) (*model.Branch, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateBranchRepository(input *model.Branch) (*model.Branch, string) {
	var branches model.Branch
	db := r.db.Model(&branches)
	errorCode := make(chan string, 1)

	checkBranchExist := db.Debug().Select("*").Where("code = ?", input.Code).Find(&branches)
	if checkBranchExist.RowsAffected > 0 {
		errorCode <- "CREATE_BRANCH_CONFLICT_409"
		return &branches, <-errorCode
	}

	branches.Name = input.Name
	branches.Code = input.Code
	branches.Address = input.Address

	addNewBranch := db.Debug().Create(&branches)
	db.Commit()
	if addNewBranch.Error != nil {
		errorCode <- "CREATE_BRANCH_FAILED_403"
		return &branches, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &branches, <-errorCode
}
