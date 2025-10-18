// **Описание**: Реализуйте функцию для создания JWT токена
	// с заданными claims и подписанием секретным ключом.
	//
	// **Входные данные**: email (string) - email пользователя
	// для включения в payload, secret (string) - секретный ключ
	// для подписи токена
	//
	// **Выходные данные**: string - готовый JWT токен в формате
	// Base64(Header).Base64(Payload).Base64(Signature),
	// error - ошибка при создании или подписании токена
	//
	// **Ограничения**: email не должен быть пустой строкой,
	// secret должен содержать минимум 8 символов
	//
	// **Примеры**:
	// Input: email = "user@example.com", secret = "mysecretkey123"
	// Output: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", nil
	//
	// Input: email = "", secret = "mysecretkey123"
	// Output: "", error("email cannot be empty")
package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(email string, secret string) (string, error) {
  if email == ""{
     return "", errors.New("email cannot be empty")
  }
  if len(secret)<8{
    return "", errors.New("secret must be at least 8 characters long")
  }
  claims:=jwt.MapClaims{
    "email": email,
  }
  token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims )
  ss,err:= token.SignedString([]byte(secret))
  if err!=nil{
    return "", err
  }
	
	return ss, nil
}
