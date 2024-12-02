/* --- Day 3: Perfectly Spherical Houses in a Vacuum ---
*
* Santa is delivering presents to an infinite two-dimensional grid of houses.
*
* He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.
*
* However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?
*
* For example:
*
*   - > delivers presents to 2 houses: one at the starting location, and one to the east.
*   - ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
*   - ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.
*
*   --- Part Two ---
*
* The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.
*
* Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.
*
* This year, how many houses receive at least one present?
*
* For example:
*
*   - ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
*   - ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
*   - ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
*/

package main

import (
  "fmt"
  "io/ioutil"
)

func main()  {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  directions := string(data)
  housesVisited := getHousesVisited(directions)
  housesVisitedWithRobot := getHousesVisitedWithRobot(directions)

  fmt.Println("Part 1 | Houses Visited:", housesVisited)
  fmt.Println("Part 2 | Houses Visited with Robot:", housesVisitedWithRobot)
}

func updatePosition(direction rune, x, y *int) {
  switch direction {
  case '>':
    *x++
  case '<':
    *x--
  case '^':
    *y++
  case 'v':
    *y--
  }
}

func getHousesVisited(directions string) (housesVisited int) {
  x, y := 0, 0
  houses := make(map[string]bool)

  key := fmt.Sprintf("(%d, %d)", x, y)
  houses[key] = true

  for _, direction := range directions {
    updatePosition(direction, &x, &y)

    key := fmt.Sprintf("(%d, %d)", x, y)
    houses[key] = true
  }

  return len(houses)
}

func getHousesVisitedWithRobot(directions string) (housesVisited int) {
  santaX, santaY := 0, 0
  robotX, robotY := 0, 0
  isSantasTurn := true
  houses := make(map[string]bool)

  key := fmt.Sprintf("(%d, %d)", santaX, santaY)
  houses[key] = true

  for _, direction := range directions {
    if isSantasTurn {
      updatePosition(direction, &santaX, &santaY)
      key = fmt.Sprintf("(%d, %d)", santaX, santaY)
    } else {
      updatePosition(direction, &robotX, &robotY)
      key = fmt.Sprintf("(%d, %d)", robotX, robotY)
    }

    isSantasTurn = !isSantasTurn 
    houses[key] = true
  }

  return len(houses)
}

