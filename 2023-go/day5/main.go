package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := "input.txt"

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n\n")
	almanach := parseInput(input)

	// min := solveOne(almanach)
	min := solveTwo(almanach)
	fmt.Println(min)
}

func parseNums(s string) []int {
	tmp := strings.Split(s, " ")
	nums := make([]int, len(tmp))

	for i, num := range tmp {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

type almanach struct {
	seeds         []seedRange
	toSoil        []mapping
	toFertilizer  []mapping
	toWater       []mapping
	toLight       []mapping
	toTemperature []mapping
	toHumidity    []mapping
	toLocation    []mapping
}

type mapping struct {
	dest int
	src  int
	rng  int
}

func parseInput(s []string) almanach {
	almanach := almanach{}

	// seeds
	almanach.seeds = parseSeeds(parseNums(strings.Split(s[0], ": ")[1]))

	// seed to soil
	almanach.toSoil = parseMapping(s[1])

	// soil to fertelizer
	almanach.toFertilizer = parseMapping(s[2])

	// fertilizer to water
	almanach.toWater = parseMapping(s[3])

	// water to light
	almanach.toLight = parseMapping(s[4])

	// light to temperature
	almanach.toTemperature = parseMapping(s[5])

	// temperature to humidity
	almanach.toHumidity = parseMapping(s[6])

	// humidity to location
	almanach.toLocation = parseMapping(s[7])

	return almanach
}

type seedRange struct {
	start int
	end   int
}

func containsSeed(s []seedRange, n int) bool {
	for _, seed := range s {
		if n >= seed.start && n <= seed.end {
			return true
		}
	}
	return false
}

func parseSeeds(nums []int) []seedRange {
	seeds := []seedRange{}
	for i := 0; i < len(nums)-1; i += 2 {
		start := nums[i]
		rng := nums[i+1]

		seeds = append(seeds, seedRange{
			start: start,
			end:   start + rng - 1,
		})
	}
	return seeds
}

func parseMapping(lines string) []mapping {
	mappings := []mapping{}
	tmp := strings.Split(lines, "\n")

	for i := 1; i < len(tmp); i++ {
		nums := parseNums(tmp[i])
		m := mapping{
			dest: nums[0],
			src:  nums[1],
			rng:  nums[2],
		}
		mappings = append(mappings, m)
	}
	return mappings
}

func resolveMapping(source int, mappings []mapping) int {
	for _, mapping := range mappings {
		if source >= mapping.src && source < mapping.src+mapping.rng {
			offset := source - mapping.src
			return mapping.dest + offset
		}
	}
	return source
}

func reverseMapping(dest int, mappings []mapping) int {
	for _, mapping := range mappings {
		if dest >= mapping.dest && dest < mapping.dest+mapping.rng {
			offset := dest - mapping.dest
			return mapping.src + offset
		}
	}
	return dest
}

func solveTwo(a almanach) int {
	for location := 0; location < 100_000_000; location++ {
		hum := reverseMapping(location, a.toLocation)
		temp := reverseMapping(hum, a.toHumidity)
		light := reverseMapping(temp, a.toTemperature)
		water := reverseMapping(light, a.toLight)
		fert := reverseMapping(water, a.toWater)
		soil := reverseMapping(fert, a.toFertilizer)
		seed := reverseMapping(soil, a.toSoil)
		if containsSeed(a.seeds, seed) {
			return location
		}
	}
	return -1
}
