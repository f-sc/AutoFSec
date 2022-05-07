package AutoFSecBackend

import (
	"fmt"
	"net/http"
)

type User struct {
	m_IsLoginSuccessfull bool
	m_UserName           string
	m_CarId              string
}

func (systemManager *MainSystemManager) CreateNewUser(username, password string) {
	if len(systemManager.m_mainDbManager.PerformRequest("select username from users where username='"+username+"'")) == 0 {
		systemManager.m_mainDbManager.PerformRequest("insert into users values ('" + username + "', '" + CryptoHashString(password) + "')")
	}
}

func (systemManager *MainSystemManager) LogIn(username string, pwdHash string) bool {

	pwdHashForUserInDb := systemManager.m_mainDbManager.PerformRequest("select pwd_hash from users where username='" + username + "'")

	if pwdHashForUserInDb == pwdHash {
		fmt.Println("\nLogin successfull")
		return true
	}
	fmt.Println("\nLogin failed")

	return false
}

func (systemManager *MainSystemManager) GetVehicleInfo(r *http.Request) string {
	sessionCookie, cookieResolveError := r.Cookie("login")
	if cookieResolveError != nil {
		return "NOT FOUND"
	}
	return systemManager.m_mainDbManager.PerformRequest("select vehiclemanufacturer from users where username='" + sessionCookie.Value + "'")
}

func (systemManager *MainSystemManager) GetSecuritySystemConnectionStatus(r *http.Request) string {
	sessionCookie, cookieResolveError := r.Cookie("login")
	if cookieResolveError != nil {
		return "NOT FOUND"
	}
	return systemManager.m_mainDbManager.PerformRequest("select onlinestatus from users where username='" + sessionCookie.Value + "'")

}
