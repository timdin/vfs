package cmd

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/mock"
)

func TestListFile(t *testing.T) {
	ctl := gomock.NewController(t)
	testStorage := mock.NewMockStorage(ctl)

	tests := []struct {
		name    string
		args    []string
		doMock  func()
		wantErr bool
	}{
		{
			name:    "list file happy case",
			args:    []string{"list-files", "timdin", "test-folder", "--" + constants.SortByCreatedFlag, constants.OrderDescString},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFile("timdin", "test-folder", constants.SortByCreated, constants.OrderDesc).Return(nil, nil)
			},
		},
		{
			name:    "list file happy case - no order",
			args:    []string{"list-files", "timdin", "test-folder", "--" + constants.SortByCreatedFlag},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFile("timdin", "test-folder", constants.SortByCreated, constants.OrderAsc).Return(nil, nil)
			},
		},
		{
			name:    "list file happy case - no sort by field",
			args:    []string{"list-files", "timdin", "test-folder", constants.OrderDescString},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFile("timdin", "test-folder", constants.SortByName, constants.OrderDesc).Return(nil, nil)
			},
		},
		{
			name:    "list file happy case - no sort field and order",
			args:    []string{"list-files", "timdin", "test-folder"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFile("timdin", "test-folder", constants.SortByName, constants.OrderAsc).Return(nil, nil)
			},
		},
		{
			name:    "list file unhappy case - bad order",
			args:    []string{"list-files", "timdin", "test-folder", "not-a-order"},
			wantErr: true,
			doMock: func() {
			},
		},
		{
			name:    "list file unhappy case - bad sort by field",
			args:    []string{"list-files", "timdin", "test-folder", "--not-a-flag"},
			wantErr: true,
			doMock: func() {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock()
			}
			// init root cmd on each case execution to purge the flag variables
			rootCmd := InitCmd(testStorage)
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
