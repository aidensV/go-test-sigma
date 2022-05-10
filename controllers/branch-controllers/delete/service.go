package deleteBranch

import model "go-test-sigma/models"

type Service interface {
	DeleteBranchService(input *InputDeleteBranch) (*model.Branch, string)
}
type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteBranchService(input *InputDeleteBranch) (*model.Branch, string) {
	branches := model.Branch{
		ID: input.ID,
	}
	deleteBranch, errCreateBranch := s.repository.DeleteBranchRepository(&branches)
	return deleteBranch, errCreateBranch
}
