package cmd

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/mock"
)

func TestRenameFolder(t *testing.T) {
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
			name:    "rename folder happy case",
			args:    []string{"rename-folder", "timdin", "test-folder", "test-folder2"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().RenameFolder("timdin", "test-folder", "test-folder2").Return(nil)
			},
		},
		{
			name:    "rename folder unhappy case - new name not provided",
			args:    []string{"rename-folder", "timdin", "test-folder"},
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
