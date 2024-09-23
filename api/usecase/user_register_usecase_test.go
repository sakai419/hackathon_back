package usecase

import (
	"api/model"
	"testing"
)

func TestCheckUser(t *testing.T) {
	type args struct {
		data model.UserJsonForHTTPPost
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "data.Name is empty",
			args: args{model.UserJsonForHTTPPost{Name: "", Age: 20}},
			wantErr: true,
		},
		{
			name: "data.Name is too long",
			args: args{model.UserJsonForHTTPPost{Name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Age: 20}},
			wantErr: true,
		},
		{
			name: "data.Age is out of range",
			args: args{model.UserJsonForHTTPPost{Name: "test", Age: 19}},
			wantErr: true,
		},
		{
			name: "data.Age is out of range",
			args: args{model.UserJsonForHTTPPost{Name: "test", Age: 81}},
			wantErr: true,
		},
		{
			name: "data.Age is in range",
			args: args{model.UserJsonForHTTPPost{Name: "test", Age: 20}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CheckUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
