package handler

import (
	"URL-Shortener-API/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	switch r.Method {
	case http.MethodPost:
		PostUrl(w, r, DB)
	case http.MethodGet:
		GetDB(w, r, DB)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ShortenHandlerId(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		shortcode := vars["code"]

		if shortcode == "" {
			http.Error(w, "Short code is nil", http.StatusBadRequest)
			return
		}

		originalURL, err := GetUrlByShortcode(DB, shortcode)
		if err != nil {

			http.Error(w, "Short code not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostUrl(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var req struct {
		OriginalURL string `json:"original_url"`
		TTLSeconds  int    `json:"ttl_seconds"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "Original URL is required", http.StatusBadRequest)
		return
	}

	shortcode := Base62Encode(uint64(time.Now().UnixNano()))

	ttl := req.TTLSeconds
	createdAt := time.Now()

	query := `INSERT INTO urls (short_code, original_url, created_at, ttl_seconds)
     VALUES ($1, $2, $3, $4)
     RETURNING id`

	var id int

	err = DB.QueryRow(query, shortcode, req.OriginalURL, createdAt, ttl).Scan(&id)
	if err != nil {
		log.Printf("Error inserting URL: query=%q shortcode=%s original_url=%s ttl=%d: %v",
			query, shortcode, req.OriginalURL, ttl, err,
		)
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		ShortCode string `json:"short_code"`
	}{
		ShortCode: shortcode,
	}

	json.NewEncoder(w).Encode(response)
}

func GetDB(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	rows, err := DB.Query("SELECT id, short_code, original_url, created_at, ttl_seconds FROM urls")
	if err != nil {
		log.Printf("Error selecting URLs: %v", err)
		http.Error(w, "Error selecting from table", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	result := []models.URL{}

	for rows.Next() {
		var res models.URL

		if err := rows.Scan(&res.Id, &res.ShortCode, &res.OriginalURL, &res.CreatedAt, &res.TTLSecond); err != nil {
			http.Error(w, "Erorr scan table", http.StatusInternalServerError)
			return
		}

		result = append(result, res)
	}
	if len(result) == 0 {
		http.Error(w, "No URLs found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
func GetUrlByShortcode(DB *sql.DB, shortcode string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_code = $1`
	var originalURL string

	err := DB.QueryRow(query, shortcode).Scan(&originalURL)
	if err != nil {

		return "", err
	}

	return originalURL, nil
}
