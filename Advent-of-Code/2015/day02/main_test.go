package main

import (
  "testing"
)

func testGetPapeSurfaceArear(t *testing.T)  {
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
      t.Errorf("getPaper(\"%q\") = %d; Expected %d", test.dimension, result, test.expected)
    }
  }
}
