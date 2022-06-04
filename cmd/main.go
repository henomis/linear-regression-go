package main

import (
	"fmt"
	"log"

	"github.com/henomis/linear-regression-go/pkg/datacleaner"
	"github.com/henomis/linear-regression-go/pkg/dataminer"
	"github.com/henomis/linear-regression-go/pkg/dataplotter"
	"github.com/henomis/linear-regression-go/pkg/datatrainer"
)

func main() {

	// GATHER DATA
	dataminer := dataminer.New(
		"production_budget_usd",
		"worldwide_gross_usd",
	)

	dataFrame, err := dataminer.GatherFromFile(
		"cost_revenue_dirty.csv",
	)
	if err != nil {
		log.Fatal("unable to gather data: ", err)
	}

	// CLEAN DATA
	dataCleaner := datacleaner.New(dataFrame)
	dataCleaner.Clean()

	dataFrame.Dump()

	// TRAIN DATA
	dataTrainer := datatrainer.New(
		dataFrame.X(),
		dataFrame.Y(),
	)

	alpha, beta := dataTrainer.LinearRegression()
	fmt.Println("alpha =", alpha, " beta =", beta)

	// PLIOT DATA
	dataPlotter := dataplotter.New(
		dataFrame,
		alpha,
		beta,
	)
	dataPlotter.SetTitles(
		"Movies production budget and gross",
		"Production budget",
		"Worldwide gross",
	)

	err = dataPlotter.PlotToFile("output.png")
	if err != nil {
		log.Fatal("unable to plot data: ", err)
	}

	fmt.Println("plot saved successfully")

}
