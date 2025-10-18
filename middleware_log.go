// **Описание**: Реализуйте middleware-функцию для логирования
	// HTTP-запросов, которая измеряет время выполнения запроса
	// и выводит информацию о методе, пути и времени выполнения.
	//
	// **Входные данные**: next http.Handler - следующий обработчик
	// в цепочке
	//
	// **Выходные данные**: http.Handler - обработчик с
	// функциональностью логирования
	//
	// **Ограничения**: Используйте стандартный пакет log для вывода,
	// время должно быть в миллисекундах
	//
	// **Примеры**:
	// Input: GET запрос на /users, выполняющийся 150ms
	// Output: "GET /users completed in 150ms"
	//
	// Input: POST запрос на /api/data, выполняющийся 75ms
	// Output: "POST /api/data completed in 75ms"
package main

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	start:= time.Now()
        next.ServeHTTP(w,r)
        duration:= time.Since(start).Milliseconds()
        log.Printf("%s %s completed in %dms", r.Method, r.URL.Path, duration)
      
	})
}
