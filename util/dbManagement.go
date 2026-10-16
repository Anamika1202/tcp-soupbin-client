package util

import (
	"tcp-soupbin-client/constants"
	"tcp-soupbin-client/model"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

var DBManagement = &model.DBManagement{}

func LoadDbConnections(myDbConnnection map[string]*sqlx.DB, redisConnnection map[string]*redis.Client) {
	DBManagement.AnalyticalRedisWriteConn = redisConnnection[constants.REDIS_ANALYTICAL_MASTER]
	DBManagement.AnalyticalRedisReadConn = redisConnnection[constants.REDIS_ANALYTICAL_SLAVE]
	DBManagement.AnalyticalMydbWriteConn = myDbConnnection[constants.MYSQL_ANALYTIC_MASTER]
	DBManagement.AnalyticalMydbWriteConn = myDbConnnection[constants.MYSQL_ANALYTIC_SLAVE]
	DBManagement.CmotsWriteMydbConn = myDbConnnection[constants.MYSQL_CMOTS_MASTER]
	DBManagement.CmotsReadMydbConn = myDbConnnection[constants.MYSQL_CMOTS_SLAVE]

}

func GetMyDbReadConn(department string) *sqlx.DB {
	switch department {
	// case constants.OMS_DEPT_ID:
	// 	return DBManagement.OmsMydbReadConn
	// case constants.ANL_DEPT_ID:
	// 	return DBManagement.AnalyticMydbReadConn
	// case constants.CMOTS_DEPT_ID:
	// 	return DBManagement.AnalyticMydbReadConn
	// case constants.MF_DEPT_ID:
	// 	return DBManagement.MfReadMydbConn
	default:
		return DBManagement.AnalyticalMydbWriteConn
	}
}
func GetRedisWriteConn(department string) *redis.Client {
	switch department {
	//case constants.OMS_DEPT_ID:
	// 	return DBManagement.OmsRedisWriteConn
	// case constants.INMEM_DEP_ID:
	// 	return DBManagement.InMemRedis
	// case constants.BRODCST_DEP_ID:
	//return DBManagement.BrodCastRedis
	default:
		return DBManagement.AnalyticalRedisReadConn
	}
}
