/*
* --- Day 2: I Was Told There Would Be No Math ---
* https://adventofcode.com/2015/day/2
*
* The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.
*
* Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.
*
* For example:
*
*     A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
*     A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
*
* All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error:", err)
  }
  defer file.Close()

  totalPaper := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    paper, err := getPaperSurfaceArea(scanner.Text())

    if err != nil {
      fmt.Println("Error:",)
    }

    totalPaper += paper
  }

  fmt.Println("Total Paper Needed:", totalPaper, "ft^2")
}

func getPaperSurfaceArea(dimensions string) (totalPaperSurfaceArea int, err error){
  dimensionValues := strings.Split(dimensions, "x")

  if len(dimensionValues) != 3 {
    return 0, fmt.Errorf("invalid dimensions format")
  }

  length, lengthErr := strconv.Atoi(dimensionValues[0])
  width, widthErr := strconv.Atoi(dimensionValues[1])
  height, heightErr := strconv.Atoi(dimensionValues[2])

  if lengthErr != nil || widthErr != nil || heightErr != nil {
    return 0, fmt.Errorf("Converting string to integer failed")
  }

  side1 := length * width
  side2 := width  * height
  side3 := length * height
 
  surfaceArea1 := 2 * length * width 
  surfaceArea2 := 2 * width  * height
  surfaceArea3 := 2 * height * length

  area := surfaceArea1 + surfaceArea2 + surfaceArea3

  slack := side1
  if side2 < slack {
    slack = side2
  }
  if side3 < slack {
    slack = side3
  }

  totalPaperSurfaceArea = area + slack

  return totalPaperSurfaceArea, nil
}
