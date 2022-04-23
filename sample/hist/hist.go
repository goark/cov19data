package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goark/cov19data"
	"github.com/goark/cov19data/filter"
	"github.com/goark/cov19data/histogram"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

func getHist() ([]*histogram.HistData, error) {
	impt, err := cov19data.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer impt.Close()
	return impt.Histogram(
		values.NewPeriod(
			values.NewDate(2020, time.Month(9), 1),
			values.NewDate(2020, time.Month(9), 28),
		),
		7,
		filter.WithCountryCode(values.CC_JP),
		filter.WithRegionCode(values.WPRO),
	)
}

func main() {
	hist, err := getHist()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	b, err := histogram.ExportCSV(hist)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
		fmt.Println(err)
	}
	// Output:
	// Date_from,Date_to,Cases,Deaths
	// 2020-09-01,2020-09-07,3991,84
	// 2020-09-08,2020-09-14,3801,79
	// 2020-09-15,2020-09-21,3483,58
	// 2020-09-22,2020-09-28,2991,48
}
