package ElevatorStates

import (
	"../ConfigFile"
	. "../Network"
	//. "fmt"
	"time"
)

//StateChan := make(chan ConfigFile.Elev)									// Fra FSM.go til StateFileNotYetMade
//ElevatorStatesChan := make(chan map[string]*ConfigFile.Elev)		// Fra ElevatorStates til HallReqAss

func DoSomethingSmart(StateChan chan ConfigFile.Elev, ElevatorStatesChan chan map[string]*ConfigFile.Elev) {
	// tar inn ELEV struct fra StateChan -> sender "noe" på NW
	//						-> sender "noe" til HallReqAss på ElevatorStatesChan
	StateNetworkRx := make(chan map[string]*ConfigFile.Elev)
	StateNetworkTx := make(chan map[string]*ConfigFile.Elev)
	go Transmitter(ConfigFile.ElevatorStatesPort, StateNetworkTx)
	go Receiver(ConfigFile.ElevatorStatesPort, StateNetworkRx)
	transmittTimer := time.NewTicker(time.Millisecond * 50).C

	AllStates := map[string]*ConfigFile.Elev{}

	for {
		select {
		case newLocalState := <-StateChan:
			AllStates[ConfigFile.LocalID] = &newLocalState
			//Printf("states sier local: %+v\n", AllStates)
			ElevatorStatesChan <- AllStates

		case newRemoteStates := <-StateNetworkRx:
			for elevID := range newRemoteStates {
				if elevID == ConfigFile.LocalID {
					// IGNORE!, we know our own state "best"
				} else {
					AllStates[elevID] = newRemoteStates[elevID]
				}
			}
			// sende hver gang får inn eller, bare hvis sikker på oppdatering?
			//Printf("states sier remote: %+v\n", AllStates)
			ElevatorStatesChan <- AllStates

		case <- transmittTimer:
			StateNetworkTx <- AllStates
		}
		
	}
}