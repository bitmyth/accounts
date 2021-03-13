package register

import "testing"

func TestTranslate(t *testing.T) {
	t.Log(NewNameExistsError().Error())
}
