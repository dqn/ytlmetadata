# ytlmetadata

Fetch YouTube Live metadata.

## Installation

```bash
$ go get github.com/dqn/ytlmetadata
```

## Usage

Support for live streaming, scheduled live and archive. Optionally, specify the language.

```go
package main

import (
  "fmt"

  "github.com/dqn/ytlmetadata"
)

func main() {
  metadataClient := ytlmetadata.New()

  // (optional) Specify the language. default is "en".
  // m.Language = "ja"

  metadata, err := metadataClient.Fetch("VIDEO_ID")
  if err != nil {
    // Handle error.
  }

  // e.g.
  fmt.Println(metadata.ViewCount)      // 17,216 watching now
  fmt.Println(metadata.ShortViewCount) // 17K
  fmt.Println(metadata.IsLive)         // true
  fmt.Println(metadata.LikeCount)      // 5.2K
  fmt.Println(metadata.DislikeCount)   // 55
  fmt.Println(metadata.Date)           // Started streaming 78 minutes ago
  fmt.Println(metadata.Title)          // lorem ipsum
  fmt.Println(metadata.Description)    // Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore
}
```

## License

MIT
