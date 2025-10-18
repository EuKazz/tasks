// **Описание**: Реализуйте функцию Chain, которая принимает
// переменное количество middleware-функций и возвращает
// функцию-обёртку для создания цепочки middleware.
// **Входные данные**: middleware ...func(http.Handler) http.Handler -
// переменное количество middleware-функций
// **Выходные данные**: func(http.Handler) http.Handler -
// функция-обёртка для применения цепочки middleware
// **Ограничения**: Middleware должны применяться в обратном
// порядке (последний middleware в списке выполняется первым)
// **Примеры**:
// Input: Chain(loggingMiddleware, authMiddleware)
// Output: Функция, которая применит сначала authMiddleware,
// затем loggingMiddleware
//
// Input: Chain(corsMiddleware, rateLimitMiddleware, authMiddleware)
// Output: Функция, которая применит middleware в порядке:
// authMiddleware → rateLimitMiddleware → corsMiddleware
