package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"URL-Shortener-API/auth"
	"URL-Shortener-API/models"

	"github.com/gorilla/mux"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	auth.RegisterFunc(w, r, DB)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	auth.LoginFunc(w, r, DB)
}

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
	vars := mux.Vars(r)
	shortcode := vars["code"]
	if shortcode == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	originalURL, err := GetUrlByShortcode(DB, shortcode)
	if err != nil {
		http.Error(w, "Short code not found", http.StatusNotFound)
		return
	}

	fmt.Println("Redirecting to:", originalURL)
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func PostUrl(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var req struct {
		OriginalURL string `json:"original_url"`
		TTLSeconds  int    `json:"ttl_seconds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	query := `
		INSERT INTO urls (short_code, original_url, created_at, ttl_seconds)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	var id int
	err := DB.QueryRow(query, shortcode, req.OriginalURL, createdAt, ttl).Scan(&id)
	if err != nil {
		log.Printf("Error inserting URL: %v", err)
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ShortCode string `json:"short_code"`
	}{ShortCode: shortcode})
}

func GetDB(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	rows, err := DB.Query("SELECT id, short_code, original_url, created_at, ttl_seconds FROM urls")
	if err != nil {
		log.Printf("Error selecting URLs: %v", err)
		http.Error(w, "Error selecting from table", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []models.URL
	for rows.Next() {
		var u models.URL
		if err := rows.Scan(&u.Id, &u.ShortCode, &u.OriginalURL, &u.CreatedAt, &u.TTLSecond); err != nil {
			http.Error(w, "Error scanning table", http.StatusInternalServerError)
			return
		}
		result = append(result, u)
	}

	if len(result) == 0 {
		http.Error(w, "No URLs found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetUrlByShortcode(DB *sql.DB, shortcode string) (string, error) {
	var originalURL string
	err := DB.QueryRow("SELECT original_url FROM urls WHERE short_code = $1", shortcode).Scan(&originalURL)
	return originalURL, err
}
