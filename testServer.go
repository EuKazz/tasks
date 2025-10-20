	// **Описание**: Реализуйте функцию для создания тестового
	// HTTP-сервера и выполнения end-to-end теста регистрации
	// пользователя через все слои приложения.
	// **Входные данные**: Функция App() возвращающая http.Handler,
	// структура User для тестовых данных
	// **Выходные данные**: Функция TestUserRegistrationE2E,
	// которая создает тестовый сервер, отправляет POST-запрос
	// и проверяет корректность ответа
	// **Ограничения**: Использовать httptest.NewServer, http.Post
	// для отправки JSON, проверить статус-код 201
	// **Примеры**:
	// Input: POST /register с JSON
	// {"email":"newuser@test.com","password":"password123"}
	// Output: HTTP 201 Created с JSON
	// {"message":"User registered successfully"}
	//
	// Input: POST /register с JSON
	// {"email":"existing@test.com","password":"pass456"}
	// Output: HTTP 400 Bad Request с JSON
	// {"error":"User already exists"}

	// TODO: Создать тестовый HTTP-сервер
	// TODO: Подготовить JSON данные для регистрации
	// TODO: Отправить POST-запрос
	// TODO: Проверить статус-код и тело ответа
	// TODO: Очистить тестовое окружение
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRegistrationE2E(t *testing.T) {
  //запуск тестового сервера
  ts:= httptest.NewServer(App())
  defer ts.Close()
  //инициализации юзера с тестовыми данными
  testUser:= User{
    Email: "test@a.ru",
    Password: "dfber"
  }
  //переводит структуру в нужный формат
  req, err:= json.Marshal(testUser)
  if err!=nil{
    t.Fatal("could not marshal this user")
  }
  //first post  - success registration of new user
  resp, err := http.Post(ts.URL+"/register", "application/json", bytes.NewReader(req))
	if err != nil {
		t.Fatal("req failed")
	}
  defer resp.Body.Close()
	//успешный кейс вернет 201
  if resp.StatusCode != 201 {
		t.Fatalf("Expected %d got %d", 201, resp.StatusCode)
	}
  //ожидаем что в теле будет одно поле - сообщение
  var messagePositive struct{
    Message string `json:"message"`
  }
//декодируем в анонимную структуру с этим полем
  err = json.NewDecoder(resp.Body).Decode(&messagePositive)
  if err!=nil{
    t.Fatalf("could not decode successfully")
  }
  //сравниваем фактическое значение с ожидаемым
  want:= "User registered successfully"
  if messagePositive.Message!= want{
    t.Fatalf("unexpected message: got %q, want %q", messagePositive.Message, want)
  }

//second post - registration of user already exists 
  //имитируем повторную регистрацию с тем же емейл
  resp2, err := http.Post(ts.URL+"/register", "application/json", bytes.NewReader(req))
	if err != nil {
		t.Fatal("req failed")
	}
  defer resp2.Body.Close()
  //ожидаем статус 400
  if resp2.StatusCode != 400 {
		t.Fatalf("Expected %d got %d", 400, resp2.StatusCode)
	}
//в отвеет ожидаем сообщение с полем ошибка
  var msgError struct{
    Error string `json:"error"`
  }
  //декодируем
  err = json.NewDecoder(resp2.Body).Decode(&msgError)
    if err != nil {
        t.Fatalf("could not decode error json: %v", err)
    }
  wantErr:= "User already exists"
  if msgError.Error!= wantErr{
    t.Fatalf("unexpected error message: got %q, want %q", msgError.Error, wantErr)
  }
}
