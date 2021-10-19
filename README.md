# kyoto charts

SSR Charts as a Component. Prototyped for using in `kyoto` pages.  
Kyoto charts are using `go-chart` under the hood to generate charts on server side and include rendered result into component's layout.

## Requirements

- `kyoto` page
- (recommended) configured SSA

## Installation

- Install go package `go get github.com/yuriizinets/kyoto-charts`
- Copy files from `templates` into your templates folder (you can ignore test files)

## Features

- No additional JS payload, only `kyoto` communication layer
- Screen size adaptive, re-generation on load

## How to use

```go
package main

import (
    "github.com/yuriizinets/kyoto"
    ktc "github.com/yuriizinets/kyoto-charts"
)

type PageIndex struct {
    ExampleChart kyoto.Component
}

func (p *PageIndex) Init() {
    // Init example series
    series := ContinuousSeries{
        Name:    "Series",
        XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
        YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(),
    }
    // Init component
    p.ExampleChart = kyoto.RegC(p, &ktc.ChartSeries{
        Title: "Basic series chart",
        Series: []interface{}{
            series,
        },
        Size: Size{
            Height: 400,
        },
        XAxis: "Axis X",
        YAxis: "Axis Y",
        Render: Render{
            Adaptive: true,
        },
    })
}
```

## Examples

### Series

![example](https://imgur.com/HNTDp7E.png)
![example](https://imgur.com/PXUmndg.png)

### Bars

![example](https://imgur.com/WiHcZpL.png)

### Donut

![example](https://imgur.com/lRBZ9Og.png)

## Support

<a target="_blank" href="https://www.buymeacoffee.com/yuriizinets"><img alt="Buy me a Coffee" src="https://github.com/egonelbre/gophers/blob/master/.thumb/animation/buy-morning-coffee-3x.gif?raw=true"></a>


Or directly with Bitcoin: `18WY5KRWVKVxjWKFzJLqqALsRrsDh4snqg`
