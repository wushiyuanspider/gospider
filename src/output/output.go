package output

type OutputType map[string][][]string

// 多种输出方式的统一调用格式
type Output interface {
	Output(data OutputType)
}

func newline() string {
	return "\r\n"
}