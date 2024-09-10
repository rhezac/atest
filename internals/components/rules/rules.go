package rules

import "slices"

func StateCheck(currentState string, newState string) (result bool) {
	switch {
	case currentState == "proposed":
		nextState := []string{"proposed", "approved", "invested", "disbursed"}
		result = slices.Contains(nextState, newState)
	case currentState == "approved":
		nextState := []string{"approved", "invested", "disbursed"}
		result = slices.Contains(nextState, newState)
	case currentState == "invested":
		nextState := []string{"invested", "disbursed"}
		result = slices.Contains(nextState, newState)
	}
	return result
}
