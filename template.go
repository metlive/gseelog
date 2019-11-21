package gseelog

import "fmt"

func logTemplate(logPath string, c *Config) string {
	//现在我想把日志输出到终端同时也把日志输出到文件
	logStartTemplate := `
	<seelog type="asynctimer" asyncinterval="1000000" minlevel="debug" maxlevel="error">
		<outputs formatid="%s">
			<console/>
			<filter levels="%s">
				<buffered  formatid="someformat" size="1000" flushperiod="100">
				   <rollingfile type="%s" filename="%s" %s="%s" maxrolls="%s" />
				</buffered>
			</filter>
		`
	//用于单独处理某级别日志过滤日志，把级别是error的通过邮件smtp方式发送出去(一般会发给相应的运维人员)
	var logErrorTemplate string
	if c.ErrorNotification {
		logErrorTemplateTemp := `
			<filter levels="error">
				<file path="%s"/>
				<smtp senderaddress="%s"
				sendername="Automatic notification services"
				hostname="%s"
				hostport="%s"
				username="%s"
				password="%s">
				<recipient address="%s"/>
				</smtp>
			</filter>
		`
		//格式化
		logErrorTemplate = fmt.Sprintf(logErrorTemplateTemp, logPath+"_error.log", c.Hostname, c.Hostname, c.Hostport, c.Username, c.Password, c.Address)
	}
	logEndTemplate := `
		</outputs>
		<formats>
			<format id="main" format="%%Date(2006-01-02 15:04:05) - [%%LEV] - %%Msg%%n"/>
			<format id="someformat" format="%%Date(2006-01-02 15:04:05) - [%%LEV] - %File %FullPath %RelFile %Msg%n"/>
		</formats>
	</seelog>
	`
	newLogTemplate := logStartTemplate + logErrorTemplate + logEndTemplate
	cfgRollTypeMap := map[string]string{"date": "datepattern", "size": "maxsize"}
	rollTypeParamKey, _ := cfgRollTypeMap[c.RollType]

	logConfig := fmt.Sprintf(newLogTemplate, c.FormatId, c.Levels, c.RollType, logPath+".log", rollTypeParamKey, c.RollTypeParam, c.RollTypeMaxRolls)
	return logConfig
}
