# this is a fork from github.com/tuihub/go-vndb, i hardfork due to my own needs

# go-vndb

Zero dependency Go api library for [VNDB.org API v2](https://api.vndb.org/kana)

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/anispwyn/go-vndb"
)

func main() {
	client := vndb.New(
		//vndb.UseSandBox,
		vndb.WithToken("YOUR_TOKEN"),
		//vndb.WithClient(YourHttpClient),
    )
	fmt.Println(client.Stats(context.Background()))
}
```
