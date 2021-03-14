// +build run

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/cov19data"
	"github.com/spiegel-im-spiegel/fetch"
)

func main() {
	impt, err := cov19data.NewWeb(context.Background(), fetch.New())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	defer impt.Close()
	if _, err := io.Copy(os.Stdout, impt.RawReader()); err != nil {
		fmt.Println(err)
	}
}
