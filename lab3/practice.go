package main

import (
	"slices"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type counter struct {
	ctr   int
	label *widget.Label
	decr  *widget.Button
	incr  *widget.Button
	del   *widget.Button
}

func makeCounter() *counter {
	label := widget.NewLabel("0")

	ret := counter{
		0,
		label,
		nil,
		nil,
		nil,
	}

	decr := widget.NewButton("-", func() {
		ret.ctr -= 1
		label.SetText(strconv.Itoa(ret.ctr))
	})

	incr := widget.NewButton("+", func() {
		ret.ctr += 1
		label.SetText(strconv.Itoa(ret.ctr))
	})

	ret.decr = decr
	ret.incr = incr
   
	return &ret
}

func main() {
	a := app.New()
	w := a.NewWindow("Testing")

	entry := widget.NewEntry()
	ctrs := []*counter{}
	box := container.NewVBox()

	button := widget.NewButton("Create", func() {
		n, err := strconv.Atoi(entry.Text)

		if err == nil {
			box.RemoveAll()

			ctrs = nil
			for range n {
				newCtr := makeCounter()

				horizontal := container.NewHBox(newCtr.decr, newCtr.label, newCtr.incr)

				del := widget.NewButton("Delete", func() {
					box.Remove(horizontal)
					slices.DeleteFunc(ctrs, func(ctr *counter) bool {
						return ctr == newCtr
					})
				})

				newCtr.del = del
				ctrs = append(ctrs, newCtr)
				horizontal.Add(newCtr.del)

				box.Add(horizontal)
			}
		}
	})

	clr_button := widget.NewButton("Clear", func() {
		box.RemoveAll()
		ctrs = nil
	})

	w.SetContent(container.NewVBox(entry, button, clr_button, box))
	w.ShowAndRun()
}
