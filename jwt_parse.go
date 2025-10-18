// **Описание**: Реализуйте функцию, которая парсит JWT-токен
	// и извлекает из него email пользователя для последующего
	// использования в контексте.
	//
	// **Входные данные**: tokenString (string) - JWT-токен в виде
	// строки, secretKey ([]byte) - секретный ключ для верификации
	// токена
	//
	// **Выходные данные**: структура JWTData с полем Email
	// (string) и ошибка
	//
	// **Ограничения**: tokenString не должен быть пустым,
	// secretKey должен содержать валидный ключ
	//
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
	// secretKey = []byte("secret")
	// Output: JWTData{Email: "user@example.com"}, nil
	//
	// Input: tokenString = "invalid.token.string",
	// secretKey = []byte("secret")
	// Output: JWTData{}, error
package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
  
)

type JWTData struct {
	Email string
}

type MyClaims struct {
    Email string `json:"email"`
    jwt.RegisteredClaims
}

func Parse(tokenString string, secretKey []byte) (*JWTData, error) {
if tokenString==""{
  return nil, errors.New("token cannot be empty")
}
   if len(secretKey) == 0 {
        return nil, errors.New("secret key is empty")
    }
 claims:= &MyClaims{}
token, err:= jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
  //проверка метода подписи
    if _, ok:= token.Method.(*jwt.SigningMethodHMAC);!ok{
    return nil, fmt.Errorf("unexpected signing method")
  }
    return []byte(secretKey), nil
  })

  if err!=nil{
    return nil,err
  }

  if !token.Valid{
     return nil, fmt.Errorf("invalid token")
  }
 
 if claims.Email == ""{
    return nil, errors.New("email not found in token claims")
 }
 return &JWTData{Email: claims.Email}, nil
}
