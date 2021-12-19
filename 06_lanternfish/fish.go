package main

import (
	"strconv"
	"strings"
	"log"
	"os"
	"fmt"
	"flag"
)

var (
	inputFile    = flag.String("input", "input", "The input file")
	days = flag.Int("days", 18, "Number of days to run")
)

const (
	ReproductionInterval = 6
	FirstCycleDelay = 2
)

type LanternFish struct {
	TimeLeft int
}

func (f *LanternFish) ProgressAndReproduce() bool {
	if f.TimeLeft == 0 {
		f.TimeLeft = ReproductionInterval
		return true
	}
	f.TimeLeft--
	return false
}

func Initialize(in []int) []*LanternFish {
	ret := make([]*LanternFish, 0, len(in))
	for _, left := range in {
		ret = append(ret, &LanternFish{
			TimeLeft: left,
		})
	}
	return ret
}

func FormatFish(fish []*LanternFish) string {
	left := make([]string, 0, len(fish))
	for _, fish := range fish {
		left = append(left, strconv.Itoa(fish.TimeLeft))
	}
	return strings.Join(left, ",")
}

func Iterate(fish []*LanternFish, d, maxDays int, doLog bool) []*LanternFish {
	if d == 0 {
		if doLog {
			log.Printf("Initial state: %s", FormatFish(fish))
		}
		return fish
	}

	addFish := 0
	for _, f := range fish {
		if f.ProgressAndReproduce() {
			addFish++
		}
	}

	if addFish > 0 {
		for n := 0; n < addFish; n++ {
			fish = append(fish, &LanternFish{
				TimeLeft: ReproductionInterval + FirstCycleDelay,
			})
		}
	}

	if doLog {
		daySuffix := "day: "
		if d > 1 {
			daySuffix = "days:"
		}
		maxDaysLen := len(strconv.Itoa(maxDays))
		days := fmt.Sprintf("%d", maxDaysLen)
		log.Printf("After %"+ days +"d %s %s", d, daySuffix, FormatFish(fish))
	}
	return fish
}

func ParseInput(in string) []int {
	in = strings.TrimSpace(in)
	splitted := strings.Split(in, ",")
	ret := make([]int, len(splitted))
	for n, digit := range splitted {
		d, err := strconv.Atoi(digit)
		if err != nil {
			ret[n] = -1
			continue
		}
		ret[n] = d
	}
	return ret
}

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}

	fishConfig := ParseInput(string(input))
	fish := Initialize(fishConfig)

	for d := 0; d <= *days; d++ {
		fish = Iterate(fish, d, *days, false)
	}
	log.Printf("After %d days, the number of fish is %d", *days, len(fish))
}
