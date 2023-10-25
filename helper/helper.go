package helper

import "reflect"

// since we could not predict the value of automatically generated fields, i.e. ID
// provide a helper function to compare two structs ignoring the empty fields in the expected param
// true for identical, false for different
func CompareStructIgnoreEmptyValues(expected, actual interface{}) bool {
	// Use reflection to get the type of the model
	modelType := reflect.TypeOf(expected)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// Check if the field is unexported (lowercase)
		if field.PkgPath == "" {
			fieldName := field.Name
			fieldValue1 := reflect.ValueOf(expected).FieldByName(fieldName)
			fieldValue2 := reflect.ValueOf(actual).FieldByName(fieldName)

			// Check if the field value is not the zero value (default value)
			if !reflect.DeepEqual(fieldValue1.Interface(), reflect.Zero(fieldValue1.Type()).Interface()) {
				// Compare the field values
				if !reflect.DeepEqual(fieldValue1.Interface(), fieldValue2.Interface()) {
					return false
				}
			}
		}
	}

	return true
}

// compare two slices of structs ignoring the empty fields in the expected param
// true for identical, false for different
func CompareStructSliceIgnoreEmptyValues(expected, actual interface{}) bool {
	valExpected := reflect.Indirect(reflect.ValueOf(expected))
	valActual := reflect.Indirect(reflect.ValueOf(actual))
	if valExpected.Len() != valActual.Len() {
		return false
	}

	for i := 0; i < valExpected.Len(); i++ {
		if !CompareStructIgnoreEmptyValues(valExpected.Index(i), valActual.Index(i)) {
			return false
		}
	}

	return true
}
