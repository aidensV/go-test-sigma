package deleteBranch

type InputDeleteBranch struct {
	ID string `validate:"required,uuid"`
}
