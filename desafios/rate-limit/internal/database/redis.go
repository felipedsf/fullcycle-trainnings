package database

import (
	"encoding/json"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/config"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/dto"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
}

func GetRedisDb() *Redis {
	c, err := redis.Dial("tcp", config.AppConfig.Redis)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Do("PING")
	if err != nil {
		log.Fatal(err)
	}

	return &Redis{}
}

func (r *Redis) getConn() redis.Conn {
	conn, err := redis.Dial("tcp", config.AppConfig.Redis)
	if err != nil {
		return nil
	}
	return conn
}

func (r *Redis) Get(id string) *dto.RateLimit {
	var rl *dto.RateLimit
	conn := r.getConn()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", id))
	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, &rl)
	if err != nil {
		return nil
	}
	return rl
}

func (r *Redis) Save(id string, limit *dto.RateLimit) {
	bytes, err := json.Marshal(limit)
	if err != nil {
		return
	}

	conn := r.getConn()
	defer conn.Close()

	_, err = conn.Do("SET", id, bytes)
	if err != nil {
		return
	}
}
