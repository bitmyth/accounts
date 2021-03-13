package i18n

import (
	"github.com/bitmyth/accounts/src/app/i18n/locale"
	"github.com/spf13/viper"
)

var Locale = locale.En

func init() {
	Locale = GetLocale()
}

func GetLocale() locale.Messages {
	l := viper.GetString("locale")

	switch l {
	case "en":
		return locale.En
	case "zh-CN":
		return locale.ZhCN
	default:
		return locale.ZhCN
	}

}

func Translate(messageCode string) string {
	return Locale[messageCode]
}
