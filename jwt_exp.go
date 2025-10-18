// **Описание**: Реализуйте функцию для проверки срока
	// действия JWT токена и извлечения времени истечения
	// из claims.
	// **Входные данные**: tokenString (string) - JWT токен
	// для проверки, secret (string) - секретный ключ
	// для верификации
	// **Выходные данные**: bool - true если токен не истек,
	// false если истек; int64 - время истечения в Unix
	// timestamp; error - ошибка при парсинге или верификации
	// **Ограничения**: tokenString не должен быть пустой
	// строкой, secret должен содержать минимум 8 символов
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	// .eyJleHAiOjE3MDAwMDAwMDB9.signature",
	// secret = "mysecretkey123"
	// Output: false, 1700000000, nil
	//
	// Input: tokenString = "", secret = "mysecretkey123"
	// Output: false, 0, error("token cannot be empty")
