package entities

import "github.com/google/uuid"

type Code struct {
	Url  string `json:"url"`
	Code string `json:"code"`
}

func NewCode(url string) *Code {
	return &Code{
		Url:  url,
		Code: "",
	}
}

func (c *Code) GenerateCode() {
	id := uuid.New()
	code := id.String()[0:6]
	c.Code = code
}
