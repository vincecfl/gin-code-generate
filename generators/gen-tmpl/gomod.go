package tmpl

// GoModContent gomod
var GoModContent = `module {{ .ProjectName }}
go 1.14
require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.12
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/spf13/viper v1.7.0
	github.com/urfave/cli/v2 v2.2.0
	go.uber.org/zap v1.15.0
	github.com/sony/sonyflake v1.0.0
)`
