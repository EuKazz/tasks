// **Описание**: Реализуйте функцию, которая парсит JWT-токен
	// и извлекает из него email пользователя для последующего
	// использования в контексте.
	//
	// **Входные данные**: tokenString (string) - JWT-токен в виде
	// строки, secretKey ([]byte) - секретный ключ для верификации
	// токена
	//
	// **Выходные данные**: структура JWTData с полем Email
	// (string) и ошибка
	//
	// **Ограничения**: tokenString не должен быть пустым,
	// secretKey должен содержать валидный ключ
	//
	// **Примеры**:
	// Input: tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
	// secretKey = []byte("secret")
	// Output: JWTData{Email: "user@example.com"}, nil
	//
	// Input: tokenString = "invalid.token.string",
	// secretKey = []byte("secret")
	// Output: JWTData{}, error
