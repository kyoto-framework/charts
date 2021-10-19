package ktc

import (
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/yuriizinets/kyoto"
)

type ChartSeries struct {
	// General
	Title  string
	Series []interface{}

	// Axis naming
	XAxis          string
	YAxis          string
	YAxisSecondary string

	// Helpers
	Render
	Size
}

type ContinuousSeries struct {
	Name    string
	XValues []float64
	YValues []float64

	ColorFill   [4]uint8
	ColorStroke [4]uint8
}

type TimeSeries struct {
	Name    string
	XValues []time.Time
	YValues []float64

	ColorFill   [4]uint8
	ColorStroke [4]uint8
}

type SMASeries struct {
	Name        string
	InnerSeries interface{}
}

type AnnotationsSeries struct {
	Name        string
	Annotations []Annotation
}

type Annotation struct {
	XValue float64
	YValue float64
	Label  string
}

func (c *ChartSeries) AfterAsync() {
	c.Render.Render(c.build())
}

func (c *ChartSeries) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Resize": func(args ...interface{}) {
			width := args[0].(float64)
			height := args[1].(float64)
			c.Size.Resize(int(width), int(height))
			c.Render.Render(c.build())
		},
	}
}

func (c *ChartSeries) build() ImplementsGoChart {
	// Build series
	series := []chart.Series{}
	for _, s := range c.Series {
		series = append(series, c.buildSeries(s))
	}
	// Build and return chart
	return chart.Chart{
		Title:  c.Title,
		Width:  c.Size.GetWidth(),
		Height: c.Size.GetHeight(),
		Series: series,

		XAxis: chart.XAxis{
			Name: c.XAxis,
		},
		YAxis: chart.YAxis{
			Name: c.YAxis,
			GridMajorStyle: chart.Style{
				Hidden:      false,
				StrokeColor: colorChartFromUint([4]uint8{33, 33, 33, 66}),
				StrokeWidth: 1.5,
			},
			GridMinorStyle: chart.Style{
				Hidden:      false,
				StrokeColor: colorChartFromUint([4]uint8{33, 33, 33, 33}),
				StrokeWidth: 1.0,
			},
		},
		YAxisSecondary: chart.YAxis{
			Name: c.YAxisSecondary,
		},
	}
}

func (c *ChartSeries) buildSeries(s interface{}) chart.Series {
	switch _s := s.(type) {
	// Continuous Series
	case ContinuousSeries:
		style := chart.Style{}
		style.FillColor = colorChartFromUint(_s.ColorFill)
		style.StrokeColor = colorChartFromUint(_s.ColorStroke)
		return chart.ContinuousSeries{
			Name:    _s.Name,
			XValues: _s.XValues,
			YValues: _s.YValues,
			Style:   style,
		}
	// Time Series
	case TimeSeries:
		style := chart.Style{}
		style.FillColor = colorChartFromUint(_s.ColorFill)
		style.StrokeColor = colorChartFromUint(_s.ColorStroke)
		return chart.TimeSeries{
			Name:    _s.Name,
			XValues: _s.XValues,
			YValues: _s.YValues,
			Style:   style,
		}
	case SMASeries:
		return chart.SMASeries{
			Name:        _s.Name,
			InnerSeries: c.buildSeries(_s.InnerSeries).(chart.ValuesProvider),
		}
	// Annotations Series
	case AnnotationsSeries:
		style := chart.Style{}
		values := []chart.Value2{}
		for _, a := range _s.Annotations {
			values = append(values, chart.Value2{
				Label:  a.Label,
				XValue: a.XValue,
				YValue: a.YValue,
			})
		}
		return chart.AnnotationSeries{
			Name:        _s.Name,
			Annotations: values,
			Style:       style,
		}
	default:
		panic("Not supported series type")
	}
}
