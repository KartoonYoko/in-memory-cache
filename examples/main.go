package main

import (
	"fmt"

	"github.com/KartoonYoko/in-memory-cache/pkg/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, err := cache.Get("userId")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)

	cache.Delete("userId")
	userId, err = cache.Get("userId")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userId)
}
