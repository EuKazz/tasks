// **Описание**: Реализуйте CORS middleware-функцию, которая
	// устанавливает необходимые заголовки для кросс-доменных
	// запросов и обрабатывает preflight OPTIONS запросы.
	// **Входные данные**: next http.Handler - следующий
	// обработчик в цепочке
	// **Выходные данные**: http.Handler - обработчик с
	// функциональностью CORS
	// **Ограничения**: Используйте заголовки
	// Access-Control-Allow-Origin, Access-Control-Allow-Methods,
	// Access-Control-Allow-Headers
	// **Примеры**:
	// Input: OPTIONS запрос с Origin: "http://localhost:3000"
	// Output: Ответ со статусом 200 и CORS заголовками
	//
	// Input: GET запрос с Origin: "https://example.com"
	// Output: Обычный ответ с добавленными CORS заголовками
