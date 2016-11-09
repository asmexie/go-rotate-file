package file

import (
	"testing"
)

func TestWrite(t *testing.T) {
	file := Open("/tmp/test-rotate.log")
	defer file.Close()

	file.SetSuffix(SuffixDay) // 设置文件名后缀
	_, err := file.WriteString("hello world")
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteString("hello world2")
	if err != nil {
		t.Fatal(err)
	}
}
