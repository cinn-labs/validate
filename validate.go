package validate

import "errors"

type Validations map[string][]string

func (v *Validations) Add(field string, message string) error {
	if (*v)[field] == nil {
		(*v)[field] = []string{}
	}
	(*v)[field] = append((*v)[field], message)
	return errors.New("Invalid Record")
}

// TODO: Handle SQL response errors
