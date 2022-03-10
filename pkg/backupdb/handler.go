package backupdb

import (
	"fmt"
	"net/http"
	"time"

	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/bookservice"
	"github.com/znbang/eventmap/pkg/eventservice"
	"github.com/znbang/eventmap/pkg/login"
	"github.com/znbang/eventmap/pkg/userservice"
)

func CreateBackupDatabaseHandler(db *dbx.DB, createTableSql string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = fmt.Fprint(w, "Method not allowed")
		}

		bak, err := dbx.Sqlite("eventmap." + time.Now().Format("20060102150405") + ".db")
		// bak, err := dbx.Open(env.Get(env.DatabaseURL))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "Open database failed: %v", err)
		}
		if err = bak.Exec(createTableSql).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "Create tables failed: %v", err)
		}
		defer bak.Close()

		models := []any{
			&[]userservice.User{},
			&[]login.UserSession{},
			&[]bookservice.Book{},
			&[]bookservice.BookJob{},
			&[]bookservice.Chapter{},
			&[]eventservice.Event{},
		}

		for _, model := range models {
			if err = dbx.CopyTable(model, db, bak); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprintf(w, "%v", err)
			}
			model = nil
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, "Done")
	}
}
