package user

type User struct {
	AuthID   string `bson: "authID" 	json: "authID"`
	UserID   string `bson: "userID" 	json: "userID"`
	Email    string `bson: "name"  		json: "name"`
	Password string `bson: "password" 	json: "password"`
}

func (u *User) CheckPassword(passwordCom string) map[string]string {
	logs := make(map[string]string)
	if u.Password == passwordCom {
		logs["status"] = "1"
		logs["msg"] = ""
		logs["result"] = u.UserID
	} else {
		logs["status"] = "215"
		logs["msg"] = "wrong password"
		logs["result"] = ""
	}

	return logs
}

func (u *User) ErrorHandle(err error) map[string]string {
	var textError string
	var textStatus string
	if err.Error() == "mongo: no documents in result" {
		textError = "User not found"
		textStatus = "215"
	} else {
		textError = err.Error()
		textStatus = "415"
	}

	logs := make(map[string]string)
	logs["status"] = textStatus
	logs["msg"] = textError
	logs["result"] = ""
	return logs
}
