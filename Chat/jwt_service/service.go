package jwtservice

import "github.com/google/uuid"

type JWTService interface {
	ValiadteToken(token string) (uuid.UUID, error)
}
