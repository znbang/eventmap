package bookservice

import (
	"encoding/json"
	"fmt"
	"net/http"
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
