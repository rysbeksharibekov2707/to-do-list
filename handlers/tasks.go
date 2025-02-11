package handlers

import (
	"encoding/json"
	"net/http"
	"todoactual/database"
	"todoactual/models"
)

// GetTasks - получение всех задач
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, details, done FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Details, &task.Done); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
