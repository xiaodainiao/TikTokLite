package log

import (
	"TikTokLite/config"
	"TikTokLite/util"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger
var sugar *zap.SugaredLogger

func InitLog() {
	var coreArr []zapcore.Core

	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		if config.Config.Level == "debug" {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		} else {
			return lev < zap.ErrorLevel && lev >= zap.InfoLevel
		}
	})
	now := util.GetCurrentTimeForString()
	path := config.GetConfig().Path.Logfile
	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path + "info" + now + ".log", //日志文件存放目录，
		MaxSize:    1,                            //文件大小限制,单位MB
		MaxBackups: 5,                            //最大保留日志文件数量
		MaxAge:     30,                           //日志文件保留天数
		Compress:   false,                        //是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path + "error" + now + ".log", //日志文件存放目录
		MaxSize:    1,                             //文件大小限制,单位MB
		MaxBackups: 5,                             //最大保留日志文件数量
		MaxAge:     30,                            //日志文件保留天数
		Compress:   false,                         //是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(1)) //zap.AddCaller()为显示文件名和行号，可省略
	sugar = log.Sugar()
}

func Infof(s string, v ...interface{}) {
	sugar.Infof(s, v...)
}

func Infow(s string, v ...interface{}) {
	sugar.Infow(s, v...)
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Debugf(s string, v ...interface{}) {
	sugar.Debugf(s, v...)
}

func Debugw(s string, v ...interface{}) {
	sugar.Debugw(s, v...)
}

func Debug(v ...interface{}) {
	sugar.Debug(v...)
}

func Errorf(s string, v ...interface{}) {
	sugar.Errorf(s, v...)
}

func Errorw(s string, v ...interface{}) {
	sugar.Errorw(s, v...)
}

func Error(v ...interface{}) {
	sugar.Error(v...)
}

func Fatalf(s string, v ...interface{}) {
	sugar.Fatalf(s, v...)
}

func Fatalw(s string, v ...interface{}) {
	sugar.Fatalw(s, v...)
}

func Fatal(v ...interface{}) {
	sugar.Error(v...)
}

func Sync() {
	log.Sync()
}
