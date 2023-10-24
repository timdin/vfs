package cmd

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/mock"
)

func TestListFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	testStorage := mock.NewMockStorage(ctl)

	tests := []struct {
		name    string
		args    []string
		doMock  func()
		wantErr bool
	}{
		{
			name:    "list folder happy case",
			args:    []string{"list-folders", "timdin", "--" + constants.SortByCreatedFlag, constants.OrderDescString},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFolder("timdin", constants.SortByCreated, constants.OrderDesc).Return(nil, nil)
			},
		},
		{
			name:    "list folder happy case w/o order",
			args:    []string{"list-folders", "timdin", "--" + constants.SortByCreatedFlag},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFolder("timdin", constants.SortByCreated, constants.OrderAsc).Return(nil, nil)
			},
		},
		{
			name:    "list folder happy case w/o sort field",
			args:    []string{"list-folders", "timdin", constants.OrderDescString},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFolder("timdin", constants.SortByName, constants.OrderDesc).Return(nil, nil)
			},
		},
		{
			name:    "list folder happy case w/o order or field",
			args:    []string{"list-folders", "timdin"},
			wantErr: false,
			doMock: func() {
				testStorage.EXPECT().ListFolder("timdin", constants.SortByName, constants.OrderAsc).Return(nil, nil)
			},
		},
		{
			name:    "list folder unhappy case - invalid sort by field",
			args:    []string{"list-folders", "timdin", "--not-a-flag"},
			wantErr: true,
		},
		{
			name:    "list folder unhappy case - invalid order field",
			args:    []string{"list-folders", "timdin", "not-a-order"},
			wantErr: true,
		},
		{
			name:    "list folder unhappy case - invalid order and sort by",
			args:    []string{"list-folders", "timdin", "--not-a-flag", "not-a-order"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock()
			}
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
