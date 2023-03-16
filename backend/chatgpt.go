package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ChatGPTChoice struct {
	Index        int                  `json:"index"`
	Message      ChatGPTChoiceMessage `json:"message"`
	FinishReason string               `json:"finish_reason"`
}

type ChatGPTChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatGPTResponse struct {
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Created int             `json:"created"`
	Choices []ChatGPTChoice `json:"choices"`
	Usage   ChatGPTUsage    `json:"usage"`
}

func ChatGPTAPI(message string) (string, error) {
	// API endpoint와 API key를 설정
	endpoint := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("ChatGPT_API_KEY")

	var responseData ChatGPTResponse

	// API 호출에 필요한 데이터를 생성
	data := []byte(fmt.Sprintf(`{
        "model": "gpt-3.5-turbo",
    	"messages": [{"role": "user", "content": "%s"}]
    }`, message))

	// API 요청을 생성
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// API 호출 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling API:", err)
		return "", err
	}
	defer resp.Body.Close()

	// API 응답 불러오기
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	// API 응답을 JSON으로 파싱
	err = json.Unmarshal(buf.Bytes(), &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return "", err
	}

	for _, choice := range responseData.Choices {
		if choice.FinishReason == "stop" {
			return choice.Message.Content, nil
		}
	}

	// API 응답을 출력합니다.
	return "", fmt.Errorf("no response")

}
