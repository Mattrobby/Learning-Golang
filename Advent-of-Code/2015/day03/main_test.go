package main

import "testing"

func TestGetVisitedHouses(t *testing.T ) {
  tests := []struct {
    directions string
    expected   int
  }{
    {">", 2},
    {"^>v<", 4},
    {"^v^v^v^v^v", 2},
  }

  for _, test := range tests {
    result := getHousesVisited(test.directions)

    if result != test.expected {
      t.Errorf("getHousesVisited(%q) = %d; Expected %d", test.directions, result, test.expected)
    }
  }
}

func TestGetVisitedHousesWithRobot(t *testing.T ) {
  tests := []struct {
    directions string
    expected   int
  }{
    {">^", 3},
    {"^>v<", 3},
    {"^v^v^v^v^v", 11},
  }

  for _, test := range tests {
    result := getHousesVisitedWithRobot(test.directions)

    if result != test.expected {
      t.Errorf("getHousesVisitedWithRobot(%q) = %d; Expected %d", test.directions, result, test.expected)
    }
  }
}
