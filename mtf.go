package compress

import "strings"

var (
	Alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals   = "0123456789"
	Whitespace = " \n\t"
	Ascii      = Whitespace + Alphabet + Numerals + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`\u0004"
)

func MtF(encode string) []int {
	var o []int
	m := Ascii // list of all ascii characters
	for _, c := range encode {
		i := strings.Index(m, string(c))   // find position of character in ascii list
		o = append(o, i)                   // add respective position to slice
		m = string(m[i]) + m[:i] + m[i+1:] // move found character to front of list
	}
	return o
}

func DeMtF(decode []int) string {
	var o string
	m := Ascii // list of all ascii characters
	for _, i := range decode {
		o += string(m[i])                  // find associated character to given index
		m = string(m[i]) + m[:i] + m[i+1:] // move character in ascii list to front of list
	}
	return o
}
