package plottabledataframe

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
)

type PlottableDataFrame struct {
	DataFrame   dataframe.DataFrame
	XColumnName string
	YColumnName string
}

func NewDataFramePlottable(
	dataFrame dataframe.DataFrame,
	xColumnName,
	yColumnName string,
) *PlottableDataFrame {
	return &PlottableDataFrame{
		DataFrame:   dataFrame,
		XColumnName: xColumnName,
		YColumnName: yColumnName,
	}
}

func (pdf *PlottableDataFrame) Len() int {
	return pdf.DataFrame.Nrow()
}

func (pdf *PlottableDataFrame) XY(i int) (x, y float64) {

	x = pdf.DataFrame.Col(pdf.XColumnName).Elem(i).Float()
	y = pdf.DataFrame.Col(pdf.YColumnName).Elem(i).Float()

	return
}

func (pdf *PlottableDataFrame) X() []float64 {
	return pdf.DataFrame.Col(pdf.XColumnName).Float()
}

func (pdf *PlottableDataFrame) Y() []float64 {
	return pdf.DataFrame.Col(pdf.YColumnName).Float()
}

func (pdf *PlottableDataFrame) Dump() {
	fmt.Println(pdf.DataFrame.Describe())
}
