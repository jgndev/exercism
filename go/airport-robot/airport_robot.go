package airportrobot

import "fmt"

// Greeter interface has two methods, LanguageName and Greet for greeting airport visitors.
type Greeter interface {
	// LanguageName accepts a string for the language name and returns a string.
	LanguageName() string
	// Greet accepts a string for the visitor's name and returns a greeting as a string.
	Greet(string) string
}

// German implements Greeter for German
type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return "Hallo " + name + "!"
}

// Italian implements Greeter for Italian
type Italian struct{}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

// Portuguese implements Greeter for Portuguese
type Portuguese struct{}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
