package mvc

import (
	"errors"
	"maps"

	"connectrpc.com/connect"
	errdetailsv1 "github.com/znbang/eventmap/internal/gen/errdetails"
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
	maps.Copy(validationError.Errors, validate.Errors)
	err := connect.NewError(connect.CodeInvalidArgument, errors.New("validation failed"))
	if detail, detailErr := connect.NewErrorDetail(&validationError); detailErr == nil {
		err.AddDetail(detail)
	}
	return err
}
