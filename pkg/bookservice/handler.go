package bookservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintf(w, "%v", err)
}

func CreateSyncStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Cache-control", "no-cache")

		ch := EventBus.Register()
		defer EventBus.Unregister(ch)

		for {
			select {
			case event := <-ch:
				if data, err := json.Marshal(event); err != nil {
					internalServerError(w, err)
				} else if _, err = fmt.Fprintf(w, "data: %v\n\n", string(data)); err != nil {
					internalServerError(w, err)
				}
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			case <-r.Context().Done():
				return
			}
		}
	}
}

func CreateDownloadBookHandler(bookService *BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		var (
			book     Book
			chapters []Chapter
			bookId   string
		)

		params := q.URL.Query()
		bookId = params.Get("id")
		if bookId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := bookService.BookRepository.FindById(&book, bookId); err != nil {
			internalServerError(w, err)
			return
		}

		if err := bookService.GetAllChapters(&chapters, bookId); err != nil {
			internalServerError(w, err)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''"+url.PathEscape(book.Title)+".txt")

		for _, chapter := range chapters {
			texts := [...]string{chapter.Title, "\n\n", chapter.Body, "\n\n"}
			for _, text := range texts {
				data := []byte(text)
				if _, err := w.Write(data); err != nil {
					internalServerError(w, err)
					return
				}
			}
		}
	}
}
