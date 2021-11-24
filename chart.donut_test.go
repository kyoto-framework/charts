package ktc

import (
	"html/template"
	"testing"

	"github.com/kyoto-framework/kyoto"
)

type TestChartDonutPage struct {
	Donut kyoto.Component
}

func (p *TestChartDonutPage) Template() *template.Template {
	return template.Must(template.New("chart.donut_test.html").Funcs(kyoto.TFuncMap()).ParseGlob("templates/*.html"))
}

func (p *TestChartDonutPage) Init() {
	p.Donut = kyoto.RegC(p, &ChartDonut{
		Title: "Donut chart",
		Values: []DonutValue{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "test"},
		},
		Render: Render{
			Adaptive: true,
		},
	})
}

func TestChartDonut(t *testing.T) {
	serve(&TestChartDonutPage{})
}
