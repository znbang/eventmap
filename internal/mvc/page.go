package mvc

import (
	"errors"

	"connectrpc.com/connect"
	"github.com/znbang/eventmap/gen/errdetails"
	"github.com/znbang/validation"
)

func PageSize(p, s int32) (page, size int) {
	page = 1
	size = 10

	if p > 0 {
		page = int(p)
	}

	if s > 0 {
		size = int(s)
	}

	return page, size
}

func NewValidationError(validate *validation.Validation) error {
	validationError := errdetailsv1.ValidationError{Errors: map[string]string{}}
	for k, v := range validate.Errors {
		validationError.Errors[k] = v
	}
	err := connect.NewError(connect.CodeInvalidArgument, errors.New("validation failed"))
	if detail, detailErr := connect.NewErrorDetail(&validationError); detailErr == nil {
		err.AddDetail(detail)
	}
	return err
}
