package ktc

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

type ImplementsGoChart interface {
	Render(rp chart.RendererProvider, w io.Writer) error
	GetWidth() int
	GetHeight() int
}

type Render struct {
	ID       string
	Type     string
	Adaptive bool
	Output   template.HTML `json:"-"`
}

func (r *Render) Render(c ImplementsGoChart) {
	// Set defaults
	if r.ID == "" {
		r.ID = strconv.Itoa(rand.Intn(100000))
	}
	if r.Type == "" {
		r.Type = "SVG"
	}
	// Render, depending on type
	switch r.Type {
	case "SVG":
		buffer := bytes.NewBuffer([]byte{})
		err := c.Render(chart.SVG, buffer)
		if err != nil {
			panic(err)
		}
		out := buffer.String()
		out = strings.ReplaceAll(out, fmt.Sprintf(`width="%d"`, c.GetWidth()), `width="100%"`)
		out = strings.ReplaceAll(out, fmt.Sprintf(`height="%d"`, c.GetHeight()), ``)
		out = strings.ReplaceAll(out, `<svg`, fmt.Sprintf(`<svg viewBox="0 0 %d %d"`, c.GetWidth(), c.GetHeight()))
		r.Output = template.HTML(out)
	default:
		panic("Render type not supported")
	}
}
