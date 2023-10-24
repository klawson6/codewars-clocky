package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const MAX_DEGREE = 360
const DEGREE_PER_HOUR = 30

func main() {
	fmt.Println("Hello we're working and good to GO")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		deg, _ := strconv.Atoi(clearString(text))
		fmt.Println("Time is: ", DegreeToTimeInt(deg))
	}
}

func DegreeToTimeFloat64(degrees float64) string {
	degrees = math.Mod(math.Mod(degrees, MAX_DEGREE)+MAX_DEGREE, MAX_DEGREE) // Convert any given int to an angle where 0 < angle < 360

	hour := math.Round(degrees / DEGREE_PER_HOUR)
	if hour == 0 {
		hour = 12
	}
	hourStr := fmt.Sprintf("%02.f", hour)

	mins := math.Round(math.Mod(degrees, DEGREE_PER_HOUR) / DEGREE_PER_HOUR * 60)
	minsStr := fmt.Sprintf("%02.f", mins)

	return hourStr + ":" + minsStr
}

func DegreeToTimeInt(degrees int) string {
	degrees = ((degrees % MAX_DEGREE) + MAX_DEGREE) % MAX_DEGREE // Convert any given int to an angle where 0 < angle < 360

	hour := degrees / DEGREE_PER_HOUR
	if hour == 0 {
		hour = 12
	}
	hourStr := fmt.Sprintf("%02d", hour)

	mins := (degrees % DEGREE_PER_HOUR) * 2
	minsStr := fmt.Sprintf("%02d", mins)

	return hourStr + ":" + minsStr
}

var nonNumericRegex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
}
