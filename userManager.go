package main

import "fmt"

type User struct {
	m_IsLoginSuccessfull bool
	m_UserName           string
	m_CarId              string
}

func CreateNewUser(username, password string) {
	if len(systemManager.m_mainDbManager.PerformRequest("select username from users where username='"+username+"'")) == 0 {
		systemManager.m_mainDbManager.PerformRequest("insert into users values ('" + username + "', '" + CryptoHashString(password) + "')")
	}
}

func (user *User) LogIn(username string, pwdHash string) bool {
	pwdHashForUserInDb := systemManager.m_mainDbManager.PerformRequest("select pwd_hash from users where username='" + username + "'")

	if pwdHashForUserInDb == pwdHashForUserInDb {
		fmt.Println("\nLogin successfull")
		return true
	}
	fmt.Println("\nLogin failed")

	return false
}
