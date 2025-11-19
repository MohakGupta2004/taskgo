package ui

import (
	"strings"
)

var bigDigits = map[rune][]string{
	'0': {
		"  ###  ",
		" #   # ",
		" #   # ",
		" #   # ",
		"  ###  ",
	},
	'1': {
		"   #   ",
		"  ##   ",
		"   #   ",
		"   #   ",
		"  ###  ",
	},
	'2': {
		"  ###  ",
		" #   # ",
		"    #  ",
		"   #   ",
		" ##### ",
	},
	'3': {
		"  ###  ",
		"     # ",
		"  ###  ",
		"     # ",
		"  ###  ",
	},
	'4': {
		" #   # ",
		" #   # ",
		" ##### ",
		"     # ",
		"     # ",
	},
	'5': {
		" ##### ",
		" #     ",
		" ####  ",
		"     # ",
		" ####  ",
	},
	'6': {
		"  ###  ",
		" #     ",
		" ####  ",
		" #   # ",
		"  ###  ",
	},
	'7': {
		" ##### ",
		"    #  ",
		"   #   ",
		"  #    ",
		" #     ",
	},
	'8': {
		"  ###  ",
		" #   # ",
		"  ###  ",
		" #   # ",
		"  ###  ",
	},
	'9': {
		"  ###  ",
		" #   # ",
		"  #### ",
		"     # ",
		"  ###  ",
	},
	':': {
		"       ",
		"   #   ",
		"       ",
		"   #   ",
		"       ",
	},
}

func RenderBigText(text string) string {
	var lines [5]string
	for _, char := range text {
		digit, ok := bigDigits[char]
		if !ok {
			// Fallback for unknown chars (like space)
			for i := 0; i < 5; i++ {
				lines[i] += "       "
			}
			continue
		}
		for i := 0; i < 5; i++ {
			lines[i] += digit[i] + " "
		}
	}
	return strings.Join(lines[:], "\n")
}
