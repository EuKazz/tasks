// **Описание**: Реализуйте функцию для создания и настройки
    // тестовой базы данных с выполнением миграций и инициализацией
    // подключения.
    //
    // **Входные данные**: Строка с именем тестовой базы данных,
    // структура конфигурации БД
    //
    // **Выходные данные**: Указатель на gorm.DB с настроенным
    // подключением к тестовой базе, функция для закрытия соединения
    //
    // **Ограничения**: Использовать GORM, поддерживать SQLite для
    // тестов, обрабатывать ошибки подключения
    //
    // **Примеры**:
    // Input: dbName = "test_app.db", config = DatabaseConfig{Driver: "sqlite"}
    // Output: *gorm.DB подключение, cleanup функция, nil error
    //
    // Input: dbName = "", config = DatabaseConfig{}
    // Output: nil, nil, error "database name is required"
package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
  "errors"
)

type DatabaseConfig struct {
    Driver string
    Host   string
    Port   int
    Name   string
}

type User struct {
    ID    int
    Email string
}

func setupTestDatabase(dbName string, config DatabaseConfig) (*gorm.DB, func(), error) {
if dbName==""{
  return nil, nil, errors.New("database name is required")
}
  if config.Driver != "sqlite"{
    return nil, nil, errors.New("driver is wrong")
  }
  db, err:= gorm.Open(sqlite.Open(dbName), &gorm.Config{})
if err!=nil{
  return nil, nil, err
}  
  //низкоуровневое соединений
  sqlDB, err:= db.DB()
  if err!=nil{
    return nil, nil, err
  }
  //функция очистки соединения
  cleanup:= func(){
    sqlDB.Close()
  }
  db.AutoMigrate(&User{})
    return db, cleanup, nil
}
