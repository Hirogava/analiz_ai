package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func AiQuestion(question string) (string, error) {
	requestBody := CompletionRequest{
        ExtraBody: map[string]interface{}{},
        Model:     modelName,
        Messages: []Message{
            {
                Role: "user",
                Content: fmt.Sprintf(
                    "Проанализируй вопрос и определи его категорию из предложенных вариантов.\n\n"+
                        "- Проблемы с авторизацией\n"+
                        "- Проблемы с оформлением заказа\n"+
                        "- Проблемы с поиском\n"+
                        "- Проблемы с отображением страниц\n"+
                        "- Технические вопросы\n"+
                        "- Другие вопросы\n"+
                        "Не пиши никакой лишней информации в ответе кроме названия категории.\n"+
                        "Вопрос:\n%s",
                    question),
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
        return "", err
    }

    category := strings.TrimSpace(completionResponse.Choices[0].Message.Content)
    
	return category, nil
}