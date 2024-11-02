/* https://adventofcode.com/2015/day/1
* --- Day 1: Not Quite Lisp ---
* 
* Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out! To save Christmas, he needs you to collect fifty stars by December 25th.
* 
* Collect stars by helping Santa solve puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
* 
* Here's an easy puzzle to warm you up.
* 
* Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.
* 
* An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.
* 
* The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.
* 
* For example:
* 
*     - (()) and ()() both result in floor 0.
*     - ((( and (()(()( both result in floor 3.
*     - ))((((( also results in floor 3.
*     - ()) and ))( both result in floor -1 (the first basement level).
*     - ))) and )())()) both result in floor -3.
* 
* To what floor do the instructions take Santa?
*
* --- Part Two ---
*
* Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.
*
* For example:
*
*     ) causes him to enter the basement at character position 1.
*     ()()) causes him to enter the basement at character position 5.
*
* What is the position of the character that causes Santa to first enter the basement?
*/

package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println(`Please enter the directions for Santa as a argument:
    Example: day1 "<directions>"`)
    return
  }

  directions := os.Args[1]

  // Part 1
  floor, err := getCorrectFloor(directions)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Println("Part 1 | Floor:", floor)

  // Part 2
  floor, err = getBasementPosition(directions)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Println("Part 2 | -1 at Index:", floor, "(index of 0 means the directions do not go to the basement)")
}

func getNextFloor(direction rune, floor *int) error {
  switch direction {
  case '(':
    *floor++
  case ')':
    *floor--
  default:
    return fmt.Errorf("'%c' is not valid character, please use '(' or ')'", direction)
  }
  return nil
}

func getCorrectFloor(directions string) (int, error) {
  floor := 0

  for _, direction := range directions {
    err := getNextFloor(direction, &floor)

    if err != nil {
      return 0, err
    }
  }
  return floor, nil
}

func getBasementPosition(directions string) (int, error) {
  floor := 0

  for index, direction := range directions {
    err := getNextFloor(direction, &floor)

    if err != nil {
      return 0, err
    }

    if floor == -1 {
      index++
      return index, nil
    }
  }
  return 0, nil
}
