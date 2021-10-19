package ktc

import (
	"html/template"
	"testing"

	"github.com/yuriizinets/kyoto"
)

type TestChartBarsPage struct {
	Bars        kyoto.Component
	BarsColored kyoto.Component
}

func (p *TestChartBarsPage) Template() *template.Template {
	return template.Must(template.New("chart.bars_test.html").Funcs(kyoto.TFuncMap()).ParseGlob("templates/*.html"))
}

func (p *TestChartBarsPage) Init() {
	bars := []Bar{
		{Value: 5.25, Label: "Test 1"},
		{Value: 4.88, Label: "Test 2"},
		{Value: 4.74, Label: "Test 3"},
		{Value: 3.22, Label: "Test 4"},
		{Value: 3, Label: "Test 5"},
		{Value: 6, Label: "Test 6"},
	}
	barsColored := []Bar{
		{Value: 5.25, Label: "Test 1", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
		{Value: 4.88, Label: "Test 2", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
		{Value: 4.74, Label: "Test 3", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
		{Value: 3.22, Label: "Test 4", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
		{Value: 3, Label: "Test 5", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
		{Value: 6, Label: "Test 6", ColorFill: ColorUintFromHex("0000FF11"), ColorStroke: ColorUintFromHex("0000FF")},
	}
	p.Bars = kyoto.RegC(p, &ChartBars{
		Title: "Bars chart",
		Bars:  bars,
		Render: Render{
			Adaptive: true,
		},
	})
	p.BarsColored = kyoto.RegC(p, &ChartBars{
		Title: "Colored bars chart",
		Bars:  barsColored,
		Render: Render{
			Adaptive: true,
		},
	})
}

func TestChartBars(t *testing.T) {
	serve(&TestChartBarsPage{})
}
