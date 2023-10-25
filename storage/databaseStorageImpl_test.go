package storage

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/timdin/vfs/model"
)

func TestDBImpl_Register(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()

	testUserName := "user"
	expectedUser := &model.User{
		Name: testUserName,
	}
	if err := db.Register(testUserName); err != nil {
		t.Error(err)
	}
	actualUser := &model.User{}
	if err := db.lookUpUser(testUserName, actualUser); err != nil {
		t.Error(err)
	}

	if compareStructIgnoreEmptyValues(*expectedUser, *actualUser) != true {
		t.Error(cmp.Diff(expectedUser, actualUser))
	}

	if err := db.lookUpUser("not-a-name", actualUser); err == nil {
		t.Error("should return error when looking up non-existent user")
	}
}

// since we could not predict the value of automatically generated fields, i.e. ID
// provide a helper function to compare two structs ignoring the empty fields in the expected param
func compareStructIgnoreEmptyValues(expected, actual interface{}) bool {
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
