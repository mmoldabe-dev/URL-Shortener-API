package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterFunc(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are requied", http.StatusBadRequest)
		return
	}
	Hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password ", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users(username, password_hash,created_at)
	VALUES ($1,$2,$3)
	RETURNING id;`

	var userID int

	err = DB.QueryRow(query, req.Username, string(Hash), time.Now()).Scan(&userID)
	if err != nil {
		http.Error(w, "Failed to register user (maybe username already exists)", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "User registed",
		"user_id": userID,
	})
}
