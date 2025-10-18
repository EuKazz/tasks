// **Описание**: Реализуйте функцию для декодирования и
	// парсинга JWT токена с извлечением email из payload.
	// **Входные данные**: tokenString (string) - JWT токен в
	// формате Base64(Header).Base64(Payload).Base64(Signature),
	// secret (string) - секретный ключ для верификации подписи
	// **Выходные данные**: string - email из payload токена,
	// error - ошибка при парсинге или верификации токена
	// **Ограничения**: tokenString не должен быть пустой
	// строкой, secret должен содержать минимум 8 символов
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20ifQ.signature",
	// secret = "mysecretkey123"
	// Output: "user@example.com", nil
	//
	// Input: tokenString = "", secret = "mysecretkey123"
	// Output: "", error("token cannot be empty")
package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
  "fmt"
)

//структура для распаковки всех полей из пейлоад токена
type MyClaims struct{
  Email string `json:"email"`
  jwt.RegisteredClaims
}

func parseJWTToken(tokenString string, secret string) (string, error) {
  if tokenString == ""{
      return "", errors.New("token cannot be empty")
    }
  if len(secret)<8{
    return "", errors.New("secret must be at least 8 characters long")
  }
  claims:= &MyClaims{}
  //парсинг(в claims запишутся декодированные данные)
  token, err:= jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
  //проверка метода подписи
    if _, ok:= token.Method.(*jwt.SigningMethodHMAC);!ok{
    return nil, fmt.Errorf("unexpected signing method")
  }
    return []byte(secret), nil
  })

  if err!=nil{
    return "",err
  }

  if !token.Valid{
     return "", fmt.Errorf("invalid token")
  }
   
  return claims.Email, nil

}
