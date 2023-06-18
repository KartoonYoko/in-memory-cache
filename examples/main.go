package main

import (
	"fmt"
	"time"

	"github.com/KartoonYoko/in-memory-cache/pkg/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second)
	userId, err := cache.Get("userId")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)

	// time.Sleep(time.Second * 3)
	// cache.Delete("userId")
	
	userId, err = cache.Get("userId")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userId)
}
