package main

import (
	"changeme/backend"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ChatGPT 검색 결과 API
func (a *App) ChatGPT(message string) string {
	respTxt, err := backend.ChatGPTAPI(message)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("%s", respTxt)
}

// Google 번역기 API
func (a *App) GoogleTranslate(targetLanguage, text string) string {
	respTxt, err := backend.GoogleTranslateAPI(targetLanguage, text)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("%s", respTxt)
}
