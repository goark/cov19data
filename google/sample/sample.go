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
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/fetch"
)

func getData() ([]*entity.JapanData, error) {
	impt, err := google.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, err
	}
	defer impt.Close()
	return impt.Data(
		filter.WithPeriod(
			values.NewPeriod(
				values.Yesterday().AddDay(-27),
				values.Yesterday().AddDay(6),
			),
		),
		filter.WithPrefJpCode(values.PrefJpCode(32)), //SHIMANE
	)
}

func main() {
	data, err := getData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	b, err := entity.ExportCSV(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
