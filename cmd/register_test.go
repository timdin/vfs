package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/mock"
)

func TestRegister(t *testing.T) {
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
			name:    "create user happy case",
			args:    []string{"register", "timdin"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().Register("timdin").Return(nil)
			},
		},
		{
			name:    "create user unhappy case",
			args:    []string{"register", "timdin"},
			wantErr: true,
			doMock: func() {
				testStorage.EXPECT().Register("timdin").Return(fmt.Errorf("timdin already exists"))
			},
		},
		{
			name:    "create user invalid command case",
			args:    []string{"register"},
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
