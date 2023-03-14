package log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l            *Logger
	outWrite     zapcore.WriteSyncer       // IO输出
	debugConsole = zapcore.Lock(os.Stdout) // 控制台标准输出
	once         sync.Once
)

type Logger struct {
	*zap.Logger
	opts      *Options
	zapConfig zap.Config
}

func GetInstance() *Logger {
	once.Do(func() {
		opt := &Options{
			//Development: true,
			Development:  false,
			AppName:      "hlog-app",
			MaxSize:      100,
			MaxBackups:   60,
			MaxAge:       30,
			Level:        "debug",
			CtxKey:       "hlog-ctx",
			WriteFile:    true,
			WriteConsole: true,
		}
		opt.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		opt.LogFileDir += "/logs/"

		l = &Logger{opts: opt}

		l.loadCfg()
		l.initLog()

	})
	return l
}

func (l *Logger) loadCfg() {
	if l.opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
		//l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		//l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}
}

func (l *Logger) setSyncers() {
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.opts.LogFileDir + "/" + l.opts.AppName + ".log",
		MaxSize:    l.opts.MaxSize,
		MaxBackups: l.opts.MaxBackups,
		MaxAge:     l.opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	})
	fmt.Println(l.opts.LogFileDir + "" + l.opts.AppName + ".log")
	return
}

func (l *Logger) initLog() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) cores() zap.Option {
	encoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l.GetLevel()
	})
	var cores []zapcore.Core
	if l.opts.WriteFile {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(encoder, outWrite, priority),
		}...)
	}
	if l.opts.WriteConsole {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(encoder, debugConsole, priority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func (l *Logger) GetLevel() (level zapcore.Level) {
	switch strings.ToLower(l.opts.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel //默认为调试模式
	}
}

// GetLogger returns logger
func GetLogger() *Logger {
	if l == nil {
		fmt.Println("Please initialize the hlog service first")
		return nil
	}
	return l
}

func (l *Logger) GetCtx(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(l.opts.CtxKey).(*zap.Logger)
	if ok {
		return log
	}
	return l.Logger
}

func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(l.opts.CtxKey).(*zap.Logger)
	if ok {
		return log
	}
	return l.Logger
}
