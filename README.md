Go Cache
================================


Cache is the simple in-memmory cache library


## Example #1

```go
package main

import (
	"fmt"
	"github.com/dameerv/golangCache.git"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")

	fmt.Println(userId)
}

