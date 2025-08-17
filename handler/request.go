package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type OpeningRequest interface {
	Validate() error
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   uint64 `json:"salary"`
}

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   uint64 `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if c.Salary == 0 {
		return errParamIsRequired("salary", "uint64")
	}
	return nil
}

func (u *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if u.Role != "" || u.Company != "" || u.Location != "" || u.Remote != nil || u.Link != "" || u.Salary > 0 {
		return nil
	}

	// If none of the fields were provided, return falsy
	return fmt.Errorf("at lest one valid field must be provided")
}
