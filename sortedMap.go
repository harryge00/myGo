package main

import (
	"fmt"
	"time"

	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
)

func compare(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func testSortedMap() {
	sm := sortedmap.New(5, compare)
	sm.Insert("type1", 7)
	sm.Insert("type2", 15)
	sm.Insert("type3", 32)
	sm.Insert("type4", 40)
	sm.Insert("type5", 50)

	iterCh, err := sm.BoundedIterCh(false, 30, 40)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer iterCh.Close()

	for rec := range iterCh.Records() {
		fmt.Printf("%+v\n", rec)
	}

}

func main() {
	testSortedMap()
	// Create an empty SortedMap with a size suggestion and a less than function:
	sm := sortedmap.New(4, asc.Time)

	// Insert example records:
	sm.Insert("OpenBSD", time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC))
	sm.Insert("UnixTime", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	sm.Insert("Linux", time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC))
	sm.Insert("GitHub", time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC))

	// Set iteration options:
	reversed := true
	lowerBound := time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC)
	upperBound := time.Now()

	// Select values > lowerBound and values <= upperBound.
	// Loop through the values, in reverse order:
	iterCh, err := sm.BoundedIterCh(reversed, lowerBound, upperBound)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer iterCh.Close()

	for rec := range iterCh.Records() {
		fmt.Printf("%+v\n", rec)
	}
}
