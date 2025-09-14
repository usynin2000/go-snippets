package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// --- Структуры под ответ API ---
type Segment struct {
	ID               int     `json:"id"`
	Start            float64 `json:"start"`
	End              float64 `json:"end"`
	Text             string  `json:"text"`
	Tokens           []int   `json:"tokens"`
	Temperature      float64 `json:"temperature"`
	AvgLogprob       float64 `json:"avg_logprob"`
	CompressionRatio float64 `json:"compression_ratio"`
	NoSpeechProb     float64 `json:"no_speech_prob"`
	Seek             int     `json:"seek"`
}

type TranscriptionResponse struct {
	Duration float64   `json:"duration"`
	Language string    `json:"language"`
	Segments []Segment `json:"segments"`
	Task     string    `json:"task"`
	Text     string    `json:"text"`
}

// --- Функция запроса в Nexara API ---
func SendMP3ToTranscriptionService(filePath string) (*TranscriptionResponse, error) {
	url := "https://api.nexara.ru/api/v1/audio/transcriptions"
	apiKey := "nx-U87qc4Q1uPFOmJb2F1ljhhni"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("ошибка копирования файла: %w", err)
	}

	if err := writer.WriteField("response_format", "verbose_json"); err != nil {
		return nil, fmt.Errorf("ошибка добавления поля: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("ошибка закрытия writer: %w", err)
	}

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка от сервера: %d\n%s", resp.StatusCode, string(respBody))
	}

	var result TranscriptionResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	return &result, nil
}

// --- Главная функция ---
func main() {
	filePath := "test_video.mp3" // <-- Укажи путь к своему файлу
	result, err := SendMP3ToTranscriptionService(filePath)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод в консоль
	fmt.Printf("Язык: %s\n", result.Language)
	fmt.Println("\nПолный текст транскрипции:\n")
	fmt.Println(result.Text)

	// Сохраняем JSON
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	_ = os.WriteFile("transcription.json", jsonData, 0644)

	// Сохраняем чистый текст
	_ = os.WriteFile("transcription.txt", []byte(result.Text), 0644)

	fmt.Println("\nРезультаты сохранены в transcription.json и transcription.txt")
}
