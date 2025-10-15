 // **Описание**: Реализуйте структуру Repository с методами
    // для CRUD операций над сущностью User.
    //
    // **Входные данные**: db *sql.DB - подключение к базе данных
    //
    // **Выходные данные**: UserRepository - структура репозитория
    // с методами Create, GetByID, Update, Delete
    //
    // **Ограничения**: Репозиторий должен инкапсулировать работу
    // с базой данных, методы должны возвращать ошибки при
    // неудачных операциях
    //
    // **Примеры**:
    // Input: NewUserRepository(db)
    // Output: &UserRepository{db: db}
    //
    // Input: repo.Create(user)
    // Output: созданный User, nil или nil, error
package main

import (
    "database/sql"
"fmt"
)

type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserRepository struct {
   db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
      db: db,
    }
}

func (r *UserRepository) Create(user User) (*User, error) {
  //подготовка запроса
  stmt, err:= r.db.Prepare("INSERT INTO users( name, email) VALUES(?,?)")
   if err!=nil{
     return nil, err
   } 
  defer stmt.Close()
  //вставить новое значение
  res, err:= stmt.Exec(user.Name, user.Email)
  if err!=nil{
     return nil, err
  }
  //получение айди, присвоенного автоматически последней вставленной записи
  id, err:= res.LastInsertId()
  if err!=nil{
     return nil, err
  }
  user.ID = uint(id)
  return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
  //подготовка запроса
  stmt, err:= r.db.Prepare("SELECT id, name, email FROM users WHERE id = ?")
   if err!=nil{
     return nil, err
   } 
  defer stmt.Close()
  //выполняем запрос - получаем 1 ряд
  row:= stmt.QueryRow(id)
  //сканируем результат в структуру
  var user User
  err = row.Scan(&user.ID, &user.Name, &user.Email)
  if err!=nil{
    if err == sql.ErrNoRows{
      //пользователь не найден
      return nil, fmt.Errorf("user not found")
    }
    return nil, err
  }
    return &user, nil
}

func (r *UserRepository) Update(id uint, user User) (*User, error) {
   //подготовка запроса
  stmt, err:= r.db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
   if err!=nil{
     return nil, err
   } 
  defer stmt.Close() 
  res, err:= stmt.Exec(user.Name, user.Email, id)
  if err != nil {
    return nil, err
}
  rowsAffected, err:= res.RowsAffected()
   if err != nil {
    return nil, err
}
  if rowsAffected == 0 {
    return nil, fmt.Errorf("no user found with id %d", id)
}
  return &user, nil
}

func (r *UserRepository) Delete(id uint) error {
  //подготовка запроса
  stmt, err:= r.db.Prepare("DELETE FROM users WHERE id = ?")
   if err!=nil{
     return  err
   } 
  defer stmt.Close() 
  res,err := stmt.Exec(id)
    if err!=nil{
     return  err
   } 
  rowsAffected, err:= res.RowsAffected()
   if err != nil {
    return err
}
  if rowsAffected == 0 {
    return fmt.Errorf("no user found with id %d", id)
}
    return nil
}
