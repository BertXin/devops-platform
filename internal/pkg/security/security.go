package security

import (
	"context"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
)

const (
	currentUser = "current-user"
)

type User interface {
	GetID() types.Long
	GetName() string
	GetRole() enum.SysRole
	GetToken() string
	GetLoginName() string
}

type systemUser struct {
	Name string
}

func (u *systemUser) GetID() types.Long {
	return 0
}

func (u *systemUser) GetName() string {
	return "系统"
}

func (u *systemUser) GetRole() enum.SysRole {
	return enum.SysRoleAdminUser
}

func (u *systemUser) GetToken() string {
	return u.Name
}

func (u *systemUser) GetLoginName() string {
	return u.Name
}

func CurrentUser(ctx context.Context) User {
	user, ok := ctx.Value(currentUser).(User)
	if !ok {
		return nil
	}
	return user
}

func SetCurrentUser(ctx context.Context, user User) context.Context {
	if user == nil {
		return ctx
	}
	return context.WithValue(ctx, currentUser, user)
}

func CurrentIsAdmin(ctx context.Context) bool {
	user := CurrentUser(ctx)

	if user == nil {
		return false
	}
	return user.GetRole() == enum.SysRoleAdminUser
}

func CurrentUserID(ctx context.Context) types.Long {
	user, ok := ctx.Value(currentUser).(User)
	if !ok {
		return 0
	}
	return user.GetID()
}

func SetCurrentUserSystem(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, currentUser, &systemUser{Name: name})
}
