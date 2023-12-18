package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
  "strconv"
)

type Mapping struct {
  sourceStart int
  destinationStart int
  rangeLen int 
}


type Mappings struct {
  seedToSoil []Mapping
  soilToFert []Mapping
  fertToWater []Mapping
  waterToLight []Mapping 
  lightToTemperature []Mapping 
  temperatureToHumidity []Mapping
  humidityToLocation []Mapping 
}

func parse(fileName string) (*Mappings, []int, error) {
  fmt.Println("--- begin parsing ---")
  f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

  mappings := &Mappings{ 
	  seedToSoil: []Mapping{},
    soilToFert: []Mapping{},
    fertToWater: []Mapping{},
	  waterToLight: []Mapping{},
    lightToTemperature: []Mapping{},
    temperatureToHumidity: []Mapping{},
    humidityToLocation: []Mapping{},
  }
  
  var seeds []int
  seedToSoil := false
  soilToFert := false
  fertToWater := false
  waterToLight := false
  lightToTemperature := false
  temperatureToHumidity := false
  humidityToLocation := false

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
      continue
    }

    if strings.HasPrefix(line, "seed-to-soil map:") {
      seedToSoil = true
      soilToFert = false 
      fertToWater = false 
      waterToLight = false
      lightToTemperature = false
      temperatureToHumidity = false
      humidityToLocation = false
      continue
    }

    if strings.HasPrefix(line, "soil-to-fertilizer map:") {
      soilToFert = true 
      seedToSoil = false 
      fertToWater = false
      waterToLight = false
      lightToTemperature = false
      temperatureToHumidity = false
      humidityToLocation = false
      continue
    }

    if strings.HasPrefix(line, "fertilizer-to-water map:") {
      fertToWater = true 
      soilToFert = false 
      seedToSoil = false 
      waterToLight = false
      lightToTemperature = false
      temperatureToHumidity = false
      humidityToLocation = false
      continue
    }
    
    if strings.HasPrefix(line, "water-to-light map:") {
      waterToLight = true 
      fertToWater = false 
      soilToFert = false 
      seedToSoil = false 
      lightToTemperature = false
      temperatureToHumidity = false
      humidityToLocation = false
      continue
    }
    
    if strings.HasPrefix(line, "light-to-temperature map:") {
      lightToTemperature = true 
      waterToLight = false
      fertToWater = false 
      soilToFert = false 
      seedToSoil = false 
      temperatureToHumidity = false
      humidityToLocation = false
      continue
    }

    if strings.HasPrefix(line, "temperature-to-humidity map:") {
      temperatureToHumidity = true 
      lightToTemperature = false 
      waterToLight = false 
      fertToWater = false 
      soilToFert = false 
      seedToSoil = false 
      humidityToLocation = false
      continue
    }
    
    if strings.HasPrefix(line, "humidity-to-location map:") {
      humidityToLocation = true 
      temperatureToHumidity = false 
      lightToTemperature = false 
      waterToLight = false 
      fertToWater = false 
      soilToFert = false 
      seedToSoil = false 
      continue
    }

    mapping := strings.Split(line, " ")
    if len(mapping) < 3 {
      fmt.Println("invalid mapping")
      continue
    }

    //fmt.Printf("destination range start %s\n", mapping[0])
    //fmt.Printf("source range start %s\n", mapping[1])
    //fmt.Printf("range length %s\n", mapping[2])

    destinationRangeStart, err := strconv.Atoi(mapping[0])
	  if err != nil {
		  log.Fatal(err)
	  }
    sourceRangeStart, err := strconv.Atoi(mapping[1])
	  if err != nil {
		  log.Fatal(err)
	  }
    rangeLength, err := strconv.Atoi(mapping[2])
	  if err != nil {
		  log.Fatal(err)
	  }
    if seedToSoil {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.seedToSoil = append(mappings.seedToSoil, mapStruct)
    }

    if soilToFert {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.soilToFert = append(mappings.soilToFert, mapStruct)
    }
    
    if fertToWater {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.fertToWater = append(mappings.fertToWater, mapStruct)
    }

    if waterToLight {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.waterToLight = append(mappings.waterToLight, mapStruct)
    }
    
    if lightToTemperature {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.lightToTemperature = append(mappings.lightToTemperature, mapStruct)
    }

    if temperatureToHumidity {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.temperatureToHumidity = append(mappings.temperatureToHumidity, mapStruct)
    }
    
    if humidityToLocation {
      mapStruct := Mapping{ sourceStart: sourceRangeStart, destinationStart: destinationRangeStart, rangeLen: rangeLength}
      mappings.humidityToLocation = append(mappings.humidityToLocation, mapStruct)
    }
  }
  fmt.Println("--- end parsing ---")

  return mappings, seeds, nil
}

func part1() {
  mappings, seeds, err := parse("input.prod")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("seeds: %d\n", seeds)
  fmt.Printf("seed to soil: %d\n", mappings.seedToSoil)
  fmt.Printf("soil to fert: %d\n", mappings.soilToFert)
  fmt.Printf("fert to water: %d\n",mappings.fertToWater)
  fmt.Printf("water to light: %d\n",mappings.waterToLight)
  fmt.Printf("light to temp: %d\n",mappings.lightToTemperature)
  fmt.Printf("temp to humidity: %d\n",mappings.temperatureToHumidity)
  fmt.Printf("humidity to location: %d\n",mappings.humidityToLocation)

  lowestLoc := 0
  for _, seed := range seeds {
    fmt.Printf("seed: %d\n", seed)
    soil := 0
    for _, mapping := range mappings.seedToSoil {
      if seed >= mapping.sourceStart && seed < mapping.sourceStart + mapping.rangeLen {
        soil = mapping.destinationStart + (seed - mapping.sourceStart)
      }
    }
    if soil == 0 {
      soil = seed
    }
    fmt.Printf("soil: %d\n", soil)

    fert := 0
    for _, mapping := range mappings.soilToFert {
      if soil >= mapping.sourceStart && soil < mapping.sourceStart + mapping.rangeLen {
        fert = mapping.destinationStart + (soil - mapping.sourceStart)
      }
    }
    if fert == 0 {
      fert = soil
    }
    fmt.Printf("fert: %d\n", fert)
    
    water := 0
    for _, mapping := range mappings.fertToWater {
      if fert >= mapping.sourceStart && fert < mapping.sourceStart + mapping.rangeLen {
        water = mapping.destinationStart + (fert - mapping.sourceStart)
      }
    }
    if water == 0 {
      water = fert
    }
    fmt.Printf("water: %d\n", water)
    
    light := 0
    for _, mapping := range mappings.waterToLight {
      if water >= mapping.sourceStart && water < mapping.sourceStart + mapping.rangeLen {
        light = mapping.destinationStart + (water - mapping.sourceStart)
      }
    }
    if light == 0 {
      light = water
    }
    fmt.Printf("light: %d\n", light)

    temp := 0
    for _, mapping := range mappings.lightToTemperature {
      if light >= mapping.sourceStart && light < mapping.sourceStart + mapping.rangeLen {
        temp = mapping.destinationStart + (light - mapping.sourceStart)
      }
    }
    if temp == 0 {
      temp = light
    }
    fmt.Printf("temp: %d\n", temp)
    
    hum := 0
    for _, mapping := range mappings.temperatureToHumidity {
      if temp >= mapping.sourceStart && temp < mapping.sourceStart + mapping.rangeLen {
        hum = mapping.destinationStart + (temp - mapping.sourceStart)
      }
    }
    if hum == 0 {
      hum = temp
    }
    fmt.Printf("hum: %d\n", hum)

    loc := 0
    for _, mapping := range mappings.humidityToLocation {
      if hum >= mapping.sourceStart && hum < mapping.sourceStart + mapping.rangeLen {
        loc = mapping.destinationStart + (hum - mapping.sourceStart)
      }
    }
    if loc == 0 {
      loc = hum
    }
    fmt.Printf("loc: %d\n", loc)

    if lowestLoc == 0 || loc < lowestLoc {
      lowestLoc = loc
    }

    fmt.Printf("lowest loc: %d\n", lowestLoc)
  }
}

func main() {
  part1()
}
