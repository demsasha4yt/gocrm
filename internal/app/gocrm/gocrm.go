package gocrm

import (
	"context"
	"net/http"
	"time"

	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

// Start aplication
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 100
	config.MaxConnLifetime = time.Second * 15
	config.ConnConfig.LogLevel = pgx.LogLevelTrace
	config.ConnConfig.Logger = logrusadapter.NewLogger(logrus.New())

	return pgxpool.ConnectConfig(context.Background(), config)
}
