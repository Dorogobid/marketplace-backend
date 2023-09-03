package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/Dorogobid/marketplace-backend/config"
	"github.com/Dorogobid/marketplace-backend/internal/repository"
	"github.com/Dorogobid/marketplace-backend/internal/server"
	svc "github.com/Dorogobid/marketplace-backend/internal/service"
)

func Execute() {
	conf := config.NewConf()

	db, err := sql.Open("postgres", conf.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := migrateDB(conf.DB.DSN()); err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	service := svc.New(repo)

	s := server.NewServer(service, conf.XAPIKey, conf.BaseURL)
	s.Logger().Fatal(s.Start(fmt.Sprintf(":%s", conf.Port)))
}
