## filetype
golang package for file type checking and just need the file path param, also this package only gets minimized file content for checking logic, which is friendly to big size file.

## Examples

#### Simple file type check

```go
package main

import (
    "fmt"

    "github.com/ijaa/filetype"
)

func main() {

    ok, _ := filetype.Is(filetype.TypeJpeg, "./sample.jpg")
	if !ok {
		fmt.Println("this file is not jpg")
    }
    
    ok, _ := filetype.IsIn([]filetype.FileType{filetype.TypeJpeg,filetype.TypePng}, "./sample.jpg")
	if !ok {
		fmt.Println("this file is not jpg or png")
	}
}
```