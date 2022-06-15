package rot13

// See https://en.wikipedia.org/wiki/ROT13

type Char interface {
	byte | rune | int
}

// Rot13 maps a character (byte, rune, int) to it's rot13 equivalent.
func Rot13[T Char](inp T) (ret T) {
	switch {
	case 'A' <= inp && inp <= 'Z':
		ret = (inp-'A'+13)%26 + 'A'
	case 'a' <= inp && inp <= 'z':
		ret = (inp-'a'+13)%26 + 'a'
	default:
		ret = inp
	}
	return
}
