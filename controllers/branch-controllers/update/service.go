package updateBranch

import (
	model "go-test-sigma/models"
)

type Service interface {
	UpdateBranchService(input *InputUpdateBranch) (*model.Branch, string)
}
type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateBranchService(input *InputUpdateBranch) (*model.Branch, string) {
	branches := model.Branch{
		ID:      input.ID,
		Name:    input.Name,
		Code:    input.Code,
		Address: input.Address,
	}
	resultUpdateBranch, errUpdateBranch := s.repository.UpdateBranchRepository(&branches)
	return resultUpdateBranch, errUpdateBranch
}
