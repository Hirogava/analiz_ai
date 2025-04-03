package ai

import (
    "bytes"
    "duabi/db"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strconv"
    "strings"
)

const (
    baseURL   = "https://openrouter.ai/api/v1/chat/completions"
    apiKey    = "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"
    modelName = "qwen/qwen2.5-vl-32b-instruct:free"
)

type CompletionRequest struct {
    ExtraBody map[string]interface{} `json:"extra_body"`
    Model     string                 `json:"model"`
    Messages  []Message              `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type CompletionResponse struct {
    Choices []Choice `json:"choices"`
}

type Choice struct {
    Message Message `json:"message"`
}

func AiRequest(prompt string) (string, error) {
    requestBody := CompletionRequest{
        ExtraBody: map[string]interface{}{},
        Model:     modelName,
        Messages: []Message{
            {
                Role:    "user",
                Content: prompt,
            },
        },
    }

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var completionResponse CompletionResponse
    err = json.Unmarshal(body, &completionResponse)
    if err != nil {
        return "", err
    }

    if len(completionResponse.Choices) == 0 {
        return "", fmt.Errorf("no choices in response")
    }

    response := strings.TrimSpace(completionResponse.Choices[0].Message.Content)
    fmt.Printf("AI Response: %s\n", response)
    return response, nil
}

func GetCategory(question string) (int, error) {
    category, err := AiRequest(fmt.Sprintf(
        "Проанализируй вопрос и определи его категорию из предложенных вариантов.\n\n"+
            "1. Проблемы с авторизацией\n"+
            "2. Проблемы с оформлением заказа\n"+
            "3. Проблемы с поиском\n"+
            "4. Проблемы с отображением страниц\n"+
            "5. Технические вопросы\n"+
            "6. Другие вопросы\n"+
            "Не пиши никакой лишней информации в ответе кроме номера категории, без знаков пунктуации.\n"+
            "Вопрос:\n%s",
        question))
    if err != nil {
        return 0, err
    }

    fmt.Printf("Category Response: %s\n", category)

    if category == "" {
        return 0, fmt.Errorf("empty category response from AI")
    }

    if category == "6" {
        return -1, nil
    }

    categoryId, err := strconv.Atoi(category)
    if err != nil {
        return 0, fmt.Errorf("invalid category ID: %w", err)
    }

    return categoryId, nil
}

func GetAnswer(db *db.DBManager, categoryId int, question string) (string, error) {
    questions, err := db.GetQuestions(categoryId)
    if err != nil {
        return "", err
    }

    prompt := fmt.Sprintf(
        "Найди наиболее похожий вариант на следующий вопрос:\n'%s'\n\nСписок возможных вариантов:\n\n",
        question,
    )

    for _, q := range questions {
        prompt += fmt.Sprintf("%d. '%s'\n", q.ID, q.Question)
    }

    prompt += "\nОтветь только id вопроса, без дополнительных знаков пунктуации.\nЕсли не нашел подходящего варианта, напиши -1.\n"

    answerId, err := AiRequest(prompt)
    if err != nil {
        return "", err
    }

    fmt.Printf("Answer ID Response: %s\n", answerId)

    if answerId == "-1" {
        return "Ваш вопрос не найден, пожалуйста, напишите в техподдержку", nil
    }

    answerIntId, err := strconv.Atoi(answerId)
    if err != nil {
        return "", fmt.Errorf("invalid answer ID: %w", err)
    }

    answer, err := db.GetAnswer(answerIntId)
    if err != nil {
        return "", err
    }

    return answer, nil
}