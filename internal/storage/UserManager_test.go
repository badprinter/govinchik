package storage

import (
	"github.com/badprinter/govinchik/internal/models"
	"reflect"
	"testing"
)

func TestUserManager_getUserByID(t *testing.T) {
	type args struct {
		id  int64
		url *string
	}
	tests := []struct {
		name string
		args args
		want *models.User
	}{
		// TODO: Add test cases.
		{
			name: "Valid user ID",
			args: args{id: 1, url: new(string)},
			want: &models.User{
				Id:         1,
				Name:       "",
				TelegramId: 0,
				IsDeleted:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManager{}
			if got := u.getUserByID(tt.args.id, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
