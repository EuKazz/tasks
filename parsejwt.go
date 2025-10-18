// **Описание**: Реализуйте функцию для декодирования и
	// парсинга JWT токена с извлечением email из payload.
	// **Входные данные**: tokenString (string) - JWT токен в
	// формате Base64(Header).Base64(Payload).Base64(Signature),
	// secret (string) - секретный ключ для верификации подписи
	// **Выходные данные**: string - email из payload токена,
	// error - ошибка при парсинге или верификации токена
	// **Ограничения**: tokenString не должен быть пустой
	// строкой, secret должен содержать минимум 8 символов
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20ifQ.signature",
	// secret = "mysecretkey123"
	// Output: "user@example.com", nil
	//
	// Input: tokenString = "", secret = "mysecretkey123"
	// Output: "", error("token cannot be empty")
