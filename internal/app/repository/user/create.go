package user

import (
	"context"

	"github.com/viettrung2103/bookmark-management-lesson/internal/app/model"
)

func (r *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
