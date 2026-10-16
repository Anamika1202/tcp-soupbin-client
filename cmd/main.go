package main

import (
	"flag"
	"fmt"
	CONST "tcp-soupbin-client/constants"
	"tcp-soupbin-client/routes"

	COMM_CONST "github.com/rs-commons/pkg/constants"
	COMM_DB "github.com/rs-commons/pkg/db"
	comm_err "github.com/rs-commons/pkg/error"

	"github.com/rs-commons/pkg/config"
	"github.com/rs-commons/pkg/web"

	"tcp-soupbin-client/util"

	"github.com/rs-commons/pkg/log"
)

func main() {
	env := initializeConfig()
	logger := initializeLogger(env)
	defer comm_err.RecoverPanic(logger)
	var dbKeys, rdbkeys []string

	for _, key := range []string{
		CONST.MYSQL_ANALYTIC_MASTER,
		CONST.MYSQL_ANALYTIC_SLAVE,
		CONST.MYSQL_CMOTS_MASTER,
		CONST.MYSQL_CMOTS_SLAVE,
	} {
		if env.Get(key) != "" {
			dbKeys = append(dbKeys, key)
		}
	}
	for _, rdbkey := range []string{
		CONST.REDIS_ANALYTICAL_MASTER,
		CONST.REDIS_ANALYTICAL_SLAVE,
	} {
		if env.Get(rdbkey) != "" {
			rdbkeys = append(rdbkeys, rdbkey)
		}
	}
	DBConnections := COMM_DB.DBConnections(env, logger, dbKeys...)
	RedisConn := COMM_DB.SetupRedisClient(env, logger, rdbkeys...)
	serverPort := env.Get(COMM_CONST.SERVER_PORT)

	routes := routes.RoutesList(logger)
	server := &web.WebServer{
		Port:    serverPort,        // Server Port address
		Routes:  routes,            // all route mappings
		Adaptor: &web.GinAdaptor{}, // Gin Web Server API
		Logger:  logger,            // Logger assignment
		Config:  env,               // Config assignment
		SuccessCallBack: func(context *web.AppContext) {
			util.LoadDbConnections(DBConnections, RedisConn)
			context.MysqlConnection[COMM_CONST.MYSQL_MASTER_DSN1] = DBConnections[CONST.MYSQL_ANALYTIC_MASTER]

		},
	}
	// Starting the Server
	server.Launch()

	logger.CommonLog.Info("Launching Server on port : ", serverPort)

	fmt.Println(env)

}

func initializeConfig() *config.Env {
	configpath := flag.String("config", CONST.CONFIG_PATH, "Path to the configuration file")
	flag.Parse()
	fmt.Printf("Using config path: %s\n", *configpath)
	env := config.ConfigFactory{
		ConfigName:    CONST.CONFIG_FILE,
		ConfigPath:    *configpath,
		ConfigAdaptor: &config.ViperAdapter{}, // Viper API Dependency Injection
	}.NewConfigFactory()
	return env
}

func initializeLogger(env *config.Env) *log.Logger {
	logger := log.NewLogger(CONST.APP_NAME, env)
	logger.CommonLog.Info("tcp-soupbin-client")
	return logger
}
