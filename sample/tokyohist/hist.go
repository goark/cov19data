package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goark/cov19data/histogram"
	"github.com/goark/cov19data/tokyo"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

func getHist() ([]*histogram.HistData, error) {
	impt, err := tokyo.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer impt.Close()
	return impt.Histogram(
		values.NewPeriod(
			values.NewDate(2022, time.Month(1), 1),
			values.NewDate(2022, time.Month(4), 22),
		),
		7,
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
		fmt.Println()
	}
	// Output:
	// Date_from,Date_to,Cases,Deaths
	// 2021-12-31,2022-01-06,1520,0
	// 2022-01-07,2022-01-13,10513,0
	// 2022-01-14,2022-01-20,37674,0
	// 2022-01-21,2022-01-27,82264,0
	// 2022-01-28,2022-02-03,119327,0
	// 2022-02-04,2022-02-10,124874,0
	// 2022-02-11,2022-02-17,104490,0
	// 2022-02-18,2022-02-24,87526,0
	// 2022-02-25,2022-03-03,79381,0
	// 2022-03-04,2022-03-10,65814,0
	// 2022-03-11,2022-03-17,57113,0
	// 2022-03-18,2022-03-24,44464,0
	// 2022-03-25,2022-03-31,52709,0
	// 2022-04-01,2022-04-07,52033,0
	// 2022-04-08,2022-04-14,52517,0
	// 2022-04-15,2022-04-21,41336,0
}
