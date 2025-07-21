package main

import (
	"context"
	"fmt"
	"github.com/Malware3447/configo"
	"github.com/Malware3447/spg"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sbh/internal/config"
	repo "sbh/internal/repository/pg"
	"sbh/internal/requests"
	"sbh/internal/requests/new_user"
	pgService "sbh/internal/services/pg"
	"syscall"
)

func main() {
	cfg, _ := configo.MustLoad[config.Config]()
	log := logrus.New()
	ctx := context.Background()

	poolPg, err := spg.NewClient(ctx, &cfg.DatabasePg)
	if err != nil {
		log.Error(fmt.Errorf("ошибка при запуске Postgres: %s", err))
		panic(err)
	}
	log.Info("Postgres успешно запущен")

	repoPg := repo.NewPgRepository(poolPg)

	pgService := pgService.NewPgService(repoPg)

	us := new_user.NewUser(pgService)

	router := requests.NewApi(us)

	err = router.Init()
	if err != nil {
		log.Error("%v", err)
	}

	log.Info("Сервис успешно запущен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-quit:
		log.Println("Завершение работы сервиса")
	}

	log.Info("Сервис успешно завершил работу")
}
