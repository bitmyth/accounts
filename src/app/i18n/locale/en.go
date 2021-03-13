package locale

type Messages map[string]string

var En = Messages{
	NameExist:          "user name exists",
	PasswordHashFailed: "password hash failed.",
	SaveFailed:         "save failed",
}
