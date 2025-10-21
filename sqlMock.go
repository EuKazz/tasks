// **Описание**: Реализуйте функцию для создания и настройки
	// mock SQL-соединения с использованием go-sqlmock для
	// тестирования операций с базой данных.
	// **Входные данные**: Структура для хранения mock базы данных
	// и ожидаемых запросов
	// **Выходные данные**: Настроенное GORM подключение с mock
	// базой, функция cleanup для закрытия соединения
	// **Ограничения**: Использовать go-sqlmock, настроить
	// ExpectQuery для SELECT запроса, возвращать тестовые данные
	// через NewRows
	// **Примеры**:
	// Input: setupMockDB() с ожидаемым запросом
	// "SELECT * FROM users WHERE email = ?"
	// Output: *gorm.DB с mock соединением, cleanup функция,
	// nil error
	//
	// Input: mockDB.ExpectQuery("SELECT").WillReturnRows(rows)
	// Output: Настроенный mock с предопределенными результатами
package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, func(), error) {
	//создает вирт бд и мок-объект ожидаемых запросов
	db, mock, err:= sqlmock.New()
  if err != nil {
        return nil, nil, nil, err
    }
   //открывает горм с драйвером- но вместо драйвера мок-соед 
  dbGorm, err:= gorm.Open(mysql.New(mysql.Config{Conn: db}), gorm.Config{})
  if err!=nil{
    return nil, nil, nil, err
  }
  //создает макет ответа бд - таблица с определенными колонками
  rows:= sqlmock.NewRows([]string{"id", "email", "password"}).
    AddRow(1, "test@g.ru", "vb34J_9ds")
  //описан ожидаемый запрос с опр аргументом
    mock.ExpectQuery("SELECT * FROM users WHERE email = ?").
    WithArgs("test@g.ru").
    WillReturnRows(rows)
   //получено сырое соед для управ
  sqlDB, err:= db.DB()
  if err!=nil{
    return nil, nil, nil, err
  }
  //функция очистки соединения
  cleanup:= func(){
    sqlDB.Close()
  }
	return dbGorm, mock, cleanup, nil
}
