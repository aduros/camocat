package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

// From http://homoglyphs.net
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
    file := os.Stdin
    if len(os.Args) > 1 && os.Args[1] != "-" {
        var err error
        file, err = os.Open(os.Args[1])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    }

    reader := bufio.NewReader(file)
    count := 0

    for {
        char, _, err := reader.ReadRune()
        if err != nil {
            if errors.Is(err, io.EOF) {
                os.Exit(0)
            } else {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
        }

        // Insert zero-width spaces between every few characters
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            count += 1
            if count >= 3 {
                count = 0
                os.Stdout.WriteString("\ufeff")
            }
        } else {
            count = 0
        }

        os.Stdout.WriteString(string(transform(char)))
    }
}
