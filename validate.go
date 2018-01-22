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

func (v *Validations) AddMany(field string, msgs []string) error {

	if (*v)[field] == nil {
		(*v)[field] = []string{}
	}
	(*v)[field] = append((*v)[field], msgs...)
	return errors.New("Invalid Record")
}

func (v *Validations) MergeSubFieldValidations(v2 Validations, p string) {
	for field, msgs := range v2 {
		v.AddMany(p+"."+field, msgs)
	}
}

// TODO: Handle SQL response errors
