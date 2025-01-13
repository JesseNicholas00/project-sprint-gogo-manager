package optional

import (
	"github.com/JesseNicholas00/GogoManager/types/optional"
	"reflect"
)

func ValidateOptionalString(field reflect.Value) interface{} {
	value, ok := field.Interface().(optional.OptionalStr)

	if ok {
		// If field is not defined, return value.V, omitnil will handle validation & skip the field
		// If field is defined but value null, return nil (invalid value), error occured
		// If field is defined and has value, return the value

		switch {
		case !value.Defined:
			return value.V
		case value.V == nil:
			return nil
		default:
			return value.V
		}
	}

	return nil
}
