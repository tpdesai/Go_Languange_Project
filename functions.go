package main


import (

  "math/rand"
)






func SandPilesMultiProc(initialcheckerboard CheckerBoard, numGens int, numProcs int) CheckerBoard {
//	winnings := 0 // amount won or lost

	c := make(chan CheckerBoard,numProcs)
  var finalboard CheckerBoard

	//play the game concurrently over numProcs processes
	for i := 0; i < numProcs; i++ {
		//create a new PRNG object that only this goroutine has access to
		//source := rand.NewSource(time.Now().UnixNano())
		//generator := rand.New(source) // PRNG object
		go SandPilesOneProc(initialcheckerboard, numGens/numProcs, c)
	}

	//grab all values from channel
	for i := 0; i < numProcs; i++ {
		finalboard = <-c
	}

	return finalboard
}

func SandPilesSerial(initialcheckerboard CheckerBoard, numGens int) CheckerBoard {
  var finalboard CheckerBoard
//	source := rand.NewSource(time.Now().UnixNano())
	//generator := rand.New(source) // PRNG object

	finalboard= SandPiles(initialcheckerboard, numGens) // true if win, false if loss

	return finalboard
}

//TotalWinOneProc takes a number of trials and an integer channel as input.
//It simulates craps numTrials times. It then places
//the total winnings into the channel.
func SandPilesOneProc(initialcheckerboard CheckerBoard, numGens int, c chan CheckerBoard) {

	//c = make(chan CheckerBoard,1)
	var finalboard CheckerBoard


	finalboard = SandPiles(initialcheckerboard, numGens) // returns +1 or -1


	c <- finalboard
}

//type CheckerBoard [][]int

func SandPiles(initialcheckerboard CheckerBoard, numGens int) CheckerBoard {

  checkerboards := make([]CheckerBoard, numGens+1)
  checkerboards[0] = initialcheckerboard
  for i:=1 ; i< numGens+1 ; i++ {

  checkerboards[i]=UpdateCheckerBoard(checkerboards[i-1])
  }
  return checkerboards[len(checkerboards)-1]
//  return checkerboards
}

func UpdateCheckerBoard(checkerboard CheckerBoard) CheckerBoard {
  checkerboardnext := CheckForFour(checkerboard)
  return checkerboardnext
}

func CheckForFour(checkerboard CheckerBoard) CheckerBoard{
  for r:=0; r<CountRows(checkerboard); r++ {
    for c:=0; c<CountCols(checkerboard); c++ {
      if checkerboard[r][c] >= 4 {
        checkerboard = MoveCoins(checkerboard,r,c)
      }else if checkerboard[r][c] <4 {
        checkerboard[r][c]=checkerboard[r][c]
    }
  }
}
  return checkerboard
}

func MoveCoins(checkerboard CheckerBoard, row, col int) CheckerBoard {
  checkerboard[row][col] = checkerboard[row][col]-4
  if InField(checkerboard, row+1, col) == true {
    checkerboard[row+1][col] = checkerboard[row+1][col]+1
  }
  if InField(checkerboard, row, col+1) == true {
    checkerboard[row][col+1] = checkerboard[row][col+1]+1
  }
  if InField(checkerboard, row-1, col) == true {
    checkerboard[row-1][col] = checkerboard[row-1][col]+1
  }
  if InField(checkerboard, row, col-1) == true {
    checkerboard[row][col-1] = checkerboard[row][col-1]+1
  }
  return checkerboard
}



// input the checkerboard with a row and column value
// output a bool true or false if the row col is in field or not

func InField(checkerboard CheckerBoard, row, col int) bool {

  if row<0  {
    return false
} else if col<0 {
    return false
} else if row> CountRows(checkerboard)-1{
    return false
} else if col> CountCols(checkerboard)-1 {
    return false
  }

  return true
}

// input a number of coins to be added to the center most square
// output the CheckerBoard
func CentralCoinsAdd(checkerboard CheckerBoard, numCoins int) CheckerBoard {

  checkerboard[CountRows(checkerboard)/2][CountCols(checkerboard)/2] = numCoins
  return checkerboard
}


func RandomCoinsAdd(checkerboard CheckerBoard, numCoins int) CheckerBoard {
  y:= rand.Intn(numCoins-1)
  numC := numCoins
  if numC !=0 {
    for i:=0; i<CountRows(checkerboard); i++ {
      for j:=0 ; j<CountCols(checkerboard); j++ {
        checkerboard[rand.Intn(CountRows(checkerboard)-1)][rand.Intn(CountCols(checkerboard)-1)] = y
        numCoins -=y

      }
    }

  }

  return checkerboard
}

// count the number of rows in the checkerboard with input as the checkerboard matrix
// output is the integer value of number of rows
func CountRows(checkerboard CheckerBoard) int {
  return len(checkerboard)
}

// count the number of columns in the checkerboard with input as the checkerboard matrix
// output is the integer value of the number of columns
func CountCols(checkerboard CheckerBoard) int{
  return len(checkerboard[0])
}


// initialize a checkerboard with number of rows and columns and numRows and numCols
// output an empty CheckerBoard
func InitializeCheckerBoard(numRows, numCols, numCoins int) CheckerBoard {
  var checkerboard CheckerBoard
  checkerboard = make(CheckerBoard, numRows)

  for i := range checkerboard {
    checkerboard[i] = make([]int, numCols)
  }
  return CentralCoinsAdd(checkerboard, numCoins)
}
func InitializeCheckerBoardRandom(numRows, numCols, numCoins int) CheckerBoard {
  var checkerboard CheckerBoard
  checkerboard = make(CheckerBoard, numRows)

  for i := range checkerboard {
    checkerboard[i] = make([]int, numCols)
  }
  return RandomCoinsAdd(checkerboard, numCoins)
}
