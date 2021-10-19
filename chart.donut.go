package ktc

import (
	"github.com/wcharczuk/go-chart/v2"
	"github.com/yuriizinets/kyoto"
)

type ChartDonut struct {
	// General
	Title  string
	Values []DonutValue

	// Helpers
	Render
	Size
}

type DonutValue struct {
	Value float64
	Label string

	ColorFill   [4]uint8
	ColorStroke [4]uint8
}

func (c *ChartDonut) AfterAsync() {
	c.Render.Render(c.build())
}

func (c *ChartDonut) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Resize": func(args ...interface{}) {
			width := args[0].(float64)
			height := args[1].(float64)
			c.Size.Resize(int(width), int(height))
			c.Render.Render(c.build())
		},
	}
}

func (c *ChartDonut) build() ImplementsGoChart {
	// Build values
	values := []chart.Value{}
	for _, v := range c.Values {
		style := chart.Style{}
		style.FillColor = colorChartFromUint(v.ColorFill)
		values = append(values, chart.Value{
			Value: v.Value,
			Label: v.Label,
			Style: style,
		})
	}
	// Build and return chart
	return chart.DonutChart{
		Title:  c.Title,
		Width:  c.Size.GetWidth(),
		Height: c.Size.GetHeight(),
		Values: values,
	}
}
