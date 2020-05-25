## filetype
golang package for file type checking, the core checker logic comes from the project, https://github.com/h2non/filetype , the main difference is this package only gets minimized file content buf for checking logic, which is friendly to big size file.

## Examples

#### Simple file type check

```go
package main

import (
    "fmt"

    "github.com/ijaa/filetype"
)

func main() {

    ok, _ := Is(types.TypeJpeg, "sample.jpg")
	if !ok {
		fmt.Println("this file is not jpg")
	}
}
```