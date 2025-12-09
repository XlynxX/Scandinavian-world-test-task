package main

import (
	"context"
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

func (a *App) Generate(passwordLength int, useNumbers, useLowercaseLetters, useUppercaseLetters bool) string {
	password, err := generatePassword(0, passwordLength, useNumbers, useLowercaseLetters, useUppercaseLetters)
	if err != nil {
		return err.Error()
	}
	return password
}
