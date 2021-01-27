// +build run

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/google"
	"github.com/spiegel-im-spiegel/cov19data/google/entity"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/fetch"
)

func getData() ([]*entity.JapanData, error) {
	impt, err := google.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, err
	}
	defer impt.Close()
	return impt.Data()
}

func main() {
	data, err := getData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	hist, err := entity.ExportHistgram(
		data,
		values.NewPeriod(
			values.Yesterday().AddDay(-27),
			values.Yesterday().AddDay(6),
		),
		7,
		filter.WithPrefJpCode(values.PrefJpCode(13)), //TOKYO
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	b, err := histogram.ExportCSV(hist)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
