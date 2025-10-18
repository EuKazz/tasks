// **Описание**: Реализуйте middleware-функцию для авторизации,
// которая проверяет наличие и валидность токена в заголовке
// Authorization и блокирует доступ неавторизованным пользователям.
//
// **Входные данные**: next http.Handler - следующий обработчик
// в цепочке
//
// **Выходные данные**: http.Handler - обработчик с
// функциональностью авторизации
//
// **Ограничения**: Токен должен начинаться с "Bearer ",
// валидный токен = "valid-token-123", при отсутствии или
// невалидном токене возвращать статус 401
//
// **Примеры**:
// Input: Запрос с заголовком "Authorization: Bearer valid-token-123"
// Output: Запрос передается дальше по цепочке
//
// Input: Запрос без заголовка Authorization или с невалидным токеном
// Output: Ответ со статусом 401 Unauthorized
package main
import(
  "strings"
  "net/http"
)
func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      authedHeader:= r.Header.Get("Authorization")
      if !strings.HasPrefix(authedHeader, "Bearer "){
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
        return
      }
      validToken:= "valid-token-123"
      token := strings.TrimPrefix(authedHeader, "Bearer ")
      if token!= validToken{
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
        return
      }
      
      next.ServeHTTP(w,r)
    })
}
