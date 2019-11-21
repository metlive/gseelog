package gseelog

import (
	"fmt"
	log "github.com/cihub/seelog"
)

type Config struct {
	Levels            string `json:"levels"`
	FormatId          string `json:"formatid"`
	RollType          string `json:"rolltype"`
	RollTypeParam     string `json:"rolltypeparam"`
	RollTypeMaxRolls  string `json:"rolltypemaxrolls"`
	ErrorNotification bool   `json:"errornotification"`
	Console           bool   `json:"console"`
	Hostname          string `json:"hostname"`
	Hostport          string `json:"hostport"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Address           string `json:"address"`
}

//根据不同服务器读取不同配置文件自定义格式
func InitLogerPool(logPath string, config *Config) error {
	logger, err := log.LoggerFromConfigAsBytes([]byte(logTemplate(logPath, config)))
	if err != nil {
		fmt.Printf("Log日志模板解析失败 %v", err)
		return err
	}
	defer logger.Flush()
	loggerErr := log.ReplaceLogger(logger)

	if loggerErr != nil {
		fmt.Println(loggerErr)
		return loggerErr
	}
	return nil
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	log.Error(v...)
}
func Errortf(format string, args ...interface{}) {
	log.Error(fmt.Sprintf(format, args...))
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	log.Warn(v...)
}

// Info compatibility alias for Info()
func Info(v ...interface{}) {
	log.Info(v...)
}

func Infotf(format string, args ...interface{}) {
	log.Info(fmt.Sprintf(format, args...))
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	log.Debug(v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(v ...interface{}) {
	log.Trace(v...)
}

func Panic(v ...interface{}) {
	log.Critical(v...)
}
