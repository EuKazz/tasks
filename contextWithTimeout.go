	// **Описание**: Реализуйте функцию, которая выполняет
	// HTTP-запрос с тайм-аутом, используя контекст для
	// ограничения времени выполнения.
	// **Входные данные**: url (строка), timeout (time.Duration)
	// **Выходные данные**: тело ответа ([]byte) и ошибка
	// **Ограничения**: timeout должен быть положительным
	// значением, url должен быть валидным HTTP URL
	// **Примеры**:
	// Input: url = "https://httpbin.org/delay/1",
	// timeout = 2*time.Second
	// Output: тело ответа, nil
	//
	// Input: url = "https://httpbin.org/delay/5",
	// timeout = 1*time.Second
	// Output: nil, context deadline exceeded
