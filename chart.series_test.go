package ktc

import (
	"html/template"
	"testing"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/yuriizinets/kyoto"
)

type TestChartSeriesPage struct {
	Series               kyoto.Component
	SeriesSMA            kyoto.Component
	SeriesTime           kyoto.Component
	SeriesOverlap        kyoto.Component
	SeriesOverlapColored kyoto.Component
	SeriesAnnotations    kyoto.Component
}

func (p *TestChartSeriesPage) Template() *template.Template {
	return template.Must(template.New("chart.series_test.html").Funcs(kyoto.TFuncMap()).ParseGlob("templates/*.html"))
}

func (p *TestChartSeriesPage) Init() {
	// Prepare shared values
	continuousSeries := ContinuousSeries{
		Name:    "Continuous",
		XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(),
	}
	smaSeries := SMASeries{
		Name:        "SMA",
		InnerSeries: continuousSeries,
	}
	tsvalues := []time.Time{}
	for i := 0; i < 30; i++ {
		tsvalues = append(tsvalues, time.Now().AddDate(0, -3, i))
	}
	timeSeries := TimeSeries{
		Name:    "Time Series",
		XValues: tsvalues,
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(),
	}
	timeSeries2 := TimeSeries{
		Name:    "Time Series 2",
		XValues: tsvalues,
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(50)}.Values(),
	}
	for i := 0; i < 10; i++ {
		timeSeries2.YValues[i*9] = timeSeries2.YValues[i*9] + 50
	}
	timeSeriesColored := TimeSeries{
		Name:        "Time Series",
		XValues:     tsvalues,
		YValues:     timeSeries.YValues,
		ColorFill:   ColorUintFromHex("0000FF11"),
		ColorStroke: ColorUintFromHex("0000FF"),
	}
	timeSeriesColored2 := TimeSeries{
		Name:        "Time Series 2",
		XValues:     tsvalues,
		YValues:     timeSeries2.YValues,
		ColorFill:   ColorUintFromHex("0000FF11"),
		ColorStroke: ColorUintFromHex("0000FF"),
	}
	annotationsSeries := AnnotationsSeries{
		Name: "Annotations",
		Annotations: []Annotation{
			{
				Label:  "Annotation 1",
				XValue: continuousSeries.XValues[33],
				YValue: continuousSeries.YValues[33],
			},
			{
				Label:  "Annotation 2",
				XValue: continuousSeries.XValues[66],
				YValue: continuousSeries.YValues[66],
			},
		},
	}
	// Init components
	p.Series = kyoto.RegC(p, &ChartSeries{
		Title: "Basic series chart",
		Series: []interface{}{
			continuousSeries,
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
	p.SeriesSMA = kyoto.RegC(p, &ChartSeries{
		Title: "Series chart with SMA",
		Series: []interface{}{
			continuousSeries,
			smaSeries,
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
	p.SeriesTime = kyoto.RegC(p, &ChartSeries{
		Title: "Time series chart",
		Series: []interface{}{
			timeSeries,
		},
		Size: Size{
			Height: 400,
		},
		XAxis: "Axis X (Time)",
		YAxis: "Axis Y",
		Render: Render{
			Adaptive: true,
		},
	})
	p.SeriesOverlap = kyoto.RegC(p, &ChartSeries{
		Title: "Time series chart with overlap",
		Series: []interface{}{
			timeSeries,
			timeSeries2,
		},
		Size: Size{
			Height: 400,
		},
		XAxis: "Axis X (Time)",
		YAxis: "Axis Y",
		Render: Render{
			Adaptive: true,
		},
	})
	p.SeriesOverlapColored = kyoto.RegC(p, &ChartSeries{
		Title: "Colored time series chart with overlap",
		Series: []interface{}{
			timeSeriesColored,
			timeSeriesColored2,
		},
		Size: Size{
			Height: 400,
		},
		XAxis: "Axis X (Time)",
		YAxis: "Axis Y",
		Render: Render{
			Adaptive: true,
		},
	})
	p.SeriesAnnotations = kyoto.RegC(p, &ChartSeries{
		Title: "Series chart with annotations",
		Series: []interface{}{
			continuousSeries,
			annotationsSeries,
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

func TestChartSeries(t *testing.T) {
	serve(&TestChartSeriesPage{})
}
