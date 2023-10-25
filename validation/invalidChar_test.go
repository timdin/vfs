package validation

import "testing"

func TestInvalidCharacterValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "happy case",
			args:    []string{"123", "abc", "a_b_c", "a-b-c"},
			wantErr: false,
		},
		{
			name:    "bad case 1",
			args:    []string{":"},
			wantErr: true,
		},
		{
			name:    "bad case 2",
			args:    []string{"$$"},
			wantErr: true,
		},
		{
			name:    "bad case 3",
			args:    []string{"^^^"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InvalidCharacterValidation(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("InvalidCharacterValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
