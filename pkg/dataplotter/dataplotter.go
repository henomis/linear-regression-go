package dataplotter

import (
	"image/color"

	plottabledataframe "github.com/henomis/linear-regression-go/pkg/plottable_dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type DataPlotter struct {
	dataFrame *plottabledataframe.PlottableDataFrame
	alpha     float64
	beta      float64

	plot *plot.Plot
}

func New(
	dataFrame *plottabledataframe.PlottableDataFrame,
	alpha,
	beta float64,
) *DataPlotter {

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)

	return &DataPlotter{
		dataFrame: dataFrame,
		alpha:     alpha,
		beta:      beta,
		plot:      plot.New(),
	}
}

func (dp *DataPlotter) SetTitles(plotTitle, xAxisTitle, yAxisTitle string) {
	dp.plot.Title.Text = plotTitle
	dp.plot.X.Label.Text = xAxisTitle
	dp.plot.Y.Label.Text = yAxisTitle
}

func (dp *DataPlotter) PlotToFile(filename string) error {

	line := plotter.NewFunction(
		func(x float64) float64 {
			return dp.beta*x + dp.alpha
		},
	)
	line.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	scatter, err := plotter.NewScatter(dp.dataFrame)
	if err != nil {
		return err
	}
	scatter.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	dp.plot.Add(scatter, line)

	if err := dp.plot.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
