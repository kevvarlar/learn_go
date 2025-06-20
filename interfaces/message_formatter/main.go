package main

// Don't Touch below this line
type Formatter interface {
	Format()(formatted string)
}

type PlainText struct {
	message string
}

func (p PlainText) Format()(formatted string) {
	return p.message
}

type Bold struct {
	message string
}

func (p Bold) Format()(formatted string) {
	return "**" + p.message + "**"
}

type Code struct {
	message string
}

func (p Code) Format()(formatted string) {
	return "`" + p.message + "`"
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
