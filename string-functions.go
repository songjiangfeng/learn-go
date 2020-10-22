package main

import (
	p "fmt"
	s "strings"
)

func main() {
	p.Println("Contains: ", s.Contains("test", "es"))
	p.Println("Count: ", s.Count("test", "t"))
	p.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	p.Println("HasSuffix: ", s.HasSuffix("test", "st"))
	p.Println("Index: ", s.Index("test", "e"))
	p.Println("Join: ", s.Join([]string{"a", "b"}, "-"))
	p.Println("Repeat: ", s.Repeat("a", 5))
	p.Println("Replace: ", s.Replace("foo", "o", "0", -1))
	p.Println("Replace: ", s.Replace("foo", "o", "0", 1))
	p.Println("Split: ", s.Split("a-b-c-d-e", "-"))
	p.Println("ToLower: ", s.ToLower("TEST"))
	p.Println("ToUpper: ", s.ToUpper("test"))
	p.Println("Len: ", len("hello"))
	p.Println("Char: ", "hello"[1])
}
