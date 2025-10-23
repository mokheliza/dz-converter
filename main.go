package main

import (
	"errors"
	"fmt"
	"strings"
)

const USDtoEUR float64 = 0.86
const USDtoRUB float64 = 81.06

func main() {

	for {
		userNowType, userSum, userTargetType, err := getUserInput()
		if err != nil {
			fmt.Println("Вы ввели неправильное значение ;(")

			if checkRepeatCalculation() {
				continue
			} else {
				break
			}
		}

		result := calculateCurrency(userNowType, userSum, userTargetType)

		fmt.Println(userSum, userNowType, "=", result, userTargetType)

		if !checkRepeatCalculation() {
			break
		}
	}

}

func getUserInput() (string, float64, string, error) {
	var userNowType string

	fmt.Print("Введите тип вашей валюты (rub usd eur) : ")
	fmt.Scan(&userNowType)
	userNowType = strings.ToLower(userNowType)
	if userNowType != "rub" && userNowType != "eur" && userNowType != "usd" {
		return "", 0.0, "", errors.New("Такой валюты нет!")
	}

	var userSum float64
	fmt.Print("Введите сумму вашей валюты: ")
	fmt.Scan(&userSum)

	var userTargetType string
	fmt.Print("Введите тип желаемой валюты (rub usd eur) : ")
	fmt.Scan(&userTargetType)
	userTargetType = strings.ToLower(userTargetType)
	if userTargetType != "rub" && userTargetType != "eur" && userTargetType != "usd" {
		return "", 0.0, "", errors.New("Такой валюты нет!")
	}

	return userNowType, userSum, userTargetType, nil
}

func calculateCurrency(
	userNowType string,
	userSum float64,
	userTargetType string,
) float64 {
	var result float64

	if userNowType == "usd" {
		if userTargetType == "rub" {
			result = userSum * USDtoRUB
		}
		if userTargetType == "eur" {
			result = userSum * USDtoEUR
		}
		if userTargetType == "usd" {
			result = userSum
		}
	}

	if userNowType == "eur" {
		if userTargetType == "rub" {
			result = userSum / USDtoEUR * USDtoRUB
		}
		if userTargetType == "usd" {
			result = userSum / USDtoEUR
		}
		if userTargetType == "eur" {
			result = userSum
		}
	}

	if userNowType == "rub" {
		if userTargetType == "eur" {
			result = userSum / USDtoRUB * USDtoEUR
		}
		if userTargetType == "usd" {
			result = userSum / USDtoRUB
		}
		if userTargetType == "rub" {
			result = userSum
		}
	}

	return result
}

func checkRepeatCalculation() bool {
	var userChoise string

	fmt.Print("Вы хотите сделать еще расчет? (y/n): ")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
