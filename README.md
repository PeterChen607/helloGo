# helloGo
Test for go module, just say hello

## Install

import code

```bash
go get github.com/PeterChen607/hellogo@latest
```

install cmd

```bash
go install github.com/PeterChen607/hellogo@latest
```

## Example

Here's a simple example as follows:

```go
package main

import (
  "fmt"
  "github.com/PeterChen607/hellogo"
)

func main() {
  result := hello.Hello("jack")
  fmt.Println(result)
}
```