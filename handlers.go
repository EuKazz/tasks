package main

import (
	"net/http"
)

// ProductHandler представляет обработчик для работы с продуктами
type ProductHandler struct {
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
}

// NewProductHandler создает новый экземпляр ProductHandler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// HandleProducts обрабатывает запросы к /products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	// Создание ServeMux и регистрация обработчиков
}
