package main

import (
	"fmt"
	"log"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	cfg := &GrepConfig{
		after:       0,
		before:      0,
		contextRows: 0,
		count:       false,
		ignoreCase:  false,
		invert:      false,
		fixed:       false,
		strNum:      true,
		regExp:      "heh",
		filename:    "develop/dev05/test.txt",
	}
	res, err := grep(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
