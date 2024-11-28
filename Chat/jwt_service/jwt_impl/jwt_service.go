package jwtimpl

import (
	"github.com/google/uuid"
	jwtservice "github.com/nilspolek/DevOps/Chat/jwt_service"
)

type svc struct {
}

func New() jwtservice.JWTService {
	return &svc{}
}

func (s *svc) ValiadteToken(token string) (uuid.UUID, error) {
	return uuid.MustParse(token), nil
}
