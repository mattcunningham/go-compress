package compress

import "strconv"

const (
	NULL = "\u0000"
)

func RLE(encode string) string {
	var o string // output string
	var l int    // last index
	for i, c := range encode {
		if i != 0 && encode[i] != encode[i-1] && string(c) != NULL {
			o += strconv.Itoa(i-l) + NULL + string(encode[i-1]) + NULL // format: amount of times character seen + NULL + respective character + NULL
			l = i
		}
	}
	if l != len(encode) { // the last few characters need to be added because they'll go undetected otherwise
		o += strconv.Itoa(len(encode)-l) + NULL + string(encode[len(encode)-1]) + NULL
	}
	return o
}

func DeRLE(decode string) string {
	var (
		o string // output string
		b string // buffer
		l int
	)
	for i, c := range decode {
		if string(c) == NULL && len(b) > 0 { // this gets respective character and adds it output string
			m, _ := strconv.Atoi(b)
			for n := 0; n < m; n++ { // add to output string amount of times necessary
				o += decode[l:i]
			}
			l = i + 1 // must add one otherwise NULL will be included
			b = ""
		} else if string(c) == NULL && len(b) == 0 { // this gets amount of times character is seen
			b = decode[l:i]
			l = i + 1 // must add one otherwise NULL will be included
		}
	}
	return o
}
