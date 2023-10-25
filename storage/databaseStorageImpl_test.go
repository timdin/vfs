package storage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/timdin/vfs/helper"
	"github.com/timdin/vfs/model"
)

func TestDBImpl_Register(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()

	tests := []struct {
		name            string
		userName        string
		additionalSetup func()
		wantErr         bool
	}{
		{
			name:     "test register user happy case",
			userName: "testuser",
			wantErr:  false,
		},
		{
			name:            "test register user unhappy case - name conflict",
			userName:        "conflict",
			additionalSetup: func() { db.Register("conflict") },
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.additionalSetup != nil {
				tt.additionalSetup()
			}
			if err := db.Register(tt.userName); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBImpl_CreateFolder(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)

	tests := []struct {
		name            string
		userName        string
		folderName      string
		description     string
		expectedFolder  *model.Folder
		wantErr         bool
		additionalSetup func()
	}{
		{
			name:        "test create folder happy case",
			folderName:  "folder1",
			userName:    testUserName,
			description: "description1",
			expectedFolder: &model.Folder{
				Name:        "folder1",
				UserID:      testUser.ID,
				Description: "description1",
			},
		},
		{
			name:        "test create folder happy case - no description",
			folderName:  "folder2",
			userName:    testUserName,
			description: "",
			expectedFolder: &model.Folder{
				Name:        "folder2",
				UserID:      testUser.ID,
				Description: "",
			},
		},
		{
			name:            "test create folder unhappy case - name conflict",
			folderName:      "conflict",
			userName:        testUserName,
			description:     "",
			wantErr:         true,
			additionalSetup: func() { db.CreateFolder(testUserName, "conflict", "") },
			// negative test does not compare
			expectedFolder: &model.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.additionalSetup != nil {
				tt.additionalSetup()
			}
			if err := db.CreateFolder(tt.userName, tt.folderName, tt.description); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			actualFolder := &model.Folder{}
			if err := db.lookUpFolder(testUser, tt.folderName, actualFolder); err != nil {
				t.Error(err)
			}
			if helper.CompareStructIgnoreEmptyValues(*tt.expectedFolder, *actualFolder) != true {
				t.Error(cmp.Diff(tt.expectedFolder, actualFolder))
			}
		})
	}

}

// look up user includes lookup unhappy cases
func TestDBImpl_lookUpUser(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	db.Register("testuser")

	tests := []struct {
		name     string
		userName string
		wantErr  bool
	}{
		{
			name:     "test look up user happy case",
			userName: "testuser",
			wantErr:  false,
		},
		{
			name:     "test look up user happy case",
			userName: "testuser2",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualUser := &model.User{}
			if err := db.lookUpUser(tt.userName, actualUser); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.lookUpUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
