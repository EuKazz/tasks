// **Описание**: Реализуйте CORS middleware-функцию, которая
	// устанавливает необходимые заголовки для кросс-доменных
	// запросов и обрабатывает preflight OPTIONS запросы.
	// **Входные данные**: next http.Handler - следующий
	// обработчик в цепочке
	// **Выходные данные**: http.Handler - обработчик с
	// функциональностью CORS
	// **Ограничения**: Используйте заголовки
	// Access-Control-Allow-Origin, Access-Control-Allow-Methods,
	// Access-Control-Allow-Headers
	// **Примеры**:
	// Input: OPTIONS запрос с Origin: "http://localhost:3000"
	// Output: Ответ со статусом 200 и CORS заголовками
	//
	// Input: GET запрос с Origin: "https://example.com"
	// Output: Обычный ответ с добавленными CORS заголовками
package main

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//получаем заголовок ориг из запроса (он присутствует при кросплатф запросе)
    origin:= r.Header.Get("Origin")
    //если его нет - значит запрос прямой или внутренний - вызываем след обработчик без корс-заголовков и прекращаем работу мидлвэр
      if origin == ""{
      next.ServeHTTP(w,r)
      return
    }
      //получаем ссылку на заголовки ответа
      header:= w.Header()
      //разрешаем доступ этому источнику
      header.Set("Access-Control-Allow-Origin", origin)
      //разрешает браузеру отправлять куки, авторизационные заголовки
      header.Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
          //если пришел префлайт запрос указываем разрешенные методы
			header.Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, PATCH")
		  //указываем разрешенные заголовки
            header.Set("Access-Control-Allow-Headers", "authorization, content-type")
		  //указываем макс время хранения ответа браузером(1 сутки)
            header.Set("Access-Control-Max-Age", "86400")
          //отправляем ОК для завершения обработки запроса и выходим из мидлвэр 
            w.WriteHeader(http.StatusOK) 
			return
		}
      next.ServeHTTP(w,r)
	})
}
