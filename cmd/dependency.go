package main

import (
	"database/sql"
	"gin-frame/api/controllers"
	"gin-frame/api/middlewares"
	"gin-frame/api/routes"
	configManager "gin-frame/lib/config"
	"gin-frame/lib/logger"
	"gin-frame/services"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func initSetting() {
	//init config
	configManager.Global = configManager.Reload()
	//init log
	log.Logger = logger.InitLogger()
}

func initializeService() (*controllers.Handler, *sql.DB, error) {

	handler := controllers.NewHandler()

	jobService := services.NewJobService()

	publicController := controllers.NewPublicController()
	publicRoutes := routes.NewPublicRoutes(handler, publicController)

	jobController := controllers.NewJobController(jobService)
	jobRoutes := routes.NewJobRoutes(handler, jobController)

	routes := routes.NewRoutes(publicRoutes, jobRoutes)
	routes.Setup()

	middlewares := middlewares.NewCorsMiddleware(handler)
	middlewares.Setup()

	mysqlDB := ConnectMysqlDB()

	return &handler, mysqlDB, nil
}

func ConnectMysqlDB() *sql.DB {
	mysqlConfig := mysql.Config{
		User:                 configManager.Global.Api.Mysql.Username,
		Passwd:               configManager.Global.Api.Mysql.Password,
		Addr:                 configManager.Global.Api.Mysql.Address,
		Net:                  "tcp",
		DBName:               configManager.Global.Api.Mysql.Database,
		AllowNativePasswords: true,
		MultiStatements:      true,
	}

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Panic().Err(err).Msgf("open database connection failed")
	}
	if err = db.Ping(); err != nil {
		log.Panic().Err(err).Msgf("mysql db connection failed")
	}

	driver, _ := migrateMySQL.WithInstance(db, &migrateMySQL.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Panic().Err(err).Msgf("mysql db migration new Instance failed")
	}

	err = m.Steps(1)
	if err != nil {
		log.Warn().Msgf("mysql db migration, %s", err.Error())
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	return db
}
