package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/mock"
)

func TestCreateFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	testStorage := mock.NewMockStorage(ctl)
	rootCmd := InitCmd(testStorage)

	tests := []struct {
		name    string
		args    []string
		doMock  func()
		wantErr bool
	}{
		{
			name:    "create folder happy case with no description",
			args:    []string{"create-folder", "timdin", "test-folder"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().CreateFolder("timdin", "test-folder", "").Return(nil)
			},
		},
		{
			name:    "create folder happy case with description",
			args:    []string{"create-folder", "timdin", "test-folder", "test description"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().CreateFolder("timdin", "test-folder", "test description").Return(nil)
			},
		},
		{
			name:    "create folder unhappy case",
			args:    []string{"create-folder", "timdin", "test-folder", "test description"},
			wantErr: true,
			doMock: func() {
				testStorage.EXPECT().CreateFolder("timdin", "test-folder", "test description").Return(fmt.Errorf("timdin not exist"))
			},
		},
		{
			name:    "create user invalid command case (folder name not provided)",
			args:    []string{"create-folder", "timdin"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock()
			}
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
