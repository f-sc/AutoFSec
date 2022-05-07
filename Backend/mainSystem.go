package AutoFSecBackend

import (
	"fmt"
	"sync"
	"time"
)

const (
	chnDO_STOP_SLEEP int = 1
)

type MainSystemManager struct {
	m_mainSysWaitGroup                sync.WaitGroup
	m_mainFuncSysManagerCommunication chan int
	m_mainDbManager                   DbManager
}

func (systemManager *MainSystemManager) SystemWaitBackground() {
	for {
		if chnDO_STOP_SLEEP == <-systemManager.m_mainFuncSysManagerCommunication {
			systemManager.m_mainSysWaitGroup.Done()
			return
		} else {
			fmt.Println("\nNew sleep started")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (systemManager *MainSystemManager) InitSystem() bool {
	fmt.Print("\nSystem initialization started")

	systemManager.m_mainSysWaitGroup.Add(1)
	go systemManager.SystemWaitBackground()

	return systemManager.m_mainDbManager.ConnectToDb()
}

func (systemManager *MainSystemManager) StopSystem() {
	systemManager.m_mainDbManager.DisconnectFromDb()

	systemManager.m_mainFuncSysManagerCommunication <- chnDO_STOP_SLEEP
}
