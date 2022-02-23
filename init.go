package hrp

import (
	"github.com/rs/zerolog"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel) // 设置全局日志级别
}
