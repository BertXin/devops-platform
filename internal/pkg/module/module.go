package module

import (
	"context"
	"devops-platform/internal/pkg/security"
	"devops-platform/pkg/types"
)

type Module struct {
	ID             types.Long `json:"id" swaggertype:"string" gorm:"primaryKey"`
	CreatedAt      types.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy      User       `json:"created_by" gorm:"embedded;embeddedPrefix:created_by_" swaggertype:"string"`
	LastModifiedAt types.Time `json:"last_modified_at" gorm:"autoUpdateTime"`
	LastModifiedBy User       `json:"last_modified_by" gorm:"embedded;embeddedPrefix:last_modified_by_" swaggertype:"string"`
}

type User struct {
	ID   types.Long `json:"id"`
	Name string     `json:"name"`
}

func (u *User) from(user security.User) {
	if user == nil {
		u.ID = 0
		u.Name = "系统"
	} else {
		u.ID = user.GetID()
		u.Name = user.GetName()
	}
}

func SystemUser() User {
	return User{ID: 0, Name: "系统"}
}

func (m *Module) AuditCreated(ctx context.Context) {
	m.CreatedBy.from(security.CurrentUser(ctx))
	m.LastModifiedBy.from(security.CurrentUser(ctx))
}

func (m *Module) AuditModified(ctx context.Context) {
	m.LastModifiedBy.from(security.CurrentUser(ctx))
}

type Operation struct {
	Operator   User       `json:"operator" gorm:"embedded;embeddedPrefix:operator_" swaggertype:"string"`
	OperatedAt types.Time `json:"operated_at" gorm:"autoCreateTime"`
}

type CreateOnlyModule struct {
	ID        types.Long `json:"id" swaggertype:"string" gorm:"primaryKey"`
	CreatedAt types.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy User       `json:"created_by" gorm:"embedded;embeddedPrefix:created_by_" swaggertype:"string"`
}

func (o *Operation) OperatingRecord(ctx context.Context) {
	o.Operator.from(security.CurrentUser(ctx))
}

func (m *CreateOnlyModule) AuditCreated(ctx context.Context) {
	m.CreatedBy.from(security.CurrentUser(ctx))
}

type DeleteStatus int

const (
	Normal  DeleteStatus = 0 // 正常
	Deleted DeleteStatus = 1 // 已删除
)

var deleteStatusMap = map[DeleteStatus]string{
	Normal:  "正常",
	Deleted: "已删除",
}
