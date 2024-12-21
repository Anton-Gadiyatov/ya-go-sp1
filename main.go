package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"ya-go-sp1-calc/calc"
)

type RequestBody struct {
	Expression string `json:"expression"`
}

type ResponseBody struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// CalculateHandler обрабатывает запросы к эндпоинту /api/v1/calculate
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendErrorResponse(w, "Expression is not valid", http.StatusBadRequest)
		return
	}

	if !isValidExpression(reqBody.Expression) {
		sendErrorResponse(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	result, err := calc.Calc(reqBody.Expression)
	if err != nil {
		sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sendSuccessResponse(w, result)
}

// isValidExpression проверяет, что выражение содержит только разрешённые символы
func isValidExpression(expression string) bool {
	allowedChars := regexp.MustCompile(`^[0-9+\-*/(). ]+$`)
	return allowedChars.MatchString(expression)
}

// sendErrorResponse отправляет ответ с ошибкой
func sendErrorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseBody{Error: errorMessage})
}

// sendSuccessResponse отправляет успешный ответ с результатом
func sendSuccessResponse(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseBody{Result: result})
}

func main() {
	// Регистрируем обработчик для эндпоинта /api/v1/calculate
	http.HandleFunc("/api/v1/calculate", CalculateHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
