package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type DBManagement struct {
	AnalyticalMydbWriteConn  *sqlx.DB
	AnalyticalMydbReadConn   *sqlx.DB
	AnalyticalRedisWriteConn *redis.Client
	AnalyticalRedisReadConn  *redis.Client
	CmotsWriteMydbConn       *sqlx.DB
	CmotsReadMydbConn        *sqlx.DB
}
