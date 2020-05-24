## filetype
go package for file type checking,the core checker logic comes froms https://github.com/h2non/filetype ,the main difference is this package only get minimized file content for checking logic

## Examples

#### Simple file type check

```go
package main

import (
    "fmt"

    "github.com/ijaa/filetype"
)

func main() {

    if filetype.IsMp3("sample.mp3") {
        fmt.Println("is mp3")
    }

    if filetype.TypeIn("sample.mp3",filetype.MP3,filetype.MP4,filetype.filetype.WAV) {
        fmt.Println("in multi custom file type")
    }

    if filetype.TypeIn("sample.mp3",filetype.Audios) {
        fmt.Println("is in audio class")
    }
}
```