	// **Описание**: Реализуйте функцию, которая выполняет
	// HTTP-запрос с тайм-аутом, используя контекст для
	// ограничения времени выполнения.
	// **Входные данные**: url (строка), timeout (time.Duration)
	// **Выходные данные**: тело ответа ([]byte) и ошибка
	// **Ограничения**: timeout должен быть положительным
	// значением, url должен быть валидным HTTP URL
	// **Примеры**:
	// Input: url = "https://httpbin.org/delay/1",
	// timeout = 2*time.Second
	// Output: тело ответа, nil
	//
	// Input: url = "https://httpbin.org/delay/5",
	// timeout = 1*time.Second
	// Output: nil, context deadline exceeded
package main

import (
	"context"
	"io"
	"net/http"
	"time"
  "errors"
)

func fetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
    if timeout <=0{
      return nil, errors.New("timeout must be positive")
    }

  if !(len(url)>0 && (url[:7]=="http://" || url[:8] =="https://")){
     return nil, errors.New("url must start with http:// or https://")
  }
  //родительский контекст
    context:= context.Background()
  //контекст с таймаутом
    ctx, cancel:= context.WithTimeout(context, timeout)
  //закрытие освободит ресурсы, связанные с таймаутом
    defer cancel()
//создает гет запрос с контекстом( по истечении запрос будет остановлен)
    req, err:= http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
  if err!=nil{
    return nil, err
  }
//запрос отправляется серверу
  resp, err:= http.DefaultClient.Do(req)
  if err!=nil{
    return nil, err
  }
  //если запрос прошел успешно - закрывает тело ответа после чтения
defer resp.Body.Close()
//все байты считываются в переменную (чтение ответа)
  body,err:= io.ReadAll(resp.Body)
  if err!=nil{
    return nil, err
  }
