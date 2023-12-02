package main

import (
    "os"
    "bufio"
    "log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

func main() {
  part2()
}

func part2() {
  f, err := os.Open("input.prod")
  if err != nil {
      log.Fatal(err)
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  constraints := map[string]int{
      "blue": 14,
      "red": 12,
      "green": 13,
  }

  gamePattern := regexp.MustCompile(`Game (\d+): `)
  var validGames []int
  sum := 0
  sumPowers := 0

  for scanner.Scan() {
      line := scanner.Text()
      //fmt.Println(line)
      valid := true
    gameN := gamePattern.FindStringSubmatch(line)
      if len(gameN) < 1 {
          fmt.Println(gameN)
          fmt.Println("game no invalid")
          break
      }
      gameNoString := gameN[1]
      rounds := strings.Split(gamePattern.ReplaceAllString(line, ""), ";")
      //fmt.Printf("rounds: %s\n", rounds)
      minPossible := map[string]int{
          "blue": 0,
          "red": 0,
          "green": 0,
      }
      for _, roundPart := range rounds {
          counts := map[string]int{
              "blue": 0,
              "red": 0,
              "green": 0,
          }
          round := strings.TrimSpace(roundPart)
          //fmt.Printf("round: %s\n", round)
          set := strings.Split(round, ",")
          for _, setPart := range set {
              //fmt.Printf("set part: %s\n", setPart)
              colorPattern := regexp.MustCompile(`blue|red|green`)
              color := colorPattern.FindString(setPart)
              //fmt.Println(color)
              countPattern := regexp.MustCompile(`(\d+)`)
              countString := countPattern.FindString(setPart)
              count, err := strconv.Atoi(countString)
              if err != nil {
                  fmt.Println("invalid input")
              }
              //fmt.Println(count)
              counts[color] = counts[color] + count
              if counts[color] > constraints[color] {
                  //fmt.Println("invalid round")
                  valid = false
              }
              //fmt.Println(counts)
          }
          for color, count := range counts {
              if minPossible[color] < count {
                  minPossible[color] = count
              }
          }

      }
      if valid {
          gameNo, err := strconv.Atoi(gameNoString)
          if err != nil {
              fmt.Println("invalid gameNo")
          }
          validGames = append(validGames, gameNo)
          sum = sum + gameNo
      }
      //fmt.Println(minPossible)
      power := 0
      for _, count := range minPossible {
          if power == 0 {
            power = count
          } else {
              power = power * count
          }
      }
      //fmt.Printf("power: %d\n", power)
      sumPowers = sumPowers + power
  }
  fmt.Printf("sum: %d\n", sum)
  fmt.Printf("sumPowers: %d\n", sumPowers)
}

func part1() {
   f, err := os.Open("input.prod")
   if err != nil {
       log.Fatal(err)
   }

   defer f.Close()

   scanner := bufio.NewScanner(f)

   constraints := map[string]int{
       "blue": 14,
       "red": 12,
       "green": 13,
   }

   gamePattern := regexp.MustCompile(`Game (\d+): `)
   var validGames []int
   sum := 0

   for scanner.Scan() {
       line := scanner.Text()
       fmt.Printf("line: %s\n", line)
       valid := true
     gameN := gamePattern.FindStringSubmatch(line)
       if len(gameN) < 1 {
           fmt.Println(gameN)
           fmt.Println("game no invalid")
           break
       }
       gameNoString := gameN[1]
       rounds := strings.Split(gamePattern.ReplaceAllString(line, ""), ";")
       //fmt.Printf("rounds: %s\n", rounds)
       for _, roundPart := range rounds {
           if !valid {
               break
           }
           counts := map[string]int{
               "blue": 0,
               "red": 0,
               "green": 0,
           }
           round := strings.TrimSpace(roundPart)
           //fmt.Printf("round: %s\n", round)
           set := strings.Split(round, ",")
           for _, setPart := range set {
               //fmt.Printf("set part: %s\n", setPart)
               colorPattern := regexp.MustCompile(`blue|red|green`)
               color := colorPattern.FindString(setPart)
               //fmt.Println(color)
               countPattern := regexp.MustCompile(`(\d+)`)
               countString := countPattern.FindString(setPart)
               count, err := strconv.Atoi(countString)
               if err != nil {
                   fmt.Println("invalid input")
               }
               //fmt.Println(count)
               counts[color] = counts[color] + count
               if counts[color] > constraints[color] {
                   //fmt.Println("invalid round")
                   valid = false
                   break
               }
               //fmt.Println(counts)
           }
       }
       if valid {
           gameNo, err := strconv.Atoi(gameNoString)
           if err != nil {
               fmt.Println("invalid gameNo")
           }
           validGames = append(validGames, gameNo)
           sum = sum + gameNo
       }
   }
   fmt.Printf("sum: %d\n", sum)
}
