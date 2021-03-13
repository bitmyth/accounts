package i18n

import (
	"github.com/bitmyth/accounts/src/app/i18n/locale"
	"github.com/bitmyth/accounts/src/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.RootPath = config.RootPath + "/../../../"
	println(config.RootPath)
	err := config.Bootstrap()
	if err != nil {
		println(err)
	}
	m.Run()
}

func TestTranslate(t *testing.T) {
	Locale = locale.En

	result := Translate(locale.NameExist)

	expect := locale.En[locale.NameExist]

	if result != expect {
		t.Error("expect:", expect, "actual:", result)
	}

	Locale = locale.ZhCN

	result = Translate(locale.NameExist)

	expect = locale.ZhCN[locale.NameExist]

	if result != expect {
		t.Error("expect:", expect, "actual:", result)
	}
}
