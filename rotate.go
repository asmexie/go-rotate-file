package rotateFile

import (
	"os"
	"sync"
	"time"
)

type TSuffix string

const (
	// 文件名后缀的格式
	SuffixHour  TSuffix = "06010215"
	SuffixDay   TSuffix = "060102"
	SuffixMonth TSuffix = "0601"

	// Log类文件的Flag
	LogFlag = os.O_WRONLY | os.O_CREATE | os.O_APPEND

	// Log Mode
	LogMode = 0666
)

type Rotate struct {
	file        *os.File
	suffixType  string
	orgFilename string

	// 文件的打开属性
	flag int
	mode os.FileMode

	// 文件名锁，保护下面几个属性
	filenameMu   sync.Mutex
	destFilename string
	destKey      string // 目标文件的key值，用以区分
}

func Open(filename string) *Rotate {
	return &Rotate{
		flag: LogFlag,
		mode: LogMode,

		orgFilename: filename,
		suffixType:  string(SuffixHour), // 默认按小时
	}
}

var nowFunc = time.Now

func (f *Rotate) SetSuffix(suffixType TSuffix) {
	f.suffixType = string(suffixType)
}

func (f *Rotate) SetFlag(flag int) {
	f.flag = flag
}

func (f *Rotate) SetMode(mode os.FileMode) {
	f.mode = mode
}

// ResetFile 重新打开文件操作对象
// 通常不需要调用该方法
// 如果Write方法写入失败，可以尝试重建File对象
func (f *Rotate) ResetFile() error {
	now := nowFunc()
	f.filenameMu.Lock()
	defer f.filenameMu.Unlock()

	key := now.Format(f.suffixType)
	name := f.orgFilename + "." + key

	// 创建新的文件对象
	file, err := os.OpenFile(name, f.flag, f.mode)
	if err != nil {
		f.filenameMu.Unlock()
		return err
	}

	// 关闭旧对象
	if f.file != nil {
		f.file.Close()
	}

	f.file, f.destFilename, f.destKey = file, name, key
	return nil
}

func (f *Rotate) Write(b []byte) (n int, err error) {
	now := nowFunc()
	f.filenameMu.Lock()
	key := now.Format(f.suffixType)
	if key != f.destKey {
		if f.file != nil {
			f.file.Close()
		}

		// 新文件名
		name := f.orgFilename + "." + key

		// 创建新的文件对象
		f.file, err = os.OpenFile(name, f.flag, f.mode)
		if err != nil {
			f.filenameMu.Unlock()
			return n, err
		}

		f.destFilename, f.destKey = name, key
	}
	f.filenameMu.Unlock()

	return f.file.Write(b)
}

func (f *Rotate) WriteString(s string) (n int, err error) {
	return f.Write([]byte(s))
}

func (f *Rotate) Close() error {
	return f.file.Close()
}
