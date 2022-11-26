package persistence

import (
	"golang-grpc-sqlboiler-mysql/pkg/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUserRepository_GetUser(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    *models.MUser
		wantErr error
	}{
		{
			name:    "ok",
			userID:  1,
			want:    &models.MUser{ID: 1, Name: "user1", Age: 11},
			wantErr: nil,
		},
		{
			name:    "ok",
			userID:  2,
			want:    &models.MUser{ID: 2, Name: "user2", Age: 22},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.User.GetUser(tt.userID)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

		})

	}

}
