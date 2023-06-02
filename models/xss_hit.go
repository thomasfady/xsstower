package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type XssHit struct {
	gorm.Model
	Ip             string
	Referer        string
	UserAgent      string
	Origin         string
	CorrelationKey string `gorm:"not null"`
	PublicId       string
	Cookies        string
	LocalStorage   string
	SessionStorage string
	Url            string
	Dom            string
	Screenshot     string
	OwnerID        int
	Owner          User `json:"-"`
	HandlerID      int
	Handler        Handler `json:"-"`
	CollectedPages []CollectedPage
}

func (hit *XssHit) CanView(user_id int) bool {
	for _, member := range hit.Handler.Members {
		if member.UserID == user_id {
			return true
		}
	}
	return false
}

func (hit *XssHit) CanUpdate(user_id int) bool {
	for _, member := range hit.Handler.Members {
		if member.UserID == user_id && (member.Permission == WRITE || member.Permission == OWNER) {
			return true
		}
	}
	return false
}

func (hit *XssHit) CanDelete(user_id int) bool {
	for _, member := range hit.Handler.Members {
		if member.UserID == user_id && member.Permission == OWNER {
			return true
		}
	}
	return false
}

func GetHitByCorrelationKey(key string) XssHit {
	var hit XssHit
	DB.Preload("CollectedPages").Preload("Handler.Members").First(&hit, "correlation_key = ?", key)
	return hit
}

func GetViewableHitPaginated(user_id int, pagination *Pagination) (hits []XssHit) {
	rbacs := GetHandleRbacForUser(user_id)
	var handlers_id []int
	for _, rbac := range rbacs {
		handlers_id = append(handlers_id, rbac.HandlerID)
	}
	DB.Scopes(Paginate(hits, pagination, DB)).Preload("Handler").Find(&hits, "handler_id IN ?", handlers_id)
	return
}

func GetHitBySharingKey(key string) XssHit {
	var hit XssHit
	DB.Preload("CollectedPages").Preload("Handler.Members").First(&hit, "public_id = ?", key)
	return hit
}

func (hit *XssHit) EnableSharing() {
	hit.PublicId = uuid.New().String()
	DB.Save(hit)
}

func (hit *XssHit) DisableSharing() {
	hit.PublicId = ""
	DB.Save(hit)
}

func (hit *XssHit) Delete() {
	DB.Delete(hit)
}
