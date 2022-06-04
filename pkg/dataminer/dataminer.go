package dataminer

import (
	"os"

	"github.com/go-gota/gota/dataframe"
	plottabledataframe "github.com/henomis/linear-regression-go/pkg/plottable_dataframe"
)

type DataMiner struct {
	xDataName string
	yDataName string
}

func New(xDataName, yDataName string) *DataMiner {
	return &DataMiner{
		xDataName: xDataName,
		yDataName: yDataName,
	}
}

func (dm *DataMiner) GatherFromFile(
	filename string,
) (*plottabledataframe.PlottableDataFrame, error) {

	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	dataFrame := dataframe.ReadCSV(csvFile).Select(
		[]string{dm.xDataName, dm.yDataName},
	)

	return plottabledataframe.NewDataFramePlottable(
		dataFrame,
		dm.xDataName,
		dm.yDataName,
	), nil

}
