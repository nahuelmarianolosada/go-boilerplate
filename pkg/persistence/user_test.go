package persistence

import (
	"reflect"
	"testing"

	"github.com/nahuelmarianolosada/go-boilerplate/pkg/models"
)

func TestGetByUsername_mustFail(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{"First test to find into an empty db", args{"test"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestCreateUser_mustNotFail(t *testing.T) {
	type args struct {
		u models.User
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"First test to create a basic user validating id into response", args{models.User{Username: "test", Password: "test"}}, 1, false},
		{"Second test to create a basic user validating id into response", args{models.User{Username: "test1", Password: "test1"}}, 2, false},
		{"Second test to create a basic user validating id into response", args{models.User{Username: "test", Password: "test"}}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

