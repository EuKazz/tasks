// **Описание**: Реализуйте функцию для создания и выполнения
	// юнит-теста JWT токена с проверкой корректности генерации
	// и парсинга.
	// **Входные данные**: Структура с данными пользователя
	// (ID, Email), секретный ключ для подписи токена
	// **Выходные данные**: Функция TestJWTTokenGeneration,
	// которая создает токен, парсит его и проверяет соответствие
	// исходным данным
	// **Ограничения**: Использовать библиотеку jwt-go, проверить
	// корректность claims, обработать ошибки парсинга
	// **Примеры**:
	// Input: userData = UserClaims{ID: 123, Email: "user@test.com"},
	// secretKey = "test-secret"
	// Output: Успешная генерация и валидация токена с корректными
	// claims
	//
	// Input: invalidToken = "invalid.jwt.token",
	// secretKey = "test-secret"
	// Output: Ошибка парсинга токена с соответствующим сообщением
	// об ошибке
package main

import (
	"testing"
	"github.com/dgrijalva/jwt-go"
  "errors"
  
  "fmt"
)

type UserClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func TestJWTTokenGeneration(t *testing.T) {
  //создаем входные данные
	claims := UserClaims{
        ID:    123,
        Email: "user@test.com",
    }
    secretKey := []byte("test-secret")
  //генерация токена
    token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims )
  
  tokenStr, err:= token.SignedString([]byte(secretKey))
  if err!=nil{
     t.Fatalf("не удалось зашифровать токен: %v", err)
  }
  //парсим валидный токен
  claimsParsed:= &UserClaims{}
  tokenParsed, err:= jwt.ParseWithClaims(tokenStr, claimsParsed, func(token *jwt.Token)(interface{}, error){
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("неожиданный signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
  })
  if err!=nil{
      t.Fatalf("не удалось распарсить токен: %v", err)
  }
  if !tokenParsed.Valid {
        t.Fatal("токен невалиден")
    }

  //  проверка claims
  if claimsParsed.ID!= claims.ID{
    t.Errorf("ожидался ID %d, получено %d",claims.ID, claimsParsed.ID)
  }
  if claimsParsed.Email != claims.Email {
        t.Errorf("ожидался Email %s, получено %s", claims.Email, claimsParsed.Email)
    }

  //парсим некорректный токен
  invalidTokenString := "invalid.jwt.token"
    _, err = jwt.ParseWithClaims(invalidTokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err == nil {
        t.Error("ожидалась ошибка при парсинге некорректного токена, но её нет")
    } else {
        t.Logf("Ожиданная ошибка при парсинге некорректного токена: %v", err)
    }
}
