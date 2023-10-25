package storage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/helper"
	"github.com/timdin/vfs/model"
)

// test for internal utility methods
// lookUpUser
// lookUpFolder
// lookUpFile

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

// look up folder includes lookup unhappy cases
func TestDBImpl_lookUpFolder(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "testuser"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)
	// set up test folder
	testFolderName := "testfolder"
	db.CreateFolder(testUserName, testFolderName, "")

	tests := []struct {
		name       string
		testUser   *model.User
		folderName string
		wantErr    bool
	}{
		{
			name:       "test look up folder happy case",
			testUser:   testUser,
			folderName: "testfolder",
			wantErr:    false,
		},
		{
			name:       "test look up folder happy case",
			testUser:   testUser,
			folderName: "testfolder2",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualFolder := &model.Folder{}
			if err := db.lookUpFolder(tt.testUser, tt.folderName, actualFolder); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.lookUpUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// tests for external production methods
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

func TestDBImpl_CreateFile(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)
	// set up test folder
	testFolderName := "folder"
	db.CreateFolder(testUserName, testFolderName, "")
	testFolder := &model.Folder{}
	db.lookUpFolder(testUser, testFolderName, testFolder)

	tests := []struct {
		name            string
		userName        string
		folderName      string
		fileName        string
		description     string
		expectedFile    *model.File
		wantErr         bool
		additionalSetup func()
	}{
		{
			name:        "test create file happy case",
			fileName:    "file1",
			folderName:  testFolderName,
			userName:    testUserName,
			description: "description1",
			expectedFile: &model.File{
				Name:        "file1",
				UserID:      testUser.ID,
				Description: "description1",
			},
		},
		{
			name:        "test create file happy case - no description",
			fileName:    "file2",
			folderName:  testFolderName,
			userName:    testUserName,
			description: "",
			expectedFile: &model.File{
				Name:        "file2",
				UserID:      testUser.ID,
				Description: "",
			},
		},
		{
			name:            "test create file unhappy case - name conflict",
			fileName:        "conflict",
			folderName:      testFolderName,
			userName:        testUserName,
			description:     "",
			wantErr:         true,
			additionalSetup: func() { db.CreateFile(testUserName, testFolderName, "conflict", "") },
			// negative test does not compare
			expectedFile: &model.File{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.additionalSetup != nil {
				tt.additionalSetup()
			}
			if err := db.CreateFile(tt.userName, tt.folderName, tt.fileName, tt.description); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			actualFile := &model.File{}
			if err := db.lookUpFile(testUser, testFolder, tt.fileName, actualFile); err != nil {
				t.Error(err)
			}
			if helper.CompareStructIgnoreEmptyValues(*tt.expectedFile, *actualFile) != true {
				t.Error(cmp.Diff(tt.expectedFile, actualFile))
			}
		})
	}
}

func TestDBImpl_DeleteFolder(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	// set up test folder
	testFolderName := "folder"
	db.CreateFolder(testUserName, testFolderName, "")

	tests := []struct {
		name       string
		userName   string
		folderName string
		wantErr    bool
	}{
		{
			name:       "delete folder happy case",
			userName:   testUserName,
			folderName: testFolderName,
		},
		{
			name:       "delete folder unhappy case - folder not exist",
			userName:   testUserName,
			folderName: "notexist",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := db.DeleteFolder(tt.userName, tt.folderName); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.DeleteFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBImpl_DeleteFile(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	// set up test folder
	testFolderName := "folder"
	db.CreateFolder(testUserName, testFolderName, "")
	// set up test file
	testFileName := "file"
	db.CreateFile(testUserName, testFolderName, testFileName, "")

	tests := []struct {
		name       string
		userName   string
		folderName string
		fileName   string
		wantErr    bool
	}{
		{
			name:       "delete file happy case",
			userName:   testUserName,
			folderName: testFolderName,
			fileName:   testFileName,
		},
		{
			name:       "delete file unhappy case - file not exist",
			userName:   testUserName,
			folderName: testFolderName,
			fileName:   "notexist",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := db.DeleteFile(tt.userName, tt.folderName, tt.fileName); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.DeleteFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBImpl_RenameFolder(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)
	// set up test folder
	testFolderName := "folder"
	db.CreateFolder(testUserName, testFolderName, "description1")
	testFolderName2 := "folder2"
	db.CreateFolder(testUserName, testFolderName2, "description1")

	tests := []struct {
		name            string
		userName        string
		originalName    string
		newName         string
		expectedFolder  *model.Folder
		wantErr         bool
		additionalSetup func()
	}{
		{
			name:         "test rename folder happy case",
			userName:     testUserName,
			originalName: testFolderName,
			newName:      "new",
			expectedFolder: &model.Folder{
				Name:        "new",
				UserID:      testUser.ID,
				Description: "description1",
			},
		},
		{
			name:           "test rename folder unhappy case - original folder not exist",
			userName:       testUserName,
			originalName:   "not exists",
			newName:        "new2",
			wantErr:        true,
			expectedFolder: &model.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.additionalSetup != nil {
				tt.additionalSetup()
			}
			if err := db.RenameFolder(tt.userName, tt.originalName, tt.newName); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.RenameFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			actualFolder := &model.Folder{}
			if err := db.lookUpFolder(testUser, tt.newName, actualFolder); (err != nil) != tt.wantErr {
				t.Error(err)
			}
			if helper.CompareStructIgnoreEmptyValues(*tt.expectedFolder, *actualFolder) != true {
				t.Error(cmp.Diff(tt.expectedFolder, actualFolder))
			}
		})
	}
}

func TestDBImpl_ListFile(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)
	// set up test folder
	testFolderName1 := "test1"
	db.CreateFolder(testUserName, testFolderName1, "description1")
	actualFolder1 := &model.Folder{}
	db.lookUpFolder(testUser, testFolderName1, actualFolder1)
	testFolderName2 := "test2"
	db.CreateFolder(testUserName, testFolderName2, "description1")
	// set up test file
	testFileName1 := "test2"
	db.CreateFile(testUserName, testFolderName1, testFileName1, "")
	testFileName2 := "test1"
	db.CreateFile(testUserName, testFolderName1, testFileName2, "")
	testFile1 := &model.File{}
	db.lookUpFile(testUser, actualFolder1, testFileName1, testFile1)
	testFile2 := &model.File{}
	db.lookUpFile(testUser, actualFolder1, testFileName2, testFile2)

	tests := []struct {
		name            string
		userName        string
		folderName      string
		expectedFiles   []*model.File
		sortBy          constants.SortByField
		order           constants.Order
		wantErr         bool
		additionalSetup func()
	}{
		{
			name:          "test list file happy case - order by created ascending",
			userName:      testUserName,
			folderName:    testFolderName1,
			sortBy:        constants.SortByCreated,
			order:         constants.OrderAsc,
			expectedFiles: []*model.File{testFile2, testFile1},
		},
		{
			name:          "test list file happy case - order by created descending",
			userName:      testUserName,
			folderName:    testFolderName1,
			sortBy:        constants.SortByCreated,
			order:         constants.OrderDesc,
			expectedFiles: []*model.File{testFile1, testFile2},
		},
		{
			name:          "test list file happy case - order by name ascending",
			userName:      testUserName,
			folderName:    testFolderName1,
			sortBy:        constants.SortByName,
			order:         constants.OrderAsc,
			expectedFiles: []*model.File{testFile2, testFile1},
		},
		{
			name:          "test list file happy case - order by name descending",
			userName:      testUserName,
			folderName:    testFolderName1,
			sortBy:        constants.SortByName,
			order:         constants.OrderDesc,
			expectedFiles: []*model.File{testFile1, testFile2},
		},
		{
			name:          "test list file happy case - no file found",
			userName:      testUserName,
			folderName:    testFolderName2,
			sortBy:        constants.SortByName,
			order:         constants.OrderDesc,
			expectedFiles: []*model.File{},
		},
		{
			name:          "test list file unhappy case - not exist folder",
			userName:      testUserName,
			folderName:    "not exist folder",
			sortBy:        constants.SortByName,
			order:         constants.OrderDesc,
			expectedFiles: []*model.File{},
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := db.ListFile(tt.userName, tt.folderName, tt.sortBy, tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.ListFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if helper.CompareStructSliceIgnoreEmptyValues(tt.expectedFiles, res) != true {
				t.Error(cmp.Diff(tt.expectedFiles, res))
			}
		})
	}
}

func TestDBImpl_ListFolder(t *testing.T) {
	defer TeardownTestDB()
	db := InitTestDB()
	// set up test user
	testUserName := "user"
	db.Register(testUserName)
	testUser := &model.User{}
	db.lookUpUser(testUserName, testUser)
	testNoFolderUserName := "no-folder-user"
	db.Register(testNoFolderUserName)
	// set up test folder
	testFolderName1 := "test2"
	db.CreateFolder(testUserName, testFolderName1, "description1")
	actualFolder1 := &model.Folder{}
	db.lookUpFolder(testUser, testFolderName1, actualFolder1)
	testFolderName2 := "test1"
	db.CreateFolder(testUserName, testFolderName2, "description1")
	actualFolder2 := &model.Folder{}
	db.lookUpFolder(testUser, testFolderName2, actualFolder2)

	tests := []struct {
		name            string
		userName        string
		expectedFolder  []*model.Folder
		sortBy          constants.SortByField
		order           constants.Order
		wantErr         bool
		additionalSetup func()
	}{
		{
			name:           "test list folder happy case - order by created ascending",
			userName:       testUserName,
			sortBy:         constants.SortByCreated,
			order:          constants.OrderAsc,
			expectedFolder: []*model.Folder{actualFolder1, actualFolder2},
		},
		{
			name:           "test list folder happy case - order by created descending",
			userName:       testUserName,
			sortBy:         constants.SortByCreated,
			order:          constants.OrderDesc,
			expectedFolder: []*model.Folder{actualFolder2, actualFolder1},
		},
		{
			name:           "test list file happy case - order by name ascending",
			userName:       testUserName,
			sortBy:         constants.SortByName,
			order:          constants.OrderAsc,
			expectedFolder: []*model.Folder{actualFolder2, actualFolder1},
		},
		{
			name:           "test list folder happy case - order by name descending",
			userName:       testUserName,
			sortBy:         constants.SortByName,
			order:          constants.OrderDesc,
			expectedFolder: []*model.Folder{actualFolder1, actualFolder2},
		},
		{
			name:           "test list folder happy case - user with no folder",
			userName:       testNoFolderUserName,
			sortBy:         constants.SortByName,
			order:          constants.OrderDesc,
			expectedFolder: []*model.Folder{},
		},
		{
			name:           "test list folder unhappy case - not exist user",
			userName:       "not-exist",
			sortBy:         constants.SortByName,
			order:          constants.OrderDesc,
			expectedFolder: []*model.Folder{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := db.ListFolder(tt.userName, tt.sortBy, tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.ListFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if helper.CompareStructSliceIgnoreEmptyValues(tt.expectedFolder, res) != true {
				t.Error(cmp.Diff(tt.expectedFolder, res))
			}
		})
	}
}
