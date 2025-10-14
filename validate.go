package main

import (
	"errors"
	"net/http"
)

func extractURLParams(r *http.Request) (map[string]string, error) {
	// **Описание**: Реализуйте функцию для извлечения и валидации
	// динамических параметров из URL-пути HTTP-запроса.
	//
	// **Входные данные**: request *http.Request - HTTP-запрос с
	// динамическими параметрами в URL
	//
	// **Выходные данные**: map[string]string - карта с извлеченными
	// параметрами, error - ошибка валидации
	//
	// **Ограничения**: URL должен содержать параметры в формате
	// /api/users/{id}/posts/{postId}, параметры не должны быть пустыми
	//
	// **Примеры**:
	// Input: URL "/api/users/123/posts/456"
	// Output: map[string]string{"id": "123", "postId": "456"}, nil
	//
	// Input: URL "/api/users//posts/456"
	// Output: nil, error("parameter 'id' cannot be empty")
  res:= make(map[string]string)
  parts:= strings.Split(r.URL.Path, "/")
    if len(parts) != 6 || parts[1] != "api" {
        return nil, errors.New("not found /api/wrong")
    }
  id:= parts[3]
  if id ==""{
    return nil, errors.New("parameter 'id' cannot be empty")
  }
  res["id"] = id
  postId:= parts[5]
  if postId ==""{
    return nil, errors.New("parameter 'postId' cannot be empty")
  }
  res["postId"] = postId
	return res, nil
}
