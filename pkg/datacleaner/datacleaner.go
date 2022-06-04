package datacleaner

import (
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	plottabledataframe "github.com/henomis/linear-regression-go/pkg/plottable_dataframe"
)

type DataCleaner struct {
	dataFrame *plottabledataframe.PlottableDataFrame
}

func New(dataFrame *plottabledataframe.PlottableDataFrame) *DataCleaner {
	return &DataCleaner{
		dataFrame: dataFrame,
	}
}

func (dc *DataCleaner) Clean() {

	// remove $ symbol and useless ','
	dc.dataFrame.DataFrame = dc.dataFrame.DataFrame.Capply(func(s series.Series) series.Series {
		return s.Map(
			func(e series.Element) series.Element {
				elementAsString := strings.ReplaceAll(e.String(), "$", "")
				elementAsString = strings.ReplaceAll(elementAsString, ",", "")
				e.Set(elementAsString)

				return e
			},
		)
	})

	// mutate series to float64 type
	dc.dataFrame.DataFrame = dc.dataFrame.DataFrame.Mutate(
		series.New(
			dc.dataFrame.DataFrame.Col(dc.dataFrame.XColumnName).Float(),
			series.Float,
			dc.dataFrame.XColumnName,
		),
	)

	dc.dataFrame.DataFrame = dc.dataFrame.DataFrame.Mutate(
		series.New(
			dc.dataFrame.DataFrame.Col(dc.dataFrame.YColumnName).Float(),
			series.Float,
			dc.dataFrame.YColumnName,
		),
	)

	// remove movies without worldwide gross (value = $0)
	dc.dataFrame.DataFrame = dc.dataFrame.DataFrame.Filter(
		dataframe.F{
			Colname:    dc.dataFrame.YColumnName,
			Comparator: series.Greater,
			Comparando: 0,
		},
	)

}
