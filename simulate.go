package main

import (
	"fmt"
	"time"

	"github.com/Dae314/placeit-sim/binPlace"
	"github.com/Dae314/placeit-sim/game"
	"github.com/Dae314/placeit-sim/midPlace"
	"github.com/Dae314/placeit-sim/randPlace"
	"github.com/Dae314/placeit-sim/utils"
	"golang.org/x/sync/errgroup"
)

type placeFunc func(*game.PlaceItGame) (int, error)

const parallelism = 1000
const maxTrials = 1000000

func playGame(g *game.PlaceItGame, c chan int, getPlace placeFunc) {
	for {
		switch g.State {
		case game.WinState, game.LoseState:
			c <- g.Score
			return
		case game.DrawState:
			g.Draw()
		case game.PlaceState:
			placement, err := getPlace(g)
			if err != nil {
				fmt.Printf("Error encountered: %v\n", err)
				return
			}
			err = g.Place(placement)
			if err != nil {
				fmt.Printf("Error encountered: %v\n", err)
				return
			}
		}
	}
}

func main() {
	start := time.Now()
	placeMethodNames := []string{
		"Random",
		"Middle",
		"Bin",
	}
	placeMethods := []placeFunc{
		randPlace.GetPlacement,
		midPlace.GetPlacement,
		binPlace.GetPlacement,
	}
	var averages []float32
	var histograms [][]int

	for _, method := range placeMethods {
		resultsC := make(chan int, maxTrials)
		var eg errgroup.Group
		eg.SetLimit(parallelism)
		var resultsS []int
		for i := 0; i < maxTrials; i++ {
			g := game.NewGame()
			eg.Go(func() error {
				playGame(&g, resultsC, method)
				return nil
			})
		}
		go func() {
			eg.Wait()
			close(resultsC)
		}()
		for r := range resultsC {
			resultsS = append(resultsS, r)
		}
		averages = append(averages, utils.Average(resultsS))
		histograms = append(histograms, utils.Histogram(resultsS, 20))
	}
	fmt.Printf("Trials: %d\n", maxTrials)
	for i, name := range placeMethodNames {
		fmt.Printf("Average for %s: %f\n", name, averages[i])
		fmt.Printf("Histogram for %s: %v\n", name, histograms[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\nSimulate took %s", elapsed)
}
