package main

import (
  "testing"
)

func TestGetCorrectFloor(t *testing.T) {
  tests := []struct {
    directions    string
    expected      int
  }{
    {"(())", 0},
    {"()()", 0},
    {"", 0},
    {"(((", 3},
    {"(()(()(", 3},
    {"))(((((", 3},
    {"())", -1},
    {"))(", -1},
    {")))", -3},
    {")())())", -3},
    {"()()())(*",0},
    {"()(%", 0},
  }

  for _, test := range tests {
    result, _ := getCorrectFloor(test.directions)

    if result != test.expected {
      t.Errorf("getCorrectFloor(%q) = %d; Expected %d", test.directions, result, test.expected)
    }
  }
}

func TestGetBasementPosition(t *testing.T) {
  tests := []struct {
    directions    string
    expected      int
  }{
    {")", 1},
    {"()())", 5},
    {"()()", 0},
  }

  for _, test := range tests {
    result, _ := getBasementPosition(test.directions)

    if result != test.expected {
      t.Errorf("testGetBasementPosition(%q) = %d; Expected %d", test.directions, result, test.expected)
    }
  }
}
