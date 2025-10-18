// **Описание**: Реализуйте функцию, которая создает контекст
	// с возможностью отмены и запускает горутину для
	// периодической отправки сообщений в канал до получения
	// сигнала отмены.
	// **Входные данные**: interval (time.Duration) - интервал
	// между сообщениями, message (string) - текст сообщения
	// для отправки
	// **Выходные данные**: канал для получения сообщений
	// (chan string) и функция отмены (context.CancelFunc)
	// **Ограничения**: interval должен быть положительным
	// значением, message не должен быть пустой строкой
	// **Примеры**:
	// Input: interval = 100*time.Millisecond, message = "ping"
	// Output: канал с периодическими "ping" сообщениями,
	// функция для остановки
	//
	// Input: interval = 500*time.Millisecond,
	// message = "heartbeat"
	// Output: канал с периодическими "heartbeat" сообщениями,
	// функция для остановки
package main

import (
	"context"
	"time"
)

func startPeriodicSender(interval time.Duration, message string) (chan string, context.CancelFunc) {
if interval<=0{
  return nil, nil
}
  if message == ""{
    return nil, nil
  }
  context:= context.Background()
  //создаем контекст с возможностью отмены
    ctx, cancel:= context.WithCancel(context)
 //создаем канал для передачи сообщений
  readCh:= make(chan string)
//запуск анонимной горутины
  go func(){
    //таймер, который будет слать сигнал с заданным интервалом
     ticker := time.NewTicker(interval)
     defer ticker.Stop()
    //при завершении горутины закроем канал
    defer close(readCh)
    for{
      select{
        //если контект был отменен - завершим горутины
      case  <-ctx.Done():
      return 
        //если пришел тик - отправляем сообщение в канал
    case <- ticker.C:
      readCh<-message
    }   
    }
  }()
//вернули канал для чтения сообщений и функцию отмены
	return readCh, cancel
}
