// **Описание**: Реализуйте программу для создания и настройки
// функции bootstrap, которая инициализирует все зависимости
// приложения и возвращает готовый к использованию http.Handler.
// **Входные данные**: Структура конфигурации приложения с
// параметрами базы данных и секретным ключом
// **Выходные данные**: Функция bootstrap возвращающая
// http.Handler и функцию cleanup для освобождения ресурсов
// **Ограничения**: Использовать dependency injection для
// внедрения зависимостей, инициализировать базу данных через
// GORM, настроить роутинг
// **Примеры**:
// Input: config = AppConfig{DatabaseURL: "test.db",
// SecretKey: "secret123"}
// Output: http.Handler с настроенными маршрутами,
// cleanup функция, nil error
//
// Input: config = AppConfig{DatabaseURL: "", SecretKey: ""}
// Output: nil, nil, error "invalid configuration"
