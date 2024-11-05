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
*   - A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
*   - A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
*
* All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*
* --- Part Two ---
*
* The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.
*
* The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.
*
* For example:
*
*   - A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
*   - A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
*
* How many total feet of ribbon should they order?
*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "sort"
  "strings"
  "strconv"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error:", err)
  }
  defer file.Close()

  totalPaper  := 0
  totalRibbon := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    paper, err := getPaperSurfaceArea(scanner.Text())
    if err != nil {
      fmt.Println("Error calculating paper surface area:", err)
      return
    }
    totalPaper += paper

    ribbon, err := getRibbonLength(scanner.Text())
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    totalRibbon += ribbon
  }

  fmt.Println("Part 1 | Total Paper Needed: ", totalPaper, "ft^2")
  fmt.Println("Part 2 | Total Ribbon Needed:", totalRibbon, "ft")
}

func getDimensions(dimensions string) (length, width, height int, err error) {
  dimensionValues := strings.Split(dimensions, "x")

  if len(dimensionValues) != 3 {
    return 0, 0, 0, fmt.Errorf("invalid dimensions format")
  }

  length, lengthErr := strconv.Atoi(dimensionValues[0])
  width, widthErr := strconv.Atoi(dimensionValues[1])
  height, heightErr := strconv.Atoi(dimensionValues[2])

  if lengthErr != nil || widthErr != nil || heightErr != nil {
    return 0, 0, 0, fmt.Errorf("Converting string to integer failed")
  }

  return length, width, height, err
}

func getPaperSurfaceArea(dimensions string) (paperSurfaceArea int, err error){
  length, width, height, err := getDimensions(dimensions)

  if err != nil {
    return 0, err
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

  paperSurfaceArea = area + slack
  return paperSurfaceArea, nil
}

func getRibbonLength(dimensions string) (ribbonLength int, err error)  {
  length, width, height, err := getDimensions(dimensions)

  if err != nil {
    return 0, err
  }

  sides := []int{length, width, height}
  sort.Ints(sides)
  sidesPerimeter := 2 * (sides[0] + sides[1])

  ribbonLength = (length * width * height) + sidesPerimeter
  return ribbonLength, nil
}
