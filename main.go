package main

import (
	"bufio"
	"os"
	"unicode"
)

var homoglyphs = map[rune]rune {
    'A': 'Α',
    'B': 'Β',
    'C': 'Ϲ',
    'D': 'Ⅾ',
    'E': 'Ε',
    // 'F': 'Ϝ',
    'G': 'Ԍ',
    'H': 'Η',
    'I': 'Ι',
    'J': 'Ј',
    'K': 'Κ',
    'L': 'Ⅼ',
    'M': 'Μ',
    'N': 'Ν',
    'O': 'Ο',
    'P': 'Ρ',
    // 'Q': 'Ｑ',
    // 'R': 'Ｒ',
    'S': 'Ѕ',
    'T': 'Τ',
    // 'U': 'Ｕ',
    // 'V': 'Ⅴ',
    // 'W': '?',
    'X': 'Χ',
    'Y': 'Υ',
    'Z': 'Ζ',

    'a': 'а',
    // 'b': '?',
    'c': 'с',
    'd': 'ⅾ',
    'e': 'е',
    // 'f': '?',
    // 'g': '?',
    'h': 'һ',
    'i': 'і',
    'j': 'ј',
    // 'k': '?',
    'l': 'ⅼ',
    'm': 'ⅿ',
    // 'n': '?',
    'o': 'ο',
    'p': 'р',
    // 'q': '?',
    // 'r': '?',
    's': 'ѕ',
    // 't': '?',
    // 'u': '?',
    'v': 'ν',
    // 'w': '?',
    'x': 'х',
    'y': 'у',
    // 'z': '?',
}

func transform (char rune) rune {
    homoglyph, exists := homoglyphs[char]
    if exists {
        return homoglyph
    } else {
        return char
    }
}

func main () {
    reader := bufio.NewReader(os.Stdin)
    count := 0

    for {
        char, _, err := reader.ReadRune()
        if err != nil {
            break // EOF
        }

        os.Stdout.WriteString(string(transform(char)))

        // Insert zero-width spaces between every few characters
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            count += 1
            if count >= 3 {
                count = 0
                os.Stdout.WriteString("​")
            }
        }
    }
}
