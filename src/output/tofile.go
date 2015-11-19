package output

import (
	//"fmt"
	"os"
)
type File struct {
	file *os.File
}

func NewFile(name string) *File {
	_, err := os.Open("output")
	if err != nil {
		// 创建该目录
		os.Mkdir("output", os.ModeDir)
	}
	f := new(File)
	f.file, err = os.Create("output/" + name)
	if err != nil {
		panic("无法创建输出文件")
	}
	return f 
}

func (f *File)Output(data OutputType) {
	for k, v := range data {
		f.file.WriteString(k + newline())
		
		for _, in_v := range v {
			f.file.WriteString(in_v[1]+ newline())
		}
	}
}

func (f *File) Close() {
	f.file.Close()
}