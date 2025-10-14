// **Описание**: Реализуйте структуру для обработчика
	// HTTP-запросов с методами для различных эндпоинтов и
	// фабричную функцию для её создания.
	// **Входные данные**: HTTP-запросы к различным эндпоинтам
	// через ServeMux
	// **Выходные данные**: HTTP-ответы с соответствующими
	// статус-кодами и JSON-данными
	// **Ограничения**: Использовать структуру-обработчик,
	// методы структуры вместо функций, фабричный конструктор
	// **Примеры**:
	// Input: GET /products
	// Output: 200 OK, {"products": ["laptop", "mouse"]}
	//
	// Input: POST /products
	// Output: 201 Created, {"message": "Product added"}

package main

import (
	"net/http"
     "encoding/json"
)

// ProductHandler представляет обработчик для работы с продуктами
type ProductHandler struct {
	
}

// NewProductHandler создает новый экземпляр ProductHandler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      resp:= map[string][]string{"products": []string{"laptop", "mouse"}}
      json.NewEncoder(w).Encode(resp)
}
func (p *ProductHandler) Post(w http.ResponseWriter, r *http.Request){
      
    w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusCreated)
      resp:= map[string]string{"message": "Product added"}
      json.NewEncoder(w).Encode(resp)
}
func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request){
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
func (p *ProductHandler) Put(w http.ResponseWriter, r *http.Request){
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
// HandleProducts обрабатывает запросы к /products
func (p *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
    case http.MethodGet:
      p.Get(w,r)
    case http.MethodPost:
      p.Post(w,r)
    case http.MethodPut:
      p.Put(w,r)
    case http.MethodDelete:
      p.Delete(w,r)
    default:
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    p:= NewProductHandler()
	mux:= http.NewServeMux()
    mux.HandleFunc("/products", p.HandleProducts)
    http.ListenAndServe(":8080", mux)
}
