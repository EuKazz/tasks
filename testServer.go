	// **Описание**: Реализуйте функцию для создания тестового
	// HTTP-сервера и выполнения end-to-end теста регистрации
	// пользователя через все слои приложения.
	// **Входные данные**: Функция App() возвращающая http.Handler,
	// структура User для тестовых данных
	// **Выходные данные**: Функция TestUserRegistrationE2E,
	// которая создает тестовый сервер, отправляет POST-запрос
	// и проверяет корректность ответа
	// **Ограничения**: Использовать httptest.NewServer, http.Post
	// для отправки JSON, проверить статус-код 201
	// **Примеры**:
	// Input: POST /register с JSON
	// {"email":"newuser@test.com","password":"password123"}
	// Output: HTTP 201 Created с JSON
	// {"message":"User registered successfully"}
	//
	// Input: POST /register с JSON
	// {"email":"existing@test.com","password":"pass456"}
	// Output: HTTP 400 Bad Request с JSON
	// {"error":"User already exists"}

	// TODO: Создать тестовый HTTP-сервер
	// TODO: Подготовить JSON данные для регистрации
	// TODO: Отправить POST-запрос
	// TODO: Проверить статус-код и тело ответа
	// TODO: Очистить тестовое окружение
