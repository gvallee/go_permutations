//
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
//
// See LICENSE.txt for license information
//

package permutations

import "fmt"

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
	if len(x) == 1 {
		return x[0]
	}

	var results PossibleValues
	for i := 0; i < len(x); i++ {
		perm := permute(x[1:])
		for j := 0; j < len(x[0]); j++ {
			for k := 0; k < len(perm); k++ {
				results = append(results, fmt.Sprintf("%s %s", x[0][j], perm[k]))
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
