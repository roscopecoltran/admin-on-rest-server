package models

type Configuration struct {
	Web struct {
		Port string
	}
	Keys struct {
		Private string
		Public  string
	}
}
