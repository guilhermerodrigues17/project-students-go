package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	Cpf    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"`
}

func errParamRequired(param string, typo string) error {
	return fmt.Errorf("the parameter '%s' is required and must be of type %s", param, typo)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.Cpf == "" {
		return errParamRequired("cpf", "string")
	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.Age <= 0 {
		return errParamRequired("age", "integer")
	}
	if s.Active == nil {
		return errParamRequired("active", "bool")
	}

	return nil
}
