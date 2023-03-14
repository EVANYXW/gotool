package zlog

// 安装以下依赖库
// go get -u go.uber.org/zap
// go get -u github.com/natefinch/lumberjack
// go get gopkg.in/alecthomas/kingpin.v2
import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"strconv"
)

const (
	defaultlogpath           = "../log/test" // 【默认】日志文件路径
	defaultlogfilename       = "test.log"    // 【默认】日志文件名称
	defaultloglevel          = "info"        // 【默认】日志打印级别 debug  info  warning  error
	defaultlogfilemaxsize    = 5             // 【日志分割】  【默认】单个日志文件最多存储量 单位(mb)
	defaultlogfilemaxbackups = 10            // 【日志分割】  【默认】日志备份文件最多数量
	logmaxage                = 1000          // 【默认】日志保留时间，单位: 天 (day)
	logcompress              = false         // 【默认】是否压缩日志
	logstdout                = false         // 【默认】是否输出到控制台
)

var Logger *zap.SugaredLogger // 定义日志打印全局变量

var (
	// kingpin 可以在启动时通过输入参数，来修改日志参数
	level = kingpin.Flag("log.level", "only log messages with the given severity or above. one of: [debug, info, warn, error]").Default(
		defaultloglevel,
	).String()
	format            = kingpin.Flag("log.format", "output format of log messages. one of: [logfmt, json]").Default("logfmt").String()
	logpath           = kingpin.Flag("log.path", "output log path").Default(defaultlogpath).String()
	logfilename       = kingpin.Flag("log.filename", "output log filename").Default(defaultlogfilename).String()
	logfilemaxsize    = kingpin.Flag("log.file-max-size", "output logfile max size, unit mb").Default(strconv.Itoa(defaultlogfilemaxsize)).Int()
	logfilemaxbackups = kingpin.Flag("log.file-max-backups", "output logfile max backups").Default(strconv.Itoa(defaultlogfilemaxbackups)).Int()
)

// 初始化 logger
func InitLogger() error {
	if Logger != nil {
		return nil
	}

	loglevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writesyncer, err := getlogwriter() // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getencoder()       // 获取日志输出编码
	level, ok := loglevel[*level] // 日志打印级别
	if !ok {
		level = loglevel["info"]
	}
	core := zapcore.NewCore(encoder, writesyncer, level)
	logger := zap.New(core, zap.AddCaller()) //  zap.addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	Logger = logger.Sugar()
	return nil
}

// 编码器(如何写入日志)
func getencoder() zapcore.Encoder {
	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder   // looger 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderconfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 info debug error
	//encoderconfig.encodecaller = zapcore.fullcallerencoder
	//encoderconfig.encodelevel = zapcore.capitalcolorlevelencoder
	if *format == "json" {
		return zapcore.NewJSONEncoder(encoderconfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderconfig) // 以logfmt格式写入
}

// 获取日志输出方式  日志文件 控制台
func getlogwriter() (zapcore.WriteSyncer, error) {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := isexist(*logpath); !exist {
		if *logpath == "" {
			*logpath = defaultlogpath
		}
		if err := os.MkdirAll(*logpath, os.ModePerm); err != nil {
			*logpath = defaultlogpath
			if err := os.MkdirAll(*logpath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	// 日志文件 与 日志切割 配置
	lumberjacklogger := &lumberjack.Logger{
		Filename:   filepath.Join(*logpath, *logfilename), // 日志文件路径
		MaxSize:    *logfilemaxsize,                       // 单个日志文件最大多少 mb
		MaxBackups: *logfilemaxbackups,                    // 日志备份数量
		MaxAge:     logmaxage,                             // 日志最长保留时间
		Compress:   logcompress,                           // 是否压缩日志
	}
	if logstdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjacklogger), zapcore.AddSync(os.Stdout)), nil
	} else {
		// 日志只输出到控制台
		return zapcore.AddSync(lumberjacklogger), nil
	}
}

// 判断文件或者目录是否存在
func isexist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
