package main

import (
    "database/sql"
)

type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserRepository struct {
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
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return nil
}

func (r *UserRepository) Create(user User) (*User, error) {
    return nil, nil
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
    return nil, nil
}

func (r *UserRepository) Update(id uint, user User) (*User, error) {
    return nil, nil
}

func (r *UserRepository) Delete(id uint) error {
    return nil
}
