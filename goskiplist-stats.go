package main

import (
	"fmt"
	"github.com/samsttx/goskiplist/skiplist"
	"math"
	"math/rand"
	"time"
)

const maxLoogAverageItem = 100
const maxLoogAverageSkipList = 10

func main() {
	seekTime(0.25)
	seekTime(0.50)
	seekTime(0.75)
}

func seekTime(p float64) {
	fmt.Println("p =", p)
	fmt.Println("")
	for i := 1; i <= 5; i++ {
		nbItems := int(math.Pow(10, float64(i)))
		fmt.Println(" nbItems =", nbItems)
		fmt.Println("  Insert =", int(averageTimeCreateSkipList(nbItems, 0.5).Nanoseconds()/1000))
		fmt.Println("  Seek   =", averageTimeSeekItemInRandomSkipList(nbItems, 0.5).Nanoseconds())
	}
	fmt.Println("")
}

func averageTimeSeekItemInRandomSkipList(nbItems int, p float64) time.Duration {
	s := generateRandomSkipList(nbItems, p)
	seekKey := rand.Intn(nbItems * 10)
	s.Set(seekKey, nil)
	return averageTimeSeekItem(seekKey, s)
}

func averageTimeSeekItem(key interface{}, s *skiplist.SkipList) time.Duration {
	sum := time.Duration(0)

	// Search key multiple time
	for i := 0; i < maxLoogAverageItem; i++ {
		start := time.Now()
		s.Seek(key)
		elapsed := time.Since(start)
		sum += elapsed
	}

	return sum / time.Duration(maxLoogAverageItem)
}

func averageTimeCreateSkipList(nbItems int, p float64) time.Duration {
	sum := time.Duration(0)

	// Generate Skip List multiple times
	for i := 0; i < maxLoogAverageSkipList; i++ {
		start := time.Now()
		generateRandomSkipList(nbItems, p)
		elapsed := time.Since(start)
		sum += elapsed
	}

	return sum / time.Duration(maxLoogAverageSkipList)
}

func generateRandomSkipList(nbItems int, p float64) *skiplist.SkipList {
	s := skiplist.NewIntMap(p)
	for i := 0; i < nbItems; i++ {
		s.Set(rand.Intn(nbItems*10), nil)
	}
	return s
}
