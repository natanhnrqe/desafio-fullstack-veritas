package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"

	"github.com/go-chi/chi/v5"
)

// GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT id, title, description, status, create_at 
		FROM tasks 
		ORDER BY create_at DESC
	`)
	if err != nil {
		log.Println("Erro no Query:", err)
		respondWithError(w, http.StatusInternalServerError, "Erro ao buscar tarefas")
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
    var t Task
    var desc string
    err := rows.Scan(&t.ID, &t.Title, &desc, &t.Status, &t.CreatedAt)
    if err != nil {
        log.Println("Erro no Scan:", err)
        respondWithError(w, http.StatusInternalServerError, "Erro ao ler tarefa")
        return
    }
    t.Description = &desc
    tasks = append(tasks, t) // ← dentro do loop
}

if err := rows.Err(); err != nil {
    log.Println("Erro após iteração:", err)
    respondWithError(w, http.StatusInternalServerError, "Erro ao listar tarefas")
    return
}

respondWithJSON(w, http.StatusOK, tasks)
}

// POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON Invalido!!!")
		return
	}

	if t.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Titulo obrigatorio!!!")
		return
	}

	if t.Status == "" {
		t.Status = "todo"
	}

	if !ValidStatuses[t.Status] {
		respondWithError(w, http.StatusBadRequest, "Status invalido!!!")
		return
	}

	var desc string
	if t.Description != nil {
		desc = *t.Description
	}

	err := DB.QueryRow(`
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, create_at
	`, t.Title, desc, t.Status).Scan(&t.ID, &t.CreatedAt)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao criar tarefa!!!")
		return
	}

	respondWithJSON(w, http.StatusCreated, t)
}

// PUT /task/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID invalido!!!")
		return
	}

	var t Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON invalido!!!")
		return
	}

	if t.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Titulo obrigatorio!!!")
		return
	}

	if !ValidStatuses[t.Status] {
		respondWithError(w, http.StatusBadRequest, "Status invalido!!!")
		return
	}

	result, err := DB.Exec(`
    UPDATE tasks 
    SET title = $1, description = $2, status = $3
    WHERE id = $4
`, t.Title, t.Description, t.Status, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao atualizar tarefa!!!")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
		return
	}
	t.ID = id
	respondWithJSON(w, http.StatusOK, t)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID invalido!!!")
		return
	}


	result, err := DB.Exec("DELETE FROM tasks WHERE id = $1", id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao deletar tarefa!!!")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Tarefa não encontrada!!!")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Tarefa deletada"})
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}