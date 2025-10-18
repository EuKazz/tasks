// **Описание**: Реализуйте функцию для проверки срока
	// действия JWT токена и извлечения времени истечения
	// из claims.
	// **Входные данные**: tokenString (string) - JWT токен
	// для проверки, secret (string) - секретный ключ
	// для верификации
	// **Выходные данные**: bool - true если токен не истек,
	// false если истек; int64 - время истечения в Unix
	// timestamp; error - ошибка при парсинге или верификации
	// **Ограничения**: tokenString не должен быть пустой
	// строкой, secret должен содержать минимум 8 символов
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	// .eyJleHAiOjE3MDAwMDAwMDB9.signature",
	// secret = "mysecretkey123"
	// Output: false, 1700000000, nil
	//
	// Input: tokenString = "", secret = "mysecretkey123"
	// Output: false, 0, error("token cannot be empty")
package main

import (
	"errors"
	"time"
    "fmt"
	"github.com/golang-jwt/jwt/v5"
)

func checkTokenExpiration(tokenString string, secret string) (bool, int64, error) {
	if tokenString == ""{
      return false, 0, errors.New("token cannot be empty")
    }
  if len(secret)<8{
    return false, 0, errors.New("secret must be at least 8 characters long")
  }
  //карта для распаковки всех ключей из пейлоад токена
  claims:= jwt.MapClaims{}
  token,err:= jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
    return []byte(secret), nil
  })
  if err != nil {
        return false, 0, err
    }
  if !token.Valid {
        return false, 0, fmt.Errorf("invalid token")
    }
//извлекаем значение поля
  expVal, ok:= claims["exp"]
  if !ok{
    return false, 0, fmt.Errorf("exp not found in token")
  }
//поле после декодирования может быть не нужного типа -поэтому переводим его в формат int64 (unix timestamp)
  var expUnix int64
  switch v:= expVal.(type) {
  case float64:
    expUnix = int64(v)
  case  int64:
    expUnix = v
  default:
    return false, 0, fmt.Errorf("exp is not a valid number")
  }
  //получаем текущее время
  now:= jwt.TimeFunc().Unix()
  //сравниваем его со временем истечения токена
  notExp:= now<expUnix
  //есть текущее время меньше полученного - токен действителен
	return notExp, expUnix, nil
}
