package datatrainer

import "gonum.org/v1/gonum/stat"

type DataTrainer struct {
	xData []float64
	yData []float64
}

func New(xData, yData []float64) *DataTrainer {
	return &DataTrainer{
		xData: xData,
		yData: yData,
	}
}

func (dt *DataTrainer) LinearRegression() (alpha, beta float64) {
	alpha, beta = stat.LinearRegression(
		dt.xData,
		dt.yData,
		nil,
		false,
	)

	return
}
