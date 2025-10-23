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
  //создает новую строку
  result:=db.Create(&user)
//возвращает ошибку вставки
	return result.Error
}

func removeData(db *gorm.DB) error {
  //удаление всех юзеров из таблицы (физически - а не мягкое удаление)
	return db.Unscoped().
      Delete(&User{}).
      Error
  //возвращает ошибку при неудачном выполнении
	
}

func setupTestDB() (*gorm.DB, func(), error) {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }
    // Автомиграция структуры таблицы
    if err := db.AutoMigrate(&User{}); err != nil {
        return nil, nil, err
    }
    // Функция очистки и закрытия
    cleanup := func() {
        // Для sqlite in-memory — просто ничего делать не надо или можно явно удалить
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }
    return db, cleanup, nil
}

func main() {
  db, cleanup, err := setupTestDB()
    if err != nil {
        t.Fatalf("ошибка инициализации БД: %v", err)
    }
  defer cleanup()
  testUser := User{Email: "test@example.com", Password: "hashed123"}
    if err := initData(db, testUser); err != nil {
        t.Fatalf("initData: %v", err)
    }

    if err := removeData(db); err != nil {
        t.Fatalf("removeData: %v", err)
    }
}
