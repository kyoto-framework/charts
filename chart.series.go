package ktc

import (
	"encoding/json"
	"time"

	"github.com/kyoto-framework/kyoto"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/yuriizinets/go-common"
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

func (s ContinuousSeries) Type() string {
	return "ContinuousSeries"
}

func (s ContinuousSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Type":        s.Type(),
		"Name":        s.Name,
		"XValues":     s.XValues,
		"YValues":     s.YValues,
		"ColorFill":   s.ColorFill,
		"ColorStroke": s.ColorStroke,
	})
}

type TimeSeries struct {
	Name    string
	XValues []time.Time
	YValues []float64

	ColorFill   [4]uint8
	ColorStroke [4]uint8
}

func (s TimeSeries) Type() string {
	return "TimeSeries"
}

func (s TimeSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Type":        s.Type(),
		"Name":        s.Name,
		"XValues":     s.XValues,
		"YValues":     s.YValues,
		"ColorFill":   s.ColorFill,
		"ColorStroke": s.ColorStroke,
	})
}

type SMASeries struct {
	Name        string
	InnerSeries interface{}
}

func (s SMASeries) Type() string {
	return "SMASeries"
}

func (s SMASeries) MarshalJSON() ([]byte, error) {
	inner, err := json.Marshal(s.InnerSeries)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(map[string]interface{}{
		"Type":        s.Type(),
		"Name":        s.Name,
		"InnerSeries": string(inner),
	})
}

type AnnotationsSeries struct {
	Name        string
	Annotations []Annotation
}

func (s AnnotationsSeries) Type() string {
	return "AnnotationsSeries"
}

func (s AnnotationsSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Type":        s.Type(),
		"Name":        s.Name,
		"Annotations": s.Annotations,
	})
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
		Title: c.Title,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
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
	// Cast series to string to avoid behavior variations
	sstr := ""
	if _s, ok := s.(string); ok {
		sstr = _s
	} else {
		sstr = common.JSONDumps(s)
	}
	// Cast series to map for details
	smap := map[string]interface{}{}
	common.JSONLoads(sstr, &smap)
	// Extract series type
	stype := smap["Type"].(string)
	// Build series, depending on type
	switch stype {
	// Continuous Series
	case "ContinuousSeries":
		var _s ContinuousSeries
		common.JSONLoads(sstr, &_s)
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
	case "TimeSeries":
		var _s TimeSeries
		common.JSONLoads(sstr, &_s)
		style := chart.Style{}
		style.FillColor = colorChartFromUint(_s.ColorFill)
		style.StrokeColor = colorChartFromUint(_s.ColorStroke)
		return chart.TimeSeries{
			Name:    _s.Name,
			XValues: _s.XValues,
			YValues: _s.YValues,
			Style:   style,
		}
	case "SMASeries":
		var _s SMASeries
		common.JSONLoads(sstr, &_s)
		return chart.SMASeries{
			Name:        _s.Name,
			InnerSeries: c.buildSeries(_s.InnerSeries).(chart.ValuesProvider),
		}
	// Annotations Series
	case "AnnotationsSeries":
		var _s AnnotationsSeries
		common.JSONLoads(sstr, &_s)
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
