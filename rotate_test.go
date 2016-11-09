package rotateFile

import (
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	file := Open("/tmp/test-rotate.log")
	defer file.Close()

	file.SetSuffix(SuffixDay) // 设置文件名后缀
	_, err := file.WriteString("hello world")
	if err != nil {
		t.Fatal(err)
	}

	// 模拟日期切换
	nowFunc = func() time.Time {
		return time.Now().Add(time.Hour * 25)
	}

	_, err = file.WriteString("hello world2")
	if err != nil {
		t.Fatal(err)
	}
}
