package main

import (
  "fmt"
  "strconv"
  "os"
  "time"
  "log"
  "runtime"

)
func main() {
  fmt.Println("start here for sandpiles")

  SIZE, err1 := strconv.Atoi(os.Args[1])
  if err1 != nil {
    panic(err1)
  }
  PILE, err1 := strconv.Atoi(os.Args[2])
  if err1 != nil {
    panic(err1)
  }


    numProcs := runtime.NumCPU()
    numGens := 10000
    RANDOM := InitializeCheckerBoardRandom(SIZE, SIZE, PILE)
    CENTRAL := InitializeCheckerBoard(SIZE, SIZE, PILE)


    start1 := time.Now()
    SandPilesMultiProc(RANDOM, SIZE, numProcs)
    elapsed1 := time.Since(start1)
    log.Printf("Simulating sandpile in parallel (random) took %s", elapsed1)

    start2 := time.Now()
    SandPilesMultiProc(CENTRAL, SIZE, numProcs)
    elapsed2 := time.Since(start2)
    log.Printf("Simulating sandpile in parallel (central) took %s", elapsed2)

    start3 := time.Now()
    SandPilesSerial(RANDOM,SIZE)
    elapsed3 := time.Since(start3)
    log.Printf("Simulating sandpile serially (random) took %s", elapsed3)

    start4 := time.Now()
    SandPilesSerial(CENTRAL,SIZE)
    elapsed4 := time.Since(start4)
    log.Printf("Simulating sandpile serially (central) took %s", elapsed4)


  	board1 := SandPilesSerial(RANDOM, numGens)
    board2 := SandPilesSerial(CENTRAL, numGens)
    board3 := SandPilesMultiProc(RANDOM, numGens,numProcs)
    board4 := SandPilesMultiProc(CENTRAL, numGens,numProcs)

  	fmt.Println("Done with checkerboards!")


  	cellWidth := 5// each cell is 1 pixel

  	image1 := DrawToCanvas(board1, cellWidth)
    image2 := DrawToCanvas(board2, cellWidth)
    image3 := DrawToCanvas(board3, cellWidth)
    image4 := DrawToCanvas(board4, cellWidth)
  	
    image1.SaveToPNG("SandPileSerialrandom.png")
    image2.SaveToPNG("SandPileSerialcentral.png")
    image3.SaveToPNG("SandPileParallelrandom.png")
    image4.SaveToPNG("SandPileParallelcentral.png")
  	fmt.Println("Images are drawn!")
  }
