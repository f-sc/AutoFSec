package main

import "fmt"

var systemManager MainSystemManager

func main() {
	if systemManager.InitSystem() {
		fmt.Println("\nSystem init successfull")
	}

	systemManager.m_mainSysWaitGroup.Wait()
}
