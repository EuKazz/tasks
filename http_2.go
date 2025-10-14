package main

import (
	"net/url"
	"strconv"
	"strings"
)

type URLComponents struct {
	Protocol string
	Host     string
	Port     string
	Path     string
}

func parseAndValidateURL(rawURL string) (URLComponents, bool) {
	// **Описание**: Реализуйте функцию для парсинга и валидации
	// URL-адресов, которая извлекает компоненты URL и проверяет
	// их корректность.
	// **Входные данные**: rawURL string - строка с URL-адресом
	// **Выходные данные**: структура URLComponents с полями
	// Protocol, Host, Port, Path и булево значение valid
	// **Ограничения**: URL должен содержать протокол (http/https),
	// хост не должен быть пустым, порт должен быть числом от 1 до
	// 65535
	// **Примеры**:
	// Input: "https://api.example.com:8080/users/profile"
	// Output: URLComponents{Protocol: "https",
	// Host: "api.example.com", Port: "8080",
	// Path: "/users/profile"}, true
	//
	// Input: "http://localhost/health"
	// Output: URLComponents{Protocol: "http", Host: "localhost",
	// Port: "80", Path: "/health"}, true
parsedURL, err:= url.Parse(rawURL)
  if err!=nil{
    return URLComponents{},false
  }
  protocol:= parsedURL.Scheme
  if protocol != "http" && protocol !="https"{
    
     return URLComponents{},false
  }
  host:= parsedURL.Hostname()
  if host == ""{
    
    return URLComponents{},false
  }
  port:= parsedURL.Port()
  if port == ""{
    if protocol == "http"{
      port = "80"
    }else if protocol =="https"{
      port ="443"
    }else{
      return URLComponents{},false
    }
  }
  portToint,err:= strconv.Atoi(port)
  if err!=nil{
    return URLComponents{},false
  }
  if portToint<1 || portToint>65535{
   
    return URLComponents{},false
  }
  finPort:= strconv.Itoa(portToint)
  res:= URLComponents{
    Protocol: protocol,
    Host: host,
    Port: finPort,
    Path: parsedURL.Path
  }
	return res, true
}
