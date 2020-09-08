# [cov19data] -- Importing WHO COVID-2019 Cases Global Data

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/cov19data.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/cov19data)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/cov19data/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/cov19data.svg)](https://github.com/spiegel-im-spiegel/cov19data/releases/latest)

## Usage

### Import This Package

```go
import "github.com/spiegel-im-spiegel/cov19data/client"
```

### Get COVID-2019 Global Data (raw data) from WHO Web Site

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/spiegel-im-spiegel/cov19data/client"
)

func main() {
    r, err := client.Default().WHOCasesData()
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    defer r.Close()

    if _, err := io.Copy(os.Stdout, r); err != nil {
        fmt.Println()
    }
}
```

### Import COVID-2019 Global Data from WHO Web Site to Structured Dataset.

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/cov19data"
    "github.com/spiegel-im-spiegel/cov19data/client"
    "github.com/spiegel-im-spiegel/cov19data/entity"
    "github.com/spiegel-im-spiegel/cov19data/values"
)

func main() {
    data, err := cov19data.ImportWHOCSV(
        client.Default(),
        entity.WithFilterPeriod(
            values.NewPeriod(
                values.NewDate(2020, time.Month(9), 1),
                values.Yesterday(),
            ),
        ),
        entity.WithCountryCode(values.CC_JP),
        entity.WithRegionCode(values.WPRO),
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v\n", err)
        return
    }

    b, err := entity.ExportWHOCSV(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        return
    }
    if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
        fmt.Println()
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
    "fmt"
    "io"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/cov19data"
    "github.com/spiegel-im-spiegel/cov19data/client"
    "github.com/spiegel-im-spiegel/cov19data/entity"
    "github.com/spiegel-im-spiegel/cov19data/histogram"
    "github.com/spiegel-im-spiegel/cov19data/values"
)

func main() {
    h, err := cov19data.MakeHistogramWHO(
        client.Default(),
        values.NewPeriod(
            values.NewDate(2020, time.Month(8), 1),
            values.Yesterday(),
        ),
        7, //step by 7 days
        entity.WithCountryCode(values.CC_JP),
        entity.WithRegionCode(values.WPRO),
    )
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }

    b, err := histogram.ExportHistCSV(h)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
        fmt.Println()
    }
    // Output:
    // Date_from,Date_to,Cases,Deaths
    // 2020-07-28,2020-08-03,8698,16
    // 2020-08-04,2020-08-10,9303,35
    // 2020-08-11,2020-08-17,7677,52
    // 2020-08-18,2020-08-24,6840,82
    // 2020-08-25,2020-08-31,5358,98
    // 2020-09-01,2020-09-07,3991,84
}
```

[cov19data]: https://github.com/spiegel-im-spiegel/cov19data
