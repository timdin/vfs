package storage

import (
	"testing"
)

func TestDBImpl_Register(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := InitTestDB()
			if err := db.Register(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
			TeardownTestDB()
		})
	}
}
