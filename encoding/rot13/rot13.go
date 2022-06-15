package rot13

// See https://en.wikipedia.org/wiki/ROT13

type Char interface {
	byte | rune | int
}

// Rot13 maps a character (byte, rune, int) to it's rot13 equivalent.
func Rot13[inp Char](b inp) inp {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}
