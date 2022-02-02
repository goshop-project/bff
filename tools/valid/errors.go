package valid

import (
	valid "github.com/asaskevich/govalidator"
)

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error = valid.Error

// Errors is an array of multiple errors and conforms to the error interface.
type Errors struct {
	valid.Errors
}

// AsError returns an array of errors, or nil if none has been appended
func (es *Errors) AsError() error {
	if es != nil && len(es.Errors) > 0 {
		return es.Errors
	}
	return nil
}

// AppendError appends an error if declared
func (es *Errors) AppendError(err error) {
	if err != nil {
		es.Errors = append(es.Errors, err)
	}
}

// AppendFieldError appends a validator error
func (es *Errors) AppendFieldError(name, validator string, err error, path ...string) {
	if len(name) > 0 {
		es.AppendError(valid.Error{
			Name:                     name,
			Err:                      err,
			CustomErrorMessageExists: err != nil,

			Validator: validator,
			Path:      path,
		})
	}
}
