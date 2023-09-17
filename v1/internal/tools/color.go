package tools

// Text attributes
func Bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}
func Underline(str string) string {
	return "\033[4m" + str + "\033[0m"
}
func Blink(str string) string {
	return "\033[5m" + str + "\033[0m"
}
func Reverse(str string) string {
	return "\033[7m" + str + "\033[0m"
}

// Foreground colors
func Black(str string) string {
	return "\033[30m" + str + "\033[0m"
}
func Red(str string) string {
	return "\033[31m" + str + "\033[0m"
}
func Green(str string) string {
	return "\033[32m" + str + "\033[0m"
}

func Yellow(str string) string {
	return "\033[33m" + str + "\033[0m"
}
func Blue(str string) string {
	return "\033[34m" + str + "\033[0m"
}
func Magenta(str string) string {
	return "\033[35m" + str + "\033[0m"
}
func Cyan(str string) string {
	return "\033[36m" + str + "\033[0m"
}
func White(str string) string {
	return "\033[37m" + str + "\033[0m"
}

// Background colors
func BgBlack(str string) string {
	return "\033[40m" + str + "\033[0m"
}
func BgRed(str string) string {
	return "\033[41m" + str + "\033[0m"
}
func BgGreen(str string) string {
	return "\033[42m" + str + "\033[0m"
}
func BgYellow(str string) string {
	return "\033[43m" + str + "\033[0m"
}
func BgBlue(str string) string {
	return "\033[44m" + str + "\033[0m"
}
func BgMagenta(str string) string {
	return "\033[45m" + str + "\033[0m"
}
func BgCyan(str string) string {
	return "\033[46m" + str + "\033[0m"
}
func BgWhite(str string) string {
	return "\033[47m" + str + "\033[0m"
}
