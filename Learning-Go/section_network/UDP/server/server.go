package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		var source string
		// Запрашиваем у пользователя ввод
		fmt.Print("Введите слово: ")
		_, err := fmt.Scan(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}

		// Отправляем сообщение
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// Устанавливаем таймаут на чтение ответа
		fmt.Println("Ответ:")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))

		// Чтение данных в отдельном цикле
		// Если сервер пришлет больше данных, они будут обработаны
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))

			// Сбрасываем таймаут до 700 миллисекунд после первых 1024 байт
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}
}
