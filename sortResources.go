package main

import (
	"fmt"
	"sort"
)

type StandardResource struct {
	ResourceTypes int
	GPU           float64
	CPU           float64
	Memory        float64
}

type StandardResources []*StandardResource

func (r StandardResources) Len() int {
	return len(r)
}

func (r StandardResources) Less(i, j int) bool {
	if r[i].ResourceTypes > r[j].ResourceTypes {
		return true
	} else if r[i].ResourceTypes < r[j].ResourceTypes {
		return false
	}
	if r[i].GPU > r[j].GPU {
		return true
	} else if r[i].GPU < r[j].GPU {
		return false
	}

	if r[i].CPU > r[j].CPU {
		return true
	} else if r[i].CPU < r[j].CPU {
		return false
	}

	if r[i].Memory > r[j].Memory {
		return true
	} else if r[i].Memory < r[j].Memory {
		return false
	}

	return false
}

func (r StandardResources) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	resources := []*StandardResource{
		// &StandardResource{
		// 	ResourceTypes: 4,
		// 	GPU:           0,
		// 	Memory:        128000,
		// },
		// &StandardResource{
		// 	ResourceTypes: 2,
		// 	GPU:           3,
		// 	CPU:           2,
		// },
		// &StandardResource{
		// 	ResourceTypes: 2,
		// 	CPU:           16,
		// },
		// &StandardResource{
		// 	ResourceTypes: 1,
		// 	Memory:        128000,
		// },
		&StandardResource{
			ResourceTypes: 1,
			Memory:        128000,
		},
		&StandardResource{
			ResourceTypes: 2,
			GPU:           1,
			CPU:           2,
		},
		&StandardResource{
			ResourceTypes: 1,
			CPU:           16,
		},
		&StandardResource{
			ResourceTypes: 1,
			GPU:           1,
		},
		&StandardResource{
			ResourceTypes: 2,
			GPU:           0.5,
			CPU:           2,
		},
		&StandardResource{
			ResourceTypes: 1,
			CPU:           8,
		},
	}
	sort.Sort(StandardResources(resources))
	fmt.Println(resources[0], resources[1], resources[2], resources[3], resources[4], resources[5])
}
