# [cov19data] -- Importing WHO COVID-2019 Cases Global Data

[![check vulns](https://github.com/spiegel-im-spiegel/cov19data/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/cov19data/actions)
[![lint status](https://github.com/spiegel-im-spiegel/cov19data/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/cov19data/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/cov19data/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/cov19data.svg)](https://github.com/spiegel-im-spiegel/cov19data/releases/latest)

This package is required Go 1.16 or later.

## Usage

### Import This Package

```go
import "github.com/spiegel-im-spiegel/cov19data"
```

### Get COVID-2019 Global Data (raw data) from WHO Web Site

```go
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
```

### Import COVID-2019 Global Data from WHO Web Site to Structured Dataset.

```go
package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/cov19data"
    "github.com/spiegel-im-spiegel/cov19data/entity"
    "github.com/spiegel-im-spiegel/cov19data/filter"
    "github.com/spiegel-im-spiegel/cov19data/values"
    "github.com/spiegel-im-spiegel/errs"
    "github.com/spiegel-im-spiegel/fetch"
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
        fmt.Fprintf(os.Stderr, "%+v\n", err)
        return
    }
    b, err := entity.ExportCSV(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
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
```

### Make Histogram data by COVID-2019 Global Data from WHO Web Site.

```go
package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/cov19data"
    "github.com/spiegel-im-spiegel/cov19data/filter"
    "github.com/spiegel-im-spiegel/cov19data/histogram"
    "github.com/spiegel-im-spiegel/cov19data/values"
    "github.com/spiegel-im-spiegel/errs"
    "github.com/spiegel-im-spiegel/fetch"
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
```

[cov19data]: https://github.com/spiegel-im-spiegel/cov19data "spiegel-im-spiegel/cov19data: Importing WHO COVID-2019 Cases Global Data"
