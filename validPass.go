// **Описание**: Реализуйте функцию для проверки валидности
// пароля путем сравнения с его хэшем, используя bcrypt.
// **Входные данные**: password (string) - исходный пароль для
// проверки, hashedPassword (string) - хэшированный пароль из
// базы данных
// **Выходные данные**: bool - true если пароль корректен,
// false если нет; error - ошибка при сравнении паролей
// **Ограничения**: password и hashedPassword не должны быть
// пустыми строками
// **Примеры**:
// Input: password = "mypassword123", hashedPassword =
// "$2a$10$N9qo8uLOickgx2ZMRZoMye"
// Output: true, nil
//
// Input: password = "wrongpassword", hashedPassword =
// "$2a$10$N9qo8uLOickgx2ZMRZoMye"
// Output: false, nil
package main
import(
  "golang.org/x/crypto/bcrypt"
  "errors"
)
func VerifyPassword(password string, hashedPassword string) (bool, error) {
  if password == "" || hashedPassword == ""{
    return false, errors.New("password and hash cannot be empty")
  }
 
  err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))  
  if err == nil{
    return true, nil
  } 
  if err == bcrypt.ErrMismatchedHashAndPassword{
    return false, nil
  }
  return false, err
}
