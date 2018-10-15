package log

import (
	"time"
	"path"
	"strings"
	log "github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var (
	Logger = log.New()
	ErrLogger = log.New()

	// fastHttp日志
	FHLogger = log.New()

	logPath  = "F:/data/log-collect/logs/"
	logLevel = "debug"
)

/*
 * @see https://github.com/rifflock/lfshook
 */
func init() {
	newLog(Logger, logPath, "log.log", logLevel)
	newLog(ErrLogger, logPath, "error.log", logLevel)
	newLog(FHLogger, logPath, "fasthttp.log", logLevel)
}

func newLog(logger *log.Logger, logPath string, logName string, logLevel string) {
	if !strings.HasSuffix(logPath, "/") {
		logPath += "/"
	}
	baseLogPath := path.Join(logPath, logName)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}

	//日志文件设置
	lfHook := lfshook.NewHook(lfshook.WriterMap {
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{})

	logger.Hooks.Add(lfHook)
}

func getLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "panic":
		return log.PanicLevel
	case "fatal":
		return log.FatalLevel
	case "error":
		return log.ErrorLevel
	case "warn", "warning":
		return log.WarnLevel
	case "info":
		return log.InfoLevel
	case "debug":
		return log.DebugLevel
	default:
		return log.DebugLevel
	}
}