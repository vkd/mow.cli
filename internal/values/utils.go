package values

import (
	"flag"
)

// IsBool checks if a given value is a bool value, i.e. implements the BoolValued interface
func IsBool(v flag.Value) bool {
	if bf, ok := v.(BoolValued); ok {
		return bf.IsBoolFlag()
	}

	return false
}

func DefaultValue(v flag.Value) string {
	if dv, ok := v.(DefaultValued); ok {
		if dv.IsDefault() {
			return ""
		}
	}
	return v.String()
}
