package api

import (
	storage "accessCloude/internal/storage"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type AccessCloude struct {
	DB *storage.Database
}

func NewAccessCloude(db *storage.Database) *AccessCloude {
	return &AccessCloude{
		DB: db,
	}
}

var _ ServerInterface = &AccessCloude{}

func UnmarshalObject[T any](r *http.Request, obj *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		comment := fmt.Errorf("error reading request body: %v", err)
		slog.Error(comment.Error())
		return comment
	}

	if unmErr := json.Unmarshal(body, &obj); unmErr != nil {
		commment := fmt.Errorf("error unmarshalling request body: %v", unmErr)
		slog.Error(commment.Error())
		return err
	}

	return nil
}

func Response[T any](w http.ResponseWriter, v T, status int) {
	w.Header().Set("Content-Type", "application/json")

	var messages interface{} = v
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		comment := fmt.Errorf("error encoding response body: %v", err)
		slog.Error(comment.Error())

		messages := map[string]interface{}{
			"message": v,
			"error":   comment.Error(),
		}

		data, err := json.Marshal(messages)
		if err != nil {
			comment := fmt.Errorf("error encoding response body: %v", err)
			slog.Error(comment.Error())
			messages["error"] = comment
			return
		}

		w.Write(data)
	}
}
