package updateBranch

import "time"

type InputUpdateBranch struct {
	ID       string `validate:"required,uuid"`
	Name     string `json:"name" validate:"required,lowercase"`
	Code     string `json:"code" validate:"required"`
	Address  string `json:"address"`
	UpdateAt time.Time
}
