// **Описание**: Реализуйте программу для создания и настройки
// тестового HTTP-сервера с использованием httptest.NewRecorder
// для тестирования HTTP хендлеров.
//
// **Входные данные**: HTTP хендлер функция, HTTP метод
// (GET/POST), URL путь, тело запроса (JSON)
//
// **Выходные данные**: Функция тестирования, которая создает
// запрос через httptest.NewRequest, выполняет его и проверяет
// статус-код и содержимое ответа
//
// **Ограничения**: Использовать httptest.NewRequest и
// httptest.NewRecorder, проверить статус-код и JSON ответ
//
// **Примеры**:
// Input: handler = getUserHandler, method = "GET",
// path = "/user/123", body = nil
// Output: HTTP 200 OK с JSON {"id":123,"name":"John Doe"}
//
// Input: handler = createUserHandler, method = "POST",
// path = "/user", body = {"name":"Jane"}
// Output: HTTP 201 Created с JSON
// {"message":"User created successfully"}
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
    "encoding/json"
  "strconv"
)

func TestHTTPHandler(t *testing.T, handler http.HandlerFunc, method string, path string, body []byte, wantStatus int, wantBody map[string]interface{} ){
	//создаем запрос -но он отправляет не запрос по сети, а объект для передачи в хэндлер
    req:= httptest.NewRequest(method, path, bytes.NewReader(body))
	//в вирт райтер хэндлер позже запишем ответ - вывод сохранится в память
    rec:= httptest.NewRecorder()
    //запускаем обработчик
    handler(rec, req)
    //получили готовый ответ
    resp:= rec.Result()
    defer resp.Body.Close()
    if resp.StatusCode!= wantStatus{
      t.Errorf("ожидался статус %d, получен %d", wantStatus, resp.StatusCode)
    }
  //декодируем тело ответа
  var got map[string]interface{}
  if err:= json.NewDecoder(resp.Body).Decode(&got); err!=nil{
    t.Fatalf("не удалось декодировать JSON ответ: %v", err)
  }
  // для каждого поля проверяем полученные значения
  for k, v := range wantBody {
        if got[k] != v {
            t.Errorf("ожидалось %q `==` %v, получено %v", k, v, got[k])
        }
    }
}
//get - handler
func getUserHandler(w http.ResponseWriter, r *http.Request){
  //для тестов обрезаем строку
  idStr:= r.URL.Path[len("/user/"):]
  //преобразует строку в целое число
  id,err:= strconv.Atoi(idStr)
  if err != nil {
        http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
        return
    }
  //выставляет заголовок, что приходит именно json
  w.Header().Set("Content-Type", "application/json")
  //превращает мапу в json и отправляет в поток клиенту
  json.NewEncoder(w).Encode(map[string]interface{}{
    "id": id,
    "name": "John Doe",
  })
}
//post - handler для сервера принимает райтер для формирования ответа и объект запроса
func createUserHandler(w http.ResponseWriter, r *http.Request) {
   //структура для входных данных запроса 
   type req struct {
        Name string `json:"name"`
    }
  //создаем переменную данного типа - в нее распарсим json из тела запроса
    var data req
  //раскодируем из тела запроса в структуру
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil || data.Name `==` "" {
        http.Error(w, `{"error":"invalid input"}`, http.StatusBadRequest)
        return
    }
  //выставляем заголовок
    w.Header().Set("Content-Type", "application/json")
  //выставляем статус
    w.WriteHeader(http.StatusCreated)
  //кодируем и отправляем json ответ с полем сообщения
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User created successfully",
    })
}
//тестовая функция принимает объект t *testing.T для логирования и фиксации ошибок
func TestHandlers(t *testing.T){
  TestHTTPHandler(t,
                  getUserHandler, 
                  http.MethodGet,
                  "/user/123", 
                  //гет-запросы обычно без тела
                  nil, 
                  http.StatusOK,
                  map[string]interface{}{"id": float64(123), "name": "John Doe"},
                 )
  //те создали вирт гет запрос и скормили его хэндлеру

  //тест пост-запроса
  //создаем тело запроса и маршаллим как слайс байт
  body, _:= json.Marshal(map[string]string{"name": "Jane"})
  TestHTTPHandler(t,
                 createUserHandler,
                 http.MethodPost,
                 "/user",
                  //сериализованное тело - новый пользователь
                 body,
                 http.StatusCreated,
                 map[string]interface{}{"message": "User created successfully"},
                 )
}
