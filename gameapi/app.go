package gameapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/task4233/techtrain-mission/gameapi/config"
	"github.com/task4233/techtrain-mission/gameapi/handler"
	"github.com/task4233/techtrain-mission/gameapi/infra"
	"github.com/task4233/techtrain-mission/gameapi/log"
	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

var (
	logger = log.MyLogger
)

// GameAPI is struct for GameAPI
type GameAPI struct {
	db *sqlx.DB
}

// NewGameAPI returns a pointer to GameAPI struct
func NewGameAPI() *GameAPI {
	return &GameAPI{}
}

// Run runs
func (g *GameAPI) Run() error {
	if err := g.setDB(); err != nil {
		return fmt.Errorf("failed setDB(): %w", err)
	}

	if err := g.injectDependencies(); err != nil {
		return fmt.Errorf("failed injectDependencies(): %w", err)
	}

	// start serving
	srv := &http.Server{Addr: ":" + config.Port()}
	logger.Warnf("Start App: listening on port %s", config.Port())

	// graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Warnf("shutdown the server with error: %+v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	logger.Warnf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Warnf("failed to shutdown: %+v", err)
		return err
	}

	return errors.New("correctly Shutdown")
}

func (g *GameAPI) setDB() (err error) {
	g.db, err = infra.NewDB()
	if err != nil {
		logger.Warnf("failed newDB: %w", err)
		os.Exit(1)
	}
	if g.db == nil {
		logger.Warnf("failed newDB: %w", err)
		os.Exit(1)
	}
	defer func() {
		cerr := g.db.Close()
		if err != nil {
			logger.Warnf("failed Close(); %w", cerr)
		}
	}()
	return err
}

func (g *GameAPI) injectDependencies() error {
	// User
	userRepo := infra.NewUserRepository(g.db)
	userUC := usecase.NewUser(userRepo)
	user := handler.NewUser(userUC)

	http.HandleFunc("/user/create", user.Create)
	http.HandleFunc("/user/get", user.Get)
	http.HandleFunc("/user/update", user.Update)

	return nil
}
