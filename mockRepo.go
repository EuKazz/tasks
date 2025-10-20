// **Описание**: Реализуйте функцию для создания mock-объекта
// UserRepository, который имитирует работу с базой данных для
// тестирования регистрации пользователей.
//
// **Входные данные**: Структура MockUserRepository с методами
// Create и FindByEmail
//
// **Выходные данные**: Реализованные методы mock-репозитория,
// возвращающие предопределенные тестовые данные
//
// **Ограничения**: Методы должны работать с моделью User,
// содержащей поля ID, Email и Password
//
// **Примеры**:
// Input: mockRepo.Create(User{Email: "test@example.com",
// Password: "hashed123"})
// Output: User{ID: 1, Email: "test@example.com",
// Password: "hashed123"}, nil
//
// Input: mockRepo.FindByEmail("test@example.com")
// Output: User{ID: 1, Email: "test@example.com",
// Password: "hashed123"}, nil
package main

import "errors"



type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type MockUserRepository struct {
	users []User
	nextID uint
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make([]User, 0),
		nextID: 1,
	}
}

func (m *MockUserRepository) Create(user User) (User, error) {
	//check if user if unique
  for _, u:= range m.users{
      if u.Email == user.Email{
        return User{}, errors.New("user with email already exists")
      }
    }
  user.ID = m.nextID
  m.nextID++
  m.users = append(m.users, user)
  return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (User, error) {
	for _, u:= range m.users{
      if u.Email == email{
        return u, nil
      }
    }
  
	return User{}, errors.New("user not found")
}
