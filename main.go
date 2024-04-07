// Arson is a package that simply commits arson on your string.
package arson

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

// Don't let this reach 5 stars or else you might get arrested.
var WantedLevel int = 0

// Function Commit commits a felony on your string.
//
// WARNING: THIS WILL TURN YOUR PROGRAM ARSONOUS
func Commit(s *string) {
	for {
		len := utf8.RuneCountInString(*s)
		i := rand.Intn(len)

		var str = strings.Builder{}
		for si := range len {
			if si == i {
				str.WriteRune('🔥')
				continue
			}
			ds := *s
			r, _ := getRuneAtIndex(ds, si)
			str.WriteRune(r)
		}
		*s = str.String()

		var test string
		for range len {
			test += "🔥"
		}
		if *s == test {
			// the entire string is engulfed in flames
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	WantedLevel += 1

	if WantedLevel >= 5 {
		panic("The CIA tracked down the GPS on your phone. You were arrested.")
	}
}

// Function Evade has your program go into hiding from law enforcement.
// This will reduce your wanted level by half.
func Evade() {
	WantedLevel /= 2
}

func getRuneAtIndex(s string, index int) (rune, int) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if i+size > index {
			return r, i
		}
		i += size
	}
	return utf8.RuneError, -1
}
