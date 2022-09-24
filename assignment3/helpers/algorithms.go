package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func CheckWaterStatus(water int) string {

	if water <= 5 {
		return "aman"
	} else if water > 5 && water <= 7 {
		return "siaga"
	} else {
		return "bahaya"
	}

}

func CheckWindStatus(wind int) string {
	if wind <= 6 {
		return "aman"
	} else if wind >= 7 && wind <= 14 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
