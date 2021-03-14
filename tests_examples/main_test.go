package main

import "testing"


type testData struct {
  values []int
  sum int
}

dataForTest := []testData{
  {[1, 2, 3], 6},
  {[5, -5, 5], 5},
}



func TestSum(t *testing.T) {
  for _, oneTest := range dataForTest {
    funcSum := Sum(oneTest.values)
    if funcSum != oneTest.sum {
      //some error log
      t.Error("Expected ", oneTest.sum, ", but got", funcSum)
    }
  }

}
