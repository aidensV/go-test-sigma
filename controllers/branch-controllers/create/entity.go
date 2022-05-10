package createBranch

type InputCreateBranch struct {
	Name    string `json:"name" validate:"required,lowercase"`
	Code    string `json:"code" validate:"required"`
	Address string `json:"address"`
}
