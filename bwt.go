package compress

import "sort"

var (
	EOT = "\u0004"
)

func BWT(encode string) string {
	var (
		ol []string // list of the rotated input string
		o  string   // the return string
	)
	encode = encode + EOT // adding end of transmission character — most helpful for decompression
	for i := 0; i < len(encode); i++ {
		ol = append(ol, encode[i:]+encode[:i]) // rotate the input string by one character
	}
	sort.Strings(ol) // sort strings lexicographically
	for _, s := range ol {
		o += string(s[len(s)-1]) // final string containing suffix of each rotated string
	}
	return o
}

func DeBWT(decode string) string {
	var o string
	ol := make([]string, len(decode))
	for j := 0; j < len(decode); j++ {
		// in each iteration, the compressed string is added as a prefix vertically to a slice of strings
		// for example, if the compressed string is "sb^Da":
		// [0]: s
		// [1]: b
		// [2]: ^D
		// [3]: a
		// this list is sorted lexicographically, then the process is repeated for the
		// total length of the input string (in the above example, there are 4 total iterations).
		for i, c := range decode {
			ol[i] = string(c) + ol[i]
		}
		sort.Strings(ol) // sort strings lexicographically
	}
	for _, i := range ol {
		if string(i[len(i)-1]) == EOT { // there will be ONLY 1 string with the EOT character at the end
			o = i[:len(i)-1] // removing EOT character
			break
		}
	}
	return o
}
