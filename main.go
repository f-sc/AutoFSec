package main

import "fmt"

var systemManager MainSystemManager

func main() {
	if systemManager.InitSystem() {
		fmt.Println("\nSystem init successfull")
	}

	CreateNewUser("XXX", "1312")

	var user User

	user.LogIn("XXX", "vA_cdyKU-lx46Ae-X_d1N1Nk99nGidot53DRPrmkB6psU17SBkHRXC0ZNFuDMOmQ")

	systemManager.m_mainSysWaitGroup.Wait()
}
