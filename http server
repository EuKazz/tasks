//Описание: Реализуйте HTTP-сервер, который обрабатывает запросы на разных маршрутах и возвращает соответствующие HTTP-статус-коды в зависимости от метода запроса.
//Входные данные: HTTP-запросы с различными методами (GET, POST, PUT, DELETE) на маршруты "/users" и "/health"
//Выходные данные: HTTP-ответы с корректными статус-кодами и телом ответа
//Ограничения: Сервер должен работать на порту 8080, поддерживать только указанные методы и маршруты
//Примеры:
//Input: GET /health
//Output: 200 OK, {"status": "healthy"}
//Input: POST /users
//Output: 201 Created, {"message": "User created"}
//nput: PUT /nonexistent
//Output: 404 Not Found
package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	

	mux := http.NewServeMux()
	
	// Регистрируйте обработчики для маршрутов
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/users", usersHandler)
	
	http.ListenAndServe(":8080", mux)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
    case http.MethodGet:
      w.Header().Set("content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      resp:= map[string]string{"status": "healthy"}
      json.NewEncoder(w).Encode(resp)
    case http.MethodPost:
       http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    case http.MethodPut:
       http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    case http.MethodDelete:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    default:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
    case http.MethodPost:
      w.Header().Set("content-Type", "application/json")
      w.WriteHeader(http.StatusCreated)
      resp:= map[string]string{"message": "User created"}
      json.NewEncoder(w).Encode(resp)
    case http.MethodGet:
       http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    case http.MethodPut:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    case http.MethodDelete:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    default:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
