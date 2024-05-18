
package main

import (
  "testing"
  "fmt"
)


////
//type CheckerBoard [][]int

func TestCountCols(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    length       int
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.length = 5

  output:= CountCols(Test.checkerboard)
  if Test.length == output {
    fmt.Println("This is correct CountCols")
  }else {
    t.Errorf("This is incorrect CountCols")
  }

}

func TestCountRows(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    length       int
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.length = 6

  output:= CountRows(Test.checkerboard)
  if Test.length == output {
    fmt.Println("This is correct CountRows")
  }else {
    t.Errorf("This is incorrect CountRows")
  }

}
func TestCountCols1(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    length       int
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.length = 5

  output:= CountCols(Test.checkerboard)
  if Test.length == output {
    fmt.Println("This is correct CountCols again")
  }else {
    t.Errorf("This is incorrect CountCols")
  }

}

func TestCountRows1(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    length       int
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.length = 5

  output:= CountRows(Test.checkerboard)
  if Test.length == output {
    fmt.Println("This is correct CountRows again")
  }else {
    t.Errorf("This is incorrect CountRows ")
  }

}

func TestInField(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    row          int
    col          int
    response     bool
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.row = 5
  Test.col = 4
  Test.response = false

  output:= InField(Test.checkerboard, Test.row, Test.col)
  if Test.response == output {
    fmt.Println("This is correct InField response")
  }else {
    t.Errorf("This is incorrect InField response ")
  }

}


func TestInField1(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    row          int
    col          int
    response     bool
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.row = 4
  Test.col = 4
  Test.response = true

  output:= InField(Test.checkerboard, Test.row, Test.col)
  if Test.response == output {
    fmt.Println("This is correct InField response")
  }else {
    t.Errorf("This is incorrect InField response ")
  }

}


func TestInField2(t *testing.T) {

	type test struct {

    checkerboard CheckerBoard
    row          int
    col          int
    response     bool
	}
  var Test test


  Test.checkerboard = CheckerBoard{[]int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0},
    []int{0,0,0,0,0}}
  Test.row = -1
  Test.col = 4
  Test.response = false

  output:= InField(Test.checkerboard, Test.row, Test.col)
  if Test.response == output {
    fmt.Println("This is correct InField response")
  }else {
    t.Errorf("This is incorrect InField response ")
  }
}
