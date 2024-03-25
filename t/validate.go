package t

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

var (
	Validate = &validate{}
)

func ValidateStruct(obj any) error {
	if Validate == nil {
		return nil
	}
	return Validate.ValidateStruct(obj)
}

type validate struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *validate) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}
func (v *validate) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}
func (v *validate) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(ValidationErrors, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

type ValidationErrors []error

// Error concatenates all error elements in SliceValidationError into a single string separated by \n.
func (err ValidationErrors) Error() string {
	n := len(err)
	switch n {
	case 0:
		return ""
	default:
		var b strings.Builder
		if err[0] != nil {
			fmt.Fprintf(&b, "[%d]: %s", 0, err[0].Error())
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				if err[i] != nil {
					b.WriteString("\n")
					fmt.Fprintf(&b, "[%d]: %s", i, err[i].Error())
				}
			}
		}
		return b.String()
	}
}
