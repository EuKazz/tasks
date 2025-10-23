// **Описание**: Реализуйте программу для создания и настройки
// тестовых данных в базе данных с последующей их очисткой
// после выполнения тестов.
// **Входные данные**: Указатель на gorm.DB подключение,
// структура User с тестовыми данными
// **Выходные данные**: Функции initData для создания тестовых
// записей и removeData для полной очистки базы данных
// **Ограничения**: Использовать GORM для операций с базой,
// обрабатывать дубликаты записей, применять Unscoped().Delete()
// для полного удаления
// **Примеры**:
// Input: db = *gorm.DB, user = User{Email: "test@example.com",
// Password: "hashed123"}
// Output: Успешное создание записи в базе, функция очистки
// готова к использованию
//
// Input: removeData(db) после выполнения тестов
// Output: База данных полностью очищена от тестовых записей
package main

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}



func initData(db *gorm.DB, user User) error {
  
  result:=db.Create(&user)

	return result.Error
}

func removeData(db *gorm.DB) error {
	return db.Unscoped().
      Delete(&User{}).
      Error
	
}

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
  db := initDb()
  user := User{Email: "test@example.com", Password: "hashed123"}
  initData(db, user)
  //
  removeData(db)
}
