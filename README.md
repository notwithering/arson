# Arson
[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/notwithering/arson/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/notwithering/arson)](https://goreportcard.com/report/github.com/notwithering/arson)

**Arson** simply commits arson on your string.

Be careful, if you're too devious government officials may attempt to track you down.

## Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/notwithering/arson"
)

func main() {
	str := "my life"
	go arson.Commit(&str)
	for {
		fmt.Println(str)
		time.Sleep(50 * time.Millisecond)
	}
}

```
```
my life
my l🔥fe
m🔥 l🔥fe
m🔥🔥l🔥fe
...
```