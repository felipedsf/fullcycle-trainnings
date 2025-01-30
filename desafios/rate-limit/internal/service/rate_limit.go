package service

import (
	"errors"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/database"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/dto"
	"time"
)

type RateLimitService struct {
	db database.DbStrategy
}

func NewService(db database.DbStrategy) *RateLimitService {
	return &RateLimitService{db: db}
}

func (r *RateLimitService) ProcessRequest(token string, ip string) error {
	var id string
	limit := r.db.Get(token)
	if limit == nil {
		limit = r.db.Get(ip)
		if limit == nil {
			limit = dto.NewRateLimit(token, ip)
			r.db.Save(limit.Id, limit)
			return nil
		} else {
			id = ip
		}
	} else {
		id = token
	}
	limit, err := r.attempt(limit)
	r.db.Save(id, limit)
	if err != nil {
		return err
	}
	return nil
}

func (r *RateLimitService) attempt(l *dto.RateLimit) (*dto.RateLimit, error) {
	if (l.BlockedUntil != nil && l.BlockedUntil.Before(time.Now())) || (l.BlockedUntil == nil && l.Until.Before(time.Now())) {
		l = r.update(l)
		return l, nil
	}

	if l.Attempts >= l.MaxAttempts {
		blockedTime := time.Now().Add(time.Duration(l.BlockInterval) * time.Second)
		l.BlockedUntil = &blockedTime
		l.Until = blockedTime
		return l, errors.New("rate limit exceeded")
	}
	l.Attempts++
	return l, nil
}

func (r *RateLimitService) update(l *dto.RateLimit) *dto.RateLimit {
	l.Until = time.Now().Add(time.Duration(l.Interval) * time.Second)
	l.BlockedUntil = nil
	l.Attempts = 1
	return l
}
