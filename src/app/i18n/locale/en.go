package locale

type Messages map[string]string

var En = Messages{
	NameExist:          "register failed,user name exists",
	PasswordHashFailed: "register failed,password hash failed.",
	SaveFailed:         "register failed,save failed",
}
