//
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
//
// See LICENSE.txt for license information
//

package permutations

import (
	"fmt"
)

type PossibleValues []string
type ParamID string
type Params map[ParamID]PossibleValues
type Permutations PossibleValues

func getAllPossibleParamValues(id ParamID, possibleValues PossibleValues) PossibleValues {
	var results PossibleValues
	for _, v := range possibleValues {
		results = append(results, fmt.Sprintf("%s=%s", id, v))
	}
	return results
}

func permute(x []PossibleValues) PossibleValues {
	var results PossibleValues
	switch len(x) {
	case 1:
		for _, v := range x {
			results = append(results, v...)
		}
	case 2:
		for i := 0; i < len(x[0]); i++ {
			for j := 0; j < len(x[1]); j++ {
				results = append(results, fmt.Sprintf("%s %s", x[0][i], x[1][j]))
			}
		}
	default:
		subResults := permute(x[1:])
		for i := 0; i < len(x[0]); i++ {
			for j := 0; j < len(subResults); j++ {
				results = append(results, fmt.Sprintf("%s %s", x[0][i], subResults[j]))
			}
		}
	}
	return results
}

func Get(params Params) Permutations {
	// Create the list of parameter IDs
	var paramsList []ParamID
	for id := range params {
		paramsList = append(paramsList, id)
	}

	var listSupportedValues []PossibleValues
	for i := 0; i < len(paramsList); i++ {
		currentID := paramsList[i]
		paramValues := getAllPossibleParamValues(currentID, params[currentID])
		listSupportedValues = append(listSupportedValues, paramValues)
	}

	return Permutations(permute(listSupportedValues))
}
