package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(write http.ResponseWriter, request *http.Request) {
	countStr := request.URL.Query().Get("count")
	if countStr == "" {
		write.WriteHeader(http.StatusBadRequest)
		if _, err := write.Write([]byte("count missing")); err != nil {
			fmt.Fprintf(write, "произошла ошибка записи ответа: %v\n", err)
		}
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		if _, err := write.Write([]byte("wrong count value")); err != nil {
			fmt.Fprintf(write, "произошла ошибка записи ответа: %v\n", err)
		}
		return
	}

	city := request.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		write.WriteHeader(http.StatusBadRequest)
		if _, err := write.Write([]byte("wrong city value")); err != nil {
			fmt.Fprintf(write, "произошла ошибка записи ответа: %v\n", err)
		}
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	write.WriteHeader(http.StatusOK)
	if _, err := write.Write([]byte(answer)); err != nil {
		fmt.Fprintf(write, "произошла ошибка записи ответа: %v\n", err)
	}
}
