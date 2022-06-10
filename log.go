package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger *zap.Logger
)

func init(){
	logger = InitLogger()
}
func Get()*zap.Logger{
	return logger
}

func Info(s string,args ...zap.Field){
	logger.Info(s,args...)
}

func Errors(s string,args ...zap.Field){
	logger.Error(s,args...)
}


func Warn(s string,args ...zap.Field){
	logger.Warn(s,args...)
}

func Sync(){
	err :=logger.Sync()
	if err!=nil{
		logger.Error(err.Error())
	}
}

func Error(s error){
	logger.Error(s.Error())
}

func InitLogger()*zap.Logger{
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/errors.log", //日志文件存放目录
		MaxSize:    3,            //文件大小限制,单位MB
		MaxBackups: 5,            //最大保留日志文件数量
		MaxAge:     30,           //日志文件保留天数
		Compress:   false,        //是否压缩处理
	})

	logger := zap.New(zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(fileWriteSyncer,zapcore.AddSync(os.Stdout)),zap.DebugLevel),
	)

	return logger
}
