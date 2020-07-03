package log

import (
	"ginLearn/global"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)
func GetLogger(logFileName string) *logrus.Logger{
	logger :=logrus.New()
	logPath := global.Vp.GetString("log.path")
	maxAge := global.Vp.GetDuration("log.maxAge")
	rotationTime := global.Vp.GetDuration("log.rotatTime")
	loglevel := global.Vp.GetString("log.level")
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge * time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime * time.Hour ), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	if loglevel == "debug"{
		logger.SetLevel(logrus.DebugLevel)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel:  os.Stderr,
		logrus.InfoLevel:  writer,
		logrus.ErrorLevel: writer,
	},
	&logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	},
	)
	logger.AddHook(lfHook)
	return logger
}
