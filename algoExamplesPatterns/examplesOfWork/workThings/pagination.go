package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Article represents the structure of an article
type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Published   string `json:"published"`
	SourceID    int    `json:"source_id"`
}

// GetUserNews gets the news for a user with pagination
func GetUserNews(w http.ResponseWriter, r *http.Request) {
	// Extract userID from URL
	vars := mux.Vars(r)
	userIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Set default values if parameters are not provided
	page := 1
	limit := 2

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get the last login time of the user
	var lastLogin string
	err = db.QueryRowContext(context.Background(), "SELECT last_login FROM users WHERE id=$1", userID).Scan(&lastLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get sources the user is subscribed to
	rows, err := db.QueryContext(context.Background(), "SELECT s.id FROM user_subscriptions us JOIN Sources s ON us.source_id = s.id WHERE us.user_id=$1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sourceIDs []int
	for rows.Next() {
		var sourceID int
		if err := rows.Scan(&sourceID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sourceIDs = append(sourceIDs, sourceID)
	}

	// If no sources, return empty response
	if len(sourceIDs) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Article{})
		return
	}

	// Get articles for the sources after last login time with pagination
	query := `
		SELECT id, title, link, description, published, source_id
		FROM articles
		WHERE source_id = ANY($1) AND published > $2
		ORDER BY published DESC
		LIMIT $3 OFFSET $4`
	articleRows, err := db.QueryContext(context.Background(), query, pq.Array(sourceIDs), lastLogin, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer articleRows.Close()

	var articles []Article
	for articleRows.Next() {
		var article Article
		if err := articleRows.Scan(&article.ID, &article.Title, &article.Link, &article.Description, &article.Published, &article.SourceID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	// Respond with the articles
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Initialize your database connection here (db)
	// db = ...

	// Create a new router
	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/sources/{id}/news", GetUserNews).Methods("GET")

	// Start the server
	http.ListenAndServe(":8080", r)
}
