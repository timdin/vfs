package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/mock"
)

func TestDeleteFolder(t *testing.T) {
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
			name:    "delete folder happy case",
			args:    []string{"delete-folder", "timdin", "test-folder"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().DeleteFolder("timdin", "test-folder").Return(nil)
			},
		},
		{
			name:    "delete folder unhappy case",
			args:    []string{"delete-folder", "timdin", "test-folder"},
			wantErr: true,
			doMock: func() {
				testStorage.EXPECT().DeleteFolder("timdin", "test-folder").Return(fmt.Errorf("timdin not exist"))
			},
		},
		{
			name:    "delete folder invalid command case (folder name not provided)",
			args:    []string{"delete-folder", "timdin"},
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
