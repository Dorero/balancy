package utils

import (
	"log"
	"net/url"
	"strconv"
)

type Validator struct{}

func (v Validator) IsNumbers(numbers url.Values) map[string]int {
	payload := make(map[string]int)

	for k, v := range numbers {
		number, err := strconv.Atoi(v[0])

		if err != nil {
			log.Println("Ошибка преобразования строки в число:", err)
			return nil
		} else {
			payload[k] = number
		}
	}

	return payload
}
