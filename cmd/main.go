package main

import (
	"fmt"
	"github.com/patyukin/go-memory-cache/internal/cache"
	"log"
)

func main() {
	count := 1
	sc := cache.NewCache(count)

	fmt.Println(sc)

	err := sc.Set("foo", "bar")
	err = sc.Set("bar", "baz")
	err = sc.Set("baz", "foo")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	value1, ok := sc.Get("foo")
	value2, ok := sc.Get("bar")
	value3, ok := sc.Get("baz")

	fmt.Println(value1, ok, value2, ok, value3, ok)
}
