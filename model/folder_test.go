package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/timdin/vfs/constants"
	"gorm.io/gorm"
)

func TestFolder_String(t *testing.T) {
	now := time.Now()
	nowString := now.Format(constants.TimeFormat)
	testModel := gorm.Model{
		CreatedAt: now,
	}
	type fields struct {
		Model       gorm.Model
		UserID      uint
		Name        string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test with description",
			fields: fields{
				Model:       testModel,
				UserID:      1,
				Name:        "test",
				Description: "test",
			},
			want: fmt.Sprintf("%v\t%v\t%v", "test", "test", nowString),
		},
		{
			name: "test with no description",
			fields: fields{
				Model:       testModel,
				UserID:      1,
				Name:        "test",
				Description: "",
			},
			want: fmt.Sprintf("%v\t%v\t%v", "test", "N/A", nowString),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Folder{
				Model:       tt.fields.Model,
				UserID:      tt.fields.UserID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
			}
			if got := fmt.Sprint(f); got != tt.want {
				t.Errorf("Folder.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
