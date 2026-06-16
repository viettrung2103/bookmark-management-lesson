package user

import (
	"context"

	"github.com/viettrung2103/bookmark-management-lesson/internal/app/model"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/utils"
)

func (s *userService) CreateUser(ctx context.Context, displayName, username, password, email string) (*model.User, error) {
	hashedPwd := utils.Hastring(password)

	user := &model.User{
		Username:    username,
		Password:    hashedPwd,
		Email:       email,
		DisplayName: displayName,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
