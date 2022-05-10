package resultBranch

import model "go-test-sigma/models"

type Service interface {
	ResultBranchService(input *InputResultBranch) (*model.Branch, string)
}
type service struct {
	repository Repository
}

func NewServiceResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultBranchService(input *InputResultBranch) (*model.Branch, string) {
	branches := model.Branch{
		ID: input.ID,
	}
	resultBranch, errCreateBranch := s.repository.ResultBranchRepository(&branches)
	return resultBranch, errCreateBranch
}
