	// **Описание**: Реализуйте функцию, которая записывает
	// пользовательские данные в контекст и создает новый
	// HTTP-запрос с обновленным контекстом.
	// **Входные данные**: r (*http.Request) - исходный
	// HTTP-запрос, userID (string) - идентификатор
	// пользователя для записи в контекст
	// **Выходные данные**: новый HTTP-запрос (*http.Request)
	// с обновленным контекстом
	// **Ограничения**: userID не должен быть пустой строкой,
	// r не должен быть nil
	// **Примеры**:
	// Input: r = HTTP-запрос, userID = "user123"
	// Output: новый HTTP-запрос с userID в контексте
	//
	// Input: r = HTTP-запрос, userID = "admin456"
	// Output: новый HTTP-запрос с admin456 в контексте
package main

import (
	"context"
	"net/http"
  "io"
  "bytes"
)

// Определяем тип ключа для контекста
type contextKey string

const UserIDKey contextKey = "userID"

func addUserToContext(r *http.Request, userID string) *http.Request {
  if userID==""{
    return nil
  }
  if r == nil{
    return nil
  }
  method:= r.Method
  url:= r.URL.String()
  bodyBytes, err:= io.ReadAll(r.Body)
  if err != nil {
        return nil
    }
  defer r.Body.Close()
  //после чтения запроса через ioданные надо снова записать для нового запроса
  body:= bytes.NewReader(bodyBytes)
  

  ctxBack:= r.Context()
  ctx:= context.WithValue(ctxBack, UserIDKey,userID )
  req,err:= http.NewRequestWithContext(ctx, method, url, body)
  if err!=nil{
    return nil
  }
	return req
}
