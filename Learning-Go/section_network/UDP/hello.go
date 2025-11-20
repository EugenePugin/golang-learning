package main

import (
	"net/http"
	"strconv"
)

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

var counter int

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("RawQuery: ", r.URL.String()) // URL с параметрами
	// fmt.Println("DBG: counter=", counter)
	switch r.Method {
	case http.MethodGet:
		{
			// fmt.Println("Processing GET...")
			w.Write([]byte(strconv.Itoa(counter)))
		}
	case http.MethodPost:
		{
			// fmt.Println("Processing POST...")
			r.ParseForm()
			receivedCount := r.FormValue("count")
			// fmt.Println("DBG:", receivedCount, counter)
			receivedCountNum, err := strconv.Atoi(receivedCount)
			// fmt.Println("DBG:", receivedCountNum)
			if err == nil {
				// fmt.Println("DBG: it is a number")
				counter += receivedCountNum
				w.Write([]byte(strconv.Itoa(counter)))
			} else {
				// fmt.Println("DBG: it is NOT a number")
				errorMsg := "это не число"
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMsg))
			}
		}
	}
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		// fmt.Println("Ошибка запуска сервера:", err)
	}
}
