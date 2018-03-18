package logger

import (
	"strings"
)

// an default cfg for seelog
const defaultConfigTemplate string = `
<seelog>
    <outputs>
        <filter levels="trace">
            <console formatid="common"/>
        </filter>
        <filter levels="debug">
            <console formatid="coloredcyan"/>
        </filter>
        <filter levels="info">
            <console formatid="coloredyellow"/>
        </filter>
        <filter levels="warn">
            <console formatid="coloredblue"/>
        </filter>
        <filter levels="error,critical">
            <splitter formatid="coloredred">
                <console/>
            </splitter>
		</filter>
        <filter formatid="common" levels="error,critical">
                <file path="{LOG_ROOT_PATH}/{LOG_NAME}_error.log"/>
        </filter>
        <file formatid="common" path="{LOG_ROOT_PATH}/{LOG_NAME}.log"/>
    </outputs>
    <formats>
        <format id="coloredblue"  format="%EscM(34)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredred"  format="%EscM(31)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredgreen"  format="%EscM(32)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredyellow"  format="%EscM(33)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredcyan"  format="%EscM(36)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredmagenta"  format="%EscM(35)[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="common"  format="[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%n"/>
        <format id="sentry"  format="%Msg%n"/>
    </formats>
</seelog>
	`

func makeLoggerCfg(path string, logName string) string {
	res := defaultConfigTemplate
	res = strings.Replace(res, "{LOG_ROOT_PATH}", path, -1)
	return strings.Replace(res, "{LOG_NAME}", logName, -1)
}
