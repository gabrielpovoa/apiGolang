package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks   = []Task{}
	nextID  = 1
	taskMux = &sync.Mutex{}
)

// main inicia o servidor HTTP e registra as rotas "/tasks" e
// "/tasks/{id}" para manipular a lista de tarefas. O servidor
// HTTP é executado na porta 8080 e o seu log de erros é
// impresso na sa da.
func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTaskByID)

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Manipula requisições para a lista de tarefas.
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Manipula requisições para uma tarefa específica por seu ID.
//
// Requisições GET buscam uma tarefa pelo seu ID.
// Requisições DELETE excluem uma tarefa pelo seu ID.
//
// Caso o ID seja inválido, o status 400 Bad Request é retornado.
// Caso o ID seja válido, mas a tarefa não exista, o status 404 Not Found é
// retornado.
// Caso o método seja desconhecido, o status 405 Method Not Allowed é retornado.
func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, id)
	case http.MethodDelete:
		deleteTask(w, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Retorna a lista de tarefas em formato JSON.
//
// A lista de tarefas é bloqueada para evitar que outras goroutines
// acessem enquanto a lista está sendo serializada.
//
// A saída da resposta é configurada para ter o tipo de conteúdo
// "application/json" e o conteúdo da lista de tarefas é
// serializado com o pacote "encoding/json".
func listTasks(w http.ResponseWriter, r *http.Request) {
	taskMux.Lock()
	defer taskMux.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Cria uma nova tarefa a partir dos dados fornecidos no corpo da requisição.
//
// Se os dados da tarefa forem válidos, um novo ID é atribuído, e a tarefa é
// adicionada à lista de tarefas. O status 201 Created é retornado junto com
// os dados da nova tarefa serializada em JSON. Caso contrário, o status
// 400 Bad Request é retornado.

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	taskMux.Lock()
	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	taskMux.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// Busca uma tarefa pelo seu id
//
// Se a tarefa for encontrada, o status 200 OK é retornado, e o conteúdo da
// tarefa é serializado em JSON e enviado na resposta.
// Caso contrário, o status 404 Not Found é retornado.
func getTask(w http.ResponseWriter, id int) {
	taskMux.Lock()
	defer taskMux.Unlock()

	for _, t := range tasks {
		if t.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.NotFound(w, nil)
}

// Deleta uma tarefa pelo seu id
//
// Se a tarefa for encontrada, o status 204 No Content é retornado.
// Caso contrário, o status 404 Not Found é retornado.
func deleteTask(w http.ResponseWriter, id int) {
	taskMux.Lock()
	defer taskMux.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, nil)
}
