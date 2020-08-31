package gocrm

import (
	"database/sql"
	"net/http"

	"log"

	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq" // ...
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

func newDB(databaseURL string) (*sql.DB, error) {
	log.Println("[SQL]: Connecting to SQL..")
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}
	log.Println("[SQL]: Success!")
	return db, nil
}
