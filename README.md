# blackpx
A tool to count true-black pixels in png, jpg and gif images.

## Usage
### Within code
Download the library:
```
go get -d github.com/gstvr/blackpx
```

Call it from your code:
```go
pakage main
import (
    "image"
    "github.com/gstvr/blackpx/pkg/blackpx"
)

func main() {
    ...
    
    img, _, err := image.Decode(...)
    panic(err)
    
    pixels := blackpx.CountTrueBlack(img)
}
```

The library comes with no image drivers imported, so you might need to import them manually depending on your needs, by adding one of the following import statements:
``` go
import _ "image/png"
import _ "image/jpg"
import _ "image/gif"
```

### From CLI
Install the CLI tool:
```
go install github.com/gstvr/blackpx@latest
```

Call it by providing path to an image:
```
$ blackpx path/to/image.png
25783
```
