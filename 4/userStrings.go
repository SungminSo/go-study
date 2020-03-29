package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("HELLO THERE"))
	f("%s\n", s.Title("tHIs will be A title!"))

	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	f("prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("prefix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("prefix: %v\n", s.HasSuffix("Mihalis", "ls"))
	f("prefix: %v\n", s.Index("Mihalis", "ha"))
	f("prefix: %v\n", s.Index("Mihalis", "Ha"))
	f("prefix: %v\n", s.Count("Mihalis", "i"))
	f("prefix: %v\n", s.Count("Mihalis", "I"))
	f("prefix: %v\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace("\t this is a line\n"))
	f("TrimLeft: %s\n", s.TrimLeft("\t this is a\t line\n", "\n\t"))
	f("TrimRight: %s\n", s.TrimRight("\t this is a\t line\n", "\n\t"))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIhalis"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("Compare: %s\n", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}