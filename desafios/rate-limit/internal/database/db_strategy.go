package database

import (
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/dto"
)

type DbStrategy interface {
	Get(id string) *dto.RateLimit
	Save(id string, limit *dto.RateLimit)
}
