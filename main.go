package main

import (
	"fmt"

	//for the gui

	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	//routine for the window
	go func() {
		window := new(app.Window)
		window.Option(app.Title("VDice"))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	//theme := material.NewTheme()

	var startButton widget.Clickable
	th := material.NewTheme()

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			fmt.Print("Closed")
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			//title := material.H1(theme, "VDice")

			// Change the color of the label.
			//black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			//title.Color = black

			// Change the position of the label.
			//title.Alignment = text.Middle

			// Draw the label to the graphics context.
			//title.Layout(gtx)

			// Pass the drawing operations to the GPU.

			if startButton.Clicked(gtx) {
				fmt.Print("Button Clicked")
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &startButton, "Start")
						return btn.Layout(gtx)
					},
				),
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			//refreshes window
			e.Frame(gtx.Ops)
		}
	}
}
