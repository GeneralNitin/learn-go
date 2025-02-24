package main

import "fmt"

type TrackingState int32

const (
	UNSUBMITTED TrackingState = iota
	SUBMITTED
	ERROR
	IGNORED
	INVALID
)

func main() {
	// Create a variable of type TrackingState
	var state TrackingState

	// Assign a value to the variable
	state = UNSUBMITTED

	// Print the value of the variable
	println(state)

	// Assign a different value to the variable
	state = SUBMITTED
	println(state)
	println(TrackingState(0))
	println(TrackingState(1))
	println(string(state))
	println(TrackingState(1).String())

}
func (c TrackingState) String() string {
	//switch c {
	//case 0:
	//	return "UNSUBMITTED"
	//case 1:
	//	return "SUBMITTED"
	//}
	return fmt.Sprintf("TrackingState(%q)", int(c))
}
