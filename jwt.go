// **Описание**: Реализуйте функцию для создания JWT токена
	// с заданными claims и подписанием секретным ключом.
	//
	// **Входные данные**: email (string) - email пользователя
	// для включения в payload, secret (string) - секретный ключ
	// для подписи токена
	//
	// **Выходные данные**: string - готовый JWT токен в формате
	// Base64(Header).Base64(Payload).Base64(Signature),
	// error - ошибка при создании или подписании токена
	//
	// **Ограничения**: email не должен быть пустой строкой,
	// secret должен содержать минимум 8 символов
	//
	// **Примеры**:
	// Input: email = "user@example.com", secret = "mysecretkey123"
	// Output: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", nil
	//
	// Input: email = "", secret = "mysecretkey123"
	// Output: "", error("email cannot be empty")
