package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(values string) string {
	var buf bytes.Buffer
	n := len(values)
	if n <= 3 {
		return values
	}
	for i, v := range values {
		if i != 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String()
}

func main() {
	s := comma1("abcdefgsefhj")
	fmt.Println(s)
}
