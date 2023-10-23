package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/mock"
)

func TestCreateFile(t *testing.T) {
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
			name:    "create file happy case with no description",
			args:    []string{"create-file", "timdin", "test-folder", "test-file"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().CreateFile("timdin", "test-folder", "test-file", "").Return(nil)
			},
		},
		{
			name:    "create file happy case with description",
			args:    []string{"create-file", "timdin", "test-folder", "test-file", "this is my test file"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().CreateFile("timdin", "test-folder", "test-file", "this is my test file").Return(nil)
			},
		},
		{
			name:    "create file unhappy case with description",
			args:    []string{"create-file", "timdin", "test-folder", "test-file", "this is my test file"},
			wantErr: true,
			doMock: func() {
				testStorage.EXPECT().CreateFile("timdin", "test-folder", "test-file", "this is my test file").Return(fmt.Errorf("timdin not exist"))
			},
		},
		{
			name:    "create file invalid command case (folder name not provided)",
			args:    []string{"create-file", "timdin", "test-file"},
			wantErr: true,
		},
		{
			name:    "create file invalid command case (file name not provided)",
			args:    []string{"create-file", "timdin", "test-folder"},
			wantErr: true,
		},
		{
			name:    "create file invalid command case (nothing provided)",
			args:    []string{"create-file", "timdin"},
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
