package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math"
	"net/http"
	"strconv"
)

type RequestBody struct {
	Statement string `json:"statement"`
}

func Calculate(statement string) interface{} {
	var calcQueue []string // rpn
	var numbStack []string // числа
	var operStack []string // операторы
	operPriority := map[string]int{
		"(": 0,
		")": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
		"V": 4,
	}

	token := ""
	for _, char := range []rune(statement) {
		if char == ' ' {
			continue
		}
		if char >= '0' && char <= '9' || char == '.' {
			token += string(char)
		} else {
			if len(token) > 0 {
				calcQueue = append(calcQueue, token)
				if len(operStack) > 2 && operStack[len(operStack)-2] == "(" && operStack[len(operStack)-1] == "-" {
					float, _ := strconv.ParseFloat(calcQueue[len(calcQueue)-1], 64)
					calcQueue = calcQueue[:len(calcQueue)-1]
					str := strconv.FormatFloat(float*-1, 'g', -1, 64)
					calcQueue = append(calcQueue, str)
					operStack = operStack[:len(operStack)-1]
				}
				token = ""
			}

			if _, ok := operPriority[string(char)]; ok {
				if string(char) == ")" {
					var oper string
					// перекладываем в очередь все операторы в скобках
					for len(operStack) > 0 {
						oper = operStack[len(operStack)-1]
						operStack = operStack[:len(operStack)-1]
						if oper == "(" {
							break
						}
						calcQueue = append(calcQueue, oper)
					}
					if oper != "(" {
						return "Unexpected \")\""
					}
				} else {
					// сортировка оператово по весу
					for len(operStack) > 0 && string(char) != "(" {
						// pop последнего оператора стека
						oper := operStack[len(operStack)-1]
						operStack = operStack[:len(operStack)-1]

						// char - текущий оператор; oper - последний в стеке
						if operPriority[string(char)] > operPriority[oper] {
							operStack = append(operStack, oper)
							break
						}
						// если вес нового меньше чем вес последнего то перекладывает последний в rpn
						if oper != "(" {
							calcQueue = append(calcQueue, oper)
						}
					}
					operStack = append(operStack, string(char))
				}
			} else {
				return fmt.Sprintf("Unexpected symbol %s", string(char))
			}
		}
	}

	if len(token) > 0 {
		calcQueue = append(calcQueue, token)
		token = ""
	}

	if len(operStack) > 0 {
		for len(operStack) > 0 {
			oper := operStack[len(operStack)-1]
			operStack = operStack[:len(operStack)-1]
			if oper == "(" {
				return "Unexpected \"(\""
			}
			calcQueue = append(calcQueue, oper)
		}
	}
	for _, token := range calcQueue {
		switch token {
		case "+":
			float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg2 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]

			float, _ = strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg1 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]

			numbStack = append(numbStack, arg2.Add(arg1).String())
		case "-":
			if len(numbStack) >= 2 {
				float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
				arg2 := decimal.NewFromFloat(float)
				numbStack = numbStack[:len(numbStack)-1]

				float, _ = strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
				arg1 := decimal.NewFromFloat(float)
				numbStack = numbStack[:len(numbStack)-1]

				numbStack = append(numbStack, arg1.Sub(arg2).String())
			} else {
				float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
				arg1 := decimal.NewFromFloat(float)
				numbStack = numbStack[:len(numbStack)-1]

				numbStack = append(numbStack, arg1.Neg().String())
			}
		case "^":
			arg2, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			numbStack = numbStack[:len(numbStack)-1]

			arg1, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)

			numbStack = numbStack[:len(numbStack)-1]

			str := strconv.FormatFloat(math.Pow(arg1, arg2), 'g', -1, 64)
			numbStack = append(numbStack, str)

		case "*":
			float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg2 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]

			float, _ = strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg1 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]

			numbStack = append(numbStack, arg2.Mul(arg1).String())
		case "V":
			float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			numbStack = numbStack[:len(numbStack)-1]
			str := strconv.FormatFloat(math.Sqrt(float), 'g', -1, 64)
			numbStack = append(numbStack, str)
		case "/":
			float, _ := strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg2 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]
			float, _ = strconv.ParseFloat(numbStack[len(numbStack)-1], 64)
			arg1 := decimal.NewFromFloat(float)
			numbStack = numbStack[:len(numbStack)-1]
			if arg2 == decimal.NewFromFloat(0) {
				return "Division by zero"
			} else {
				numbStack = append(numbStack, arg1.Div(arg2).String())
				fmt.Println(arg1.Div(arg2).String())
			}
		default:
			numbStack = append(numbStack, token)
		}
	}
	return numbStack[len(numbStack)-1]
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "РАб")
	})
	r.POST("/calculate", func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		statement := requestBody.Statement

		// Проверяем, что поле "statement" было передано
		if statement == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не передано выражение"})
			return
		}

		// Вызываем функцию calculate с переданным выражением
		result := Calculate(statement)

		// Возвращаем результат вычислений
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.Run(":8080")

}
