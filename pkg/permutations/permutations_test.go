//
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
//
// See LICENSE.txt for license information
//

package permutations

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	params := make(Params)
	params["A"] = []string{"0", "1"}
	params["B"] = []string{"0", "1"}
	params["C"] = []string{"0", "1", "2"}

	permutations := Get(params)
	for idx, p := range permutations {
		fmt.Printf("Permutation %d: %s\n", idx, p)
	}
	fmt.Printf("Total number of permutations: %d\n", len(permutations))
}
