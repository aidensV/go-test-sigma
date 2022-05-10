package resultsBranch

import model "go-test-sigma/models"

type Service interface {
	ResultsBranchService() (*[]model.Branch, string)
}
type service struct {
	repository Repository
}

func NewServiceResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsBranchService() (*[]model.Branch, string) {
	resultBranch, errCreateBranch := s.repository.ResultsBranchRepository()
	return resultBranch, errCreateBranch
}
