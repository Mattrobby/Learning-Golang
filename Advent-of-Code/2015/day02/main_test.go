package main

import (
  "testing"
)

func TestGetPapeSurfaceArear(t *testing.T)  {
  tests := []struct {
    dimension string
    expected  int
  }{
    {"2x3x4", 58},
    {"1x1x10", 43},
  }

  for _, test := range tests {
    result, _ := getPaperSurfaceArea(test.dimension)

    if result != test.expected {
      t.Errorf("getPaperSurfaceArea(\"%q\") = %d; Expected %d", test.dimension, result, test.expected)
    }
  }
}

func TestGetRibbonLength(t *testing.T) {
  tests := []struct {
    dimension string
    expected int
  }{
    {"2x3x4", 34},
    {"1x1x10", 14},
  }

  for _, test := range tests {
    result, _ := getRibbonLength(test.dimension)

    if result != test.expected {
      t.Errorf("geRibbonLength(\"%q\") = %d; Expected %d", test.dimension, result, test.expected)
    }
  }
}
