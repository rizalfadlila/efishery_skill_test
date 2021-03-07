package helper

import (
	"context"

	"github.com/fetch_app/constants"
	"github.com/fetch_app/exceptions"
)

// UserContext contains user data parsed from context
type UserContext struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
}

// ParseUserContext parse user context data
func ParseUserContext(ctx context.Context) (userContext UserContext, err error) {
	var exist bool

	userContext.Name, exist = ctx.Value(constants.KeyName).(string)
	if !exist {
		return UserContext{}, exceptions.ErrMissingName
	}

	userContext.Email, exist = ctx.Value(constants.KeyEmail).(string)
	if !exist {
		return UserContext{}, exceptions.ErrMissingEmail
	}

	userContext.Phone, exist = ctx.Value(constants.KeyPhone).(string)
	if !exist {
		return UserContext{}, exceptions.ErrMissingPhone
	}

	userContext.Role, exist = ctx.Value(constants.KeyRole).(string)
	if !exist {
		return UserContext{}, exceptions.ErrMissingRole
	}

	userContext.Timestamp, exist = ctx.Value(constants.KeyTimestamp).(string)
	if !exist {
		return UserContext{}, exceptions.ErrMissingRole
	}

	return
}
