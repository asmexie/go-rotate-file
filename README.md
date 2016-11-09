# go-rotate-file

文件操作对象，支持自动切割。

## Install 

```sh
go get -u github.com/ibbd-dev/go-rotate-file
```

## Example

```go
package main

import (
	"github.com/ibbd-dev/go-rotate-file"
)

func TestWrite(t *testing.T) {
	file := rotateFile.Open("/tmp/test-rotate.log")
	defer file.Close()

	file.SetSuffix(SuffixDay) // 设置文件名后缀
	_, err := file.WriteString("hello world")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("hello world2")
	if err != nil {
		panic(err)
	}
}
```

