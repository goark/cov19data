//go:build run
// +build run

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goark/cov19data"
	"github.com/goark/cov19data/entity"
	"github.com/goark/cov19data/filter"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

func getData() ([]*entity.GlobalData, error) {
	impt, err := cov19data.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer impt.Close()
	return impt.Data(
		filter.WithPeriod(
			values.NewPeriod(
				values.NewDate(2020, time.Month(9), 1),
				values.NewDate(2020, time.Month(9), 7),
			),
		),
		filter.WithCountryCode(values.CC_JP),
		filter.WithRegionCode(values.WPRO),
	)
}

func main() {
	data, err := getData()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	b, err := entity.ExportCSV(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
		fmt.Println(err)
	}
	// Output:
	// Date_reported,Country_code,Country,WHO_region,New_cases,Cumulative_cases,New_deaths,Cumulative_deaths
	// 2020-09-01,JP,Japan,WPRO,527,68392,17,1296
	// 2020-09-02,JP,Japan,WPRO,609,69001,11,1307
	// 2020-09-03,JP,Japan,WPRO,598,69599,12,1319
	// 2020-09-04,JP,Japan,WPRO,669,70268,11,1330
	// 2020-09-05,JP,Japan,WPRO,608,70876,19,1349
	// 2020-09-06,JP,Japan,WPRO,543,71419,8,1357
	// 2020-09-07,JP,Japan,WPRO,437,71856,6,1363
}
