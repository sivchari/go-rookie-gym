package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sivchari/go-rookie-gym/handler"
	"github.com/sivchari/go-rookie-gym/infrastructure"
	groupdb "github.com/sivchari/go-rookie-gym/infrastructure/group"
	userdb "github.com/sivchari/go-rookie-gym/infrastructure/user"
	groupuc "github.com/sivchari/go-rookie-gym/usecase/group"
	useruc "github.com/sivchari/go-rookie-gym/usecase/user"
)

const (
	OSExitOK    = 0
	OSExitError = 1
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	return listen(ctx, 8080)
}

func listen(ctx context.Context, port int) int {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("failed to open a db err = %s", err.Error())
		return OSExitError
	}
	defer db.Close()
	if err := db.PingContext(context.Background()); err != nil {
		log.Printf("failed to ping err = %s", err.Error())
		return OSExitError
	}
	m := infrastructure.NewTxManager(db)
	groupr := groupdb.NewDB(m)
	userr := userdb.NewDB(m)
	groupuc := groupuc.NewUsecase(groupr)
	useruc := useruc.NewUsecase(userr, groupr, m)
	h := handler.NewHandler(groupuc, useruc)

	mux := http.NewServeMux()
	mux.Handle("/group", handler.Validate(h.GroupHandler()))
	mux.Handle("/groups", handler.Validate(h.GroupsHandler()))
	mux.Handle("/user", h.UserHandler())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	// SIGTERM, SIGINTを検知すると通知する
	sigctx, c := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer c()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("[ERROR] %s\n", err.Error())
		}
	}()

	// Signalがくるまでブロック
	<-sigctx.Done()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// Graceful Shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		return OSExitError
	}

	log.Println("[INFO] shuwdown...")
	return OSExitOK
}
