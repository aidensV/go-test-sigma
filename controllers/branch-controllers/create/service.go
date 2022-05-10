package createBranch

import (
	model "go-test-sigma/models"
)

type Service interface {
	CreateBranchService(input *InputCreateBranch) (*model.Branch, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateBranchService(input *InputCreateBranch) (*model.Branch, string) {
	branches := model.Branch{
		Name:    input.Name,
		Code:    input.Code,
		Address: input.Address,
	}
	resultCreateBranch, errCreateBranch := s.repository.CreateBranchRepository(&branches)
	return resultCreateBranch, errCreateBranch
}
