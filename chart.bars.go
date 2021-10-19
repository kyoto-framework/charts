package ktc

import (
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"github.com/yuriizinets/kyoto"
)

type ChartBars struct {
	Title string
	Bars  []Bar

	// Axis naming
	YAxis string

	Render
	Size
}

type Bar struct {
	Value float64
	Label string

	ColorFill   [4]uint8
	ColorStroke [4]uint8
}

func (c *ChartBars) AfterAsync() {
	c.Render.Render(c.build())
}

func (c *ChartBars) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Resize": func(args ...interface{}) {
			width := args[0].(float64)
			height := args[1].(float64)
			c.Size.Resize(int(width), int(height))
			c.Render.Render(c.build())
		},
	}
}

func (c *ChartBars) build() ImplementsGoChart {
	// Build bars
	bars := []chart.Value{}
	for _, b := range c.Bars {
		style := chart.Style{}
		style.StrokeWidth = 1
		style.FillColor = colorChartFromUint(b.ColorFill)
		style.StrokeColor = colorChartFromUint(b.ColorStroke)
		bars = append(bars, chart.Value{
			Value: b.Value,
			Label: b.Label,
			Style: style,
		})
	}
	// Build and return chart
	return chart.BarChart{
		Title:  c.Title,
		Width:  c.Size.GetWidth(),
		Height: c.Size.GetHeight(),
		Bars:   bars,
		YAxis: chart.YAxis{
			Name:  c.YAxis,
			Style: chart.Shown(),
			GridMajorStyle: chart.Style{
				Hidden:      false,
				StrokeColor: drawing.ColorBlack,
				StrokeWidth: 1.5,
			},
			GridMinorStyle: chart.Style{
				Hidden:      false,
				StrokeColor: drawing.Color{R: 0, G: 0, B: 0, A: 100},
				StrokeWidth: 1.0,
			},
		},
	}
}
