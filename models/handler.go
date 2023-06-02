package models

import (
	"database/sql/driver"
	"errors"

	"github.com/thomasfady/xsstower/config"
	"gorm.io/gorm"
)

type PermisionType string

const (
	READ  PermisionType = "READ"
	WRITE PermisionType = "WRITE"
	OWNER PermisionType = "OWNER"
)

func (t PermisionType) Value() (driver.Value, error) {
	switch t {
	case READ, WRITE, OWNER:
		return string(t), nil
	}
	return nil, errors.New("Invalid permission type value")
}

func (t *PermisionType) Scan(value interface{}) error {
	var pt PermisionType
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.(string)
	if !ok {
		return errors.New("Invalid data for permission type")
	}
	pt = PermisionType(st)

	switch pt {
	case READ, WRITE, OWNER:
		*t = pt
		return nil
	}
	return errors.New("Invalid permission type value")
}

type HandlerRbac struct {
	gorm.Model
	UserID     int
	HandlerID  int
	User       User
	Permission PermisionType
	Handler    Handler
}
type Handler struct {
	gorm.Model
	Path           string `gorm:"size:255;not null"`
	Domain         string `gorm:"default:*"`
	OwnerID        int
	Owner          User `json:"-"`
	Screenshot     bool
	Dom            bool
	CollectedPages []string      `gorm:"serializer:json"`
	Members        []HandlerRbac `gorm:"foreignKey:HandlerID"`
}

func GetHandleRbacForUser(userId int) (rbacs []HandlerRbac) {
	DB.Preload("User").Preload("Handler").Find(&rbacs, "user_id = ?", userId)
	return
}

func GetViewableHandlers(userId int) (handlers []Handler) {
	rbacs := GetHandleRbacForUser(userId)
	var handlers_id []int
	for _, rbac := range rbacs {
		handlers_id = append(handlers_id, rbac.HandlerID)
	}
	DB.Preload("Members.User").Find(&handlers, "id IN ?", handlers_id)
	return
}

func GetHandlerById(id interface{}) (handler Handler) {
	DB.Preload("Members.User").First(&handler, id)
	return
}

func (h *Handler) GetRoleForUser(user_id int) string {
	for _, member := range h.Members {
		if member.UserID == user_id {
			return string(member.Permission)
		}
	}
	return ""
}

func (h *Handler) CanView(user_id int) bool {
	return len(h.GetRoleForUser(user_id)) > 0
}

func (h *Handler) CanUpdate(user_id int) bool {
	role := h.GetRoleForUser(user_id)
	return role == string(WRITE) || role == string(OWNER)
}

func (h *Handler) CanDelete(user_id int) bool {
	return h.GetRoleForUser(user_id) == string(OWNER)
}

func (h *Handler) Delete() {
	DB.Delete(h)
}

func (h *Handler) Save() {
	DB.Save(h)
}

func HandleIsFree(path string, domain string) (bool, string) {
	var handler Handler
	DB.First(&handler, "path = ? and domain = ?", path, domain)
	if handler.ID != 0 {
		return false, "Handler already taken"
	}

	DB.First(&handler, "path = ? and domain = '*'", path)
	if handler.ID != 0 {
		return false, "Handler already taken (This path with a wildcard domain)"
	}

	DB.First(&handler, "path = '*' and domain = ?", domain)
	if handler.ID != 0 {
		return false, "Handler already taken (This domain with a wildcard path)"
	}

	if domain == "*" {
		if config.GetConfig("handlers.wildcard.domain") != true {
			return false, "Wildcard domain creation disabled"
		}
		DB.First(&handler, "path = ?", path)
		if handler.ID != 0 {
			return false, "Handler already exists for this path. Wildcard creation is not possible"
		}
	}
	if path == "*" {
		if config.GetConfig("handlers.wildcard.path") != true {
			return false, "Wildcard path creation disabled"
		}
		DB.First(&handler, "domain = ?", domain)
		if handler.ID != 0 {
			return false, "Handler already exists for this domain. Wildcard creation is not possible"
		}
	}
	return true, ""
}
