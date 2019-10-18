package main

func StrRev(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		as += string(s[i])
	}
	return as
}
