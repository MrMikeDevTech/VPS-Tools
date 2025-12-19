package utils

var cRed = "\033[31m"
var cGreen = "\033[32m"
var cBlue = "\033[34m"
var cCyan = "\033[36m"
var cYellow = "\033[33m"
var cOrange = "\033[38;5;208m"
var cPink = "\033[38;5;213m"
var cPurple = "\033[35m"
var cMagenta = "\033[35m"
var cWhite = "\033[37m"
var cBlack = "\033[30m"
var cGray = "\033[90m"

var cReset = "\033[0m"
var cBold = "\033[1m"

var Colors = struct {
	Red, Green, Blue, Cyan, Yellow, Orange, Pink, Purple, Magenta, White, Black, Gray, Reset, Bold string
}{
	Red:     cRed,
	Green:   cGreen,
	Blue:    cBlue,
	Cyan:    cCyan,
	Yellow:  cYellow,
	Orange:  cOrange,
	Pink:    cPink,
	Purple:  cPurple,
	Magenta: cMagenta,
	White:   cWhite,
	Black:   cBlack,
	Gray:    cGray,
	Reset:   cReset,
	Bold:    cBold,
}
