package main

import (
	"fmt"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"math"
)

func init() {
	TestMap["gui.chart"] = &GuiChart{}
}

type GuiChart struct{}

func (t *GuiChart) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Creates ChartLine panel
	cl1 := gui.NewChartLine(500, 300)
	cl1.SetMargins(10, 10, 10, 10)
	cl1.SetBorders(10, 10, 10, 10)
	cl1.SetBordersColor(&math32.Black)
	cl1.SetColor(&math32.White)
	cl1.SetPaddings(8, 8, 8, 8)
	cl1.SetPaddingsColor(&math32.Green)
	cl1.SetPosition(0, 0)

	cl1.SetFormatX("%3.2f")
	cl1.SetFormatY("%2.1f")
	ctx.Gui.Add(cl1)

	minY := float32(-12.0)
	maxY := float32(12.0)
	startX := 0
	firstX := float32(0)
	stepX := float32(2.0)
	countStepX := float32(2.0)

	cl1.SetRangeY(minY, maxY)

	// Graph1
	var g1 *gui.LineGraph
	//data1 := []float32{0, 1, 1.5, 3, 20, 5, -8, 7, 7.5, 9, 9.5}
	//data1 := []float32{0, 1, -1, 2, -2, 3, -3, 4}
	stepX = float32(2.0)
	countStepX = float32(2.0)

	data1 := make([]float32, 0)
	var x float32
	for x = 0; x < 2*math.Pi*10; x += 0.1 {
		//data1 = append(data1, 10*math32.Sin(x)*math32.Sin(x/10))
		data1 = append(data1, 10*math32.Sin(x))
	}
	stepX = float32(2.0)
	countStepX = float32(20)

	cbG1 := gui.NewCheckBox("Graph1")
	cbG1.SetPosition(cl1.Position().X+10, cl1.Position().Y+cl1.Height()+10)
	cbG1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbG1.Value() {
			g1 = cl1.AddGraph(&math32.Color{0, 0, 1}, data1)
			g1.SetLineWidth(2.0)
		} else {
			cl1.RemoveGraph(g1)
			g1 = nil
		}
	})
	ctx.Gui.Add(cbG1)

	// Graph2
	var g2 *gui.LineGraph
	data2 := make([]float32, 0)
	var x2 float32
	for x2 = 0; x2 < 2*math.Pi*10; x2 += 0.1 {
		data2 = append(data2, 5*math32.Cos(x2/3))
	}
	cbG2 := gui.NewCheckBox("Graph2")
	cbG2.SetPosition(cbG1.Position().X+cbG1.Width()+10, cbG1.Position().Y)
	cbG2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbG2.Value() {
			g2 = cl1.AddGraph(&math32.Color{1, 0, 0}, data2)
			g2.SetLineWidth(2.0)
		} else {
			cl1.RemoveGraph(g2)
			g2 = nil
		}
	})
	ctx.Gui.Add(cbG2)

	// Title checkbox
	cbTitle := gui.NewCheckBox("Chart Title")
	cbTitle.SetPosition(cbG1.Position().X, cbG1.Position().Y+cbG1.Height()+10)
	cbTitle.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbTitle.Value() {
			title := gui.NewLabel("Chart Title")
			title.SetBgColor(&math32.White)
			title.SetFontSize(16)
			cl1.SetTitle(title)
		} else {
			cl1.SetTitle(nil)
		}
	})
	cbTitle.SetValue(true)
	ctx.Gui.Add(cbTitle)

	// X Scale checkbox
	cbX := gui.NewCheckBox("X Scale")
	cbX.SetPosition(cbTitle.Position().X+cbTitle.Width()+10, cbTitle.Position().Y)
	cbX.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbX.Value() {
			cl1.SetScaleX(5, &math32.Color{0.8, 0.8, 0.8})
			cl1.SetRangeX(firstX, stepX, countStepX)
		} else {
			cl1.ClearScaleX()
		}
	})
	cbX.SetValue(true)
	ctx.Gui.Add(cbX)

	// Y Scale checkbox
	cbY := gui.NewCheckBox("Y Scale")
	cbY.SetPosition(cbX.Position().X+cbX.Width()+10, cbX.Position().Y)
	cbY.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbY.Value() {
			cl1.SetScaleY(5, &math32.Color{0.8, 0.8, 0.8})
		} else {
			cl1.ClearScaleY()
		}
	})
	cbY.SetValue(true)
	ctx.Gui.Add(cbY)

	// startX
	sG1sx := gui.NewHSlider(100, 20)
	sG1sx.SetPosition(cbTitle.Position().X, cbTitle.Position().Y+cbTitle.Height()+10)
	sG1sx.SetValue(0)
	sG1sx.SetText(fmt.Sprintf("startX:%d", startX))
	sG1sx.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		startX = int(sG1sx.Value() * float32(len(data1)))
		sG1sx.SetText(fmt.Sprintf("startX:%d", startX))
		firstX := float32(startX)
		cl1.SetRangeX(firstX, stepX, countStepX)
		if g1 != nil {
			g1.SetData(data1[startX:])
		}
		if g2 != nil {
			g2.SetData(data2[startX:])
		}
	})
	ctx.Gui.Add(sG1sx)

	// step and countStepX
	sG1cx := gui.NewHSlider(100, 20)
	sG1cx.SetPosition(sG1sx.Position().X+sG1sx.Width()+10, sG1sx.Position().Y)
	sG1cx.SetValue(float32(stepX))
	sG1cx.SetText(fmt.Sprintf("stepX:%2.1f", stepX))
	sG1cx.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		stepX = float32(math.Trunc(float64(1 + 10*sG1cx.Value())))
		countStepX = 10 * stepX
		sG1cx.SetText(fmt.Sprintf("stepX:%2.1f", stepX))
		cl1.SetRangeX(0, stepX, countStepX)
	})
	ctx.Gui.Add(sG1cx)

	// minY
	sG1y0 := gui.NewHSlider(100, 20)
	sG1y0.SetPosition(sG1cx.Position().X+sG1cx.Width()+10, sG1cx.Position().Y)
	sG1y0.SetText(fmt.Sprintf("minY:%3.1f", minY))
	sG1y0.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		minY = -20 + 20*sG1y0.Value()
		sG1y0.SetText(fmt.Sprintf("minY:%3.1f", minY))
		cl1.SetRangeY(minY, maxY)
	})
	ctx.Gui.Add(sG1y0)

	// maxY
	sG1y1 := gui.NewHSlider(100, 20)
	sG1y1.SetPosition(sG1y0.Position().X+sG1y0.Width()+10, sG1y0.Position().Y)
	sG1y1.SetText(fmt.Sprintf("maxY:%3.1f", maxY))
	sG1y1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		maxY = 20 * sG1y1.Value()
		sG1y1.SetText(fmt.Sprintf("maxY:%3.1f", maxY))
		cl1.SetRangeY(minY, maxY)
	})
	ctx.Gui.Add(sG1y1)

	// Y auto range
	cbYa := gui.NewCheckBox("Y auto")
	cbYa.SetPosition(sG1y1.Position().X+sG1y1.Width()+10, sG1y1.Position().Y)
	cbYa.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if cbYa.Value() {
			cl1.SetRangeYauto(true)
		} else {
			cl1.SetRangeYauto(false)
		}
	})
	ctx.Gui.Add(cbYa)
}

func (t *GuiChart) Render(ctx *Context) {

}