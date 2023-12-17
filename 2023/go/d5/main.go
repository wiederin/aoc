package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
  "strconv"
)

type SeedMapping struct {
  seed, fertilizer, water, light, temperature, humidity, location int
}

type SeedMappings struct {
  Mappings []SeedMapping
}

func parse(fileName string) {
  f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

  var seeds []int
  seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
    if line == "" {
      continue
    }
    if strings.HasPrefix(line, "seeds:") {
      parts := strings.Split(line, ":")
      trimmedPart := strings.TrimSpace(parts[1])
      seedNums := strings.Split(trimmedPart, " ")
      for _, num := range seedNums {
        if num == "" {
          continue
        }
        numInt, _ := strconv.Atoi(num)
        seeds = append(seeds, numInt)
      }
    }

    if seedToSoil {

    }

    if strings.HasPrefix(line, "seed-to-soil map:") {
      seedToSoil = true
    }

    if strings.HasPrefix(line, "soil-to-fertilizer map:") {
      soilToFert = true
    }

    if strings.HasPrefix(line, "fertilizer-to-water map:") {
      fertToWater = true
    }

    if strings.HasPrefix(line, "water-to-light map:") {
      waterToLight = true
    }

    if strings.HasPrefix(line, "light-to-temperature map:") {
      lightToTemp = true
    }

    if strings.HasPrefix(line, "temperature-to-humidity map: ") {
      lightToTemp = true
    }

    fmt.Println(line)
  }
  fmt.Println(seeds)
}

func part1() {
  parse("input.test")
}

func main() {
  part1()
}
