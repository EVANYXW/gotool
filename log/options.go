package log

type Options struct {
	Development  bool
	LogFileDir   string
	AppName      string
	MaxSize      int //文件多大开始切分
	MaxBackups   int //保留文件个数
	MaxAge       int //文件保留最大实际
	Level        string
	CtxKey       string //通过 ctx 传递 hlog 对象
	WriteFile    bool
	WriteConsole bool
}
