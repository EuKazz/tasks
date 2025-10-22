// **Описание**: Реализуйте программу для создания и настройки
// тестового HTTP-сервера с использованием httptest.NewRecorder
// для тестирования HTTP хендлеров.
//
// **Входные данные**: HTTP хендлер функция, HTTP метод
// (GET/POST), URL путь, тело запроса (JSON)
//
// **Выходные данные**: Функция тестирования, которая создает
// запрос через httptest.NewRequest, выполняет его и проверяет
// статус-код и содержимое ответа
//
// **Ограничения**: Использовать httptest.NewRequest и
// httptest.NewRecorder, проверить статус-код и JSON ответ
//
// **Примеры**:
// Input: handler = getUserHandler, method = "GET",
// path = "/user/123", body = nil
// Output: HTTP 200 OK с JSON {"id":123,"name":"John Doe"}
//
// Input: handler = createUserHandler, method = "POST",
// path = "/user", body = {"name":"Jane"}
// Output: HTTP 201 Created с JSON
// {"message":"User created successfully"}
