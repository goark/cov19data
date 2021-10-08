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

	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/tokyo"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

func getHist() ([]*histogram.HistData, error) {
	impt, err := tokyo.NewWeb(context.Background(), fetch.New())
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
	// 2020-09-01,2020-09-07,1032,0
	// 2020-09-08,2020-09-14,1234,0
	// 2020-09-15,2020-09-21,1223,0
	// 2020-09-22,2020-09-28,1029,0
}
