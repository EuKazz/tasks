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
