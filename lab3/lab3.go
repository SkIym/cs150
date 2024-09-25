package main

import (
	"fmt"
	// "image/color"
	// "log"
	// "net/http"
	"os"
	"strings"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type source struct {
	label  *widget.Label
	status *widget.Label
}

func makeSource(label string, url string) source {

	source := source{
		widget.NewLabel(label),
		nil,
	}

	status := widget.NewLabel("Fetching")
	// status.Color = color.RGBA{255, 127, 127, 255}

	source.status = status

	return source
}

var countdownDuration int = 30

func main() {
	a := app.New()
	w := a.NewWindow("Your window title here")
	box := container.NewVBox()

	raw, err := os.ReadFile("urls.csv")

	if err != nil {
		return
	}

	text := string(raw)

	lines := strings.Split(text, "\n")
	for _, s := range lines {
		line := strings.Split(s, ",")
		label, url := line[0], line[1]
		source := makeSource(label, url)
		horizontal := container.NewHBox(source.label, source.status)
		box.Add(horizontal)
		fmt.Println(label, url)
	}

	// tIMER LOGIC
	// timer := time.NewTimer(time.Second * 1)

	// go func() {
	// 	var currentTime int
	// 	var x int

	// 	for {
	// 		fmt.Println(currentTime)
	// 		timer = time.NewTimer(time.Second * 1)
	// 		<-timer.C

	// 		currentTime, _ = strconv.Atoi(display.timeLeft.Text)

	// 		if currentTime == 1 {
	// 			x, _ = strconv.Atoi(display.counter.Text)
	// 			x += strdisplay.entry.text
	// 			display.counter.SetText(strconv.Itoa(x))
	// 			currentTime = countdownDuration
	// 		} else {
	// 			currentTime -= 1
	// 		}

	// 		display.timeLeft.SetText(strconv.Itoa(currentTime))

	// 	}
	// }()

	// FETCHING PARSING

	// resp, err := http.Get("https://dcs.upd.edu.ph")
	// if err != nil || resp.StatusCode != 200 {
	// 	log.Fatalln("URL fetch failed")
	// }

	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	log.Fatalln("HTTP parsing failed")
	// }

	// found := false

	// doc.Find("a").Each(func(i int, p *goquery.Selection) {
	// 	if strings.Contains(p.Text(), "Diliman") {
	// 		found = true
	// 	}
	// })

	// fmt.Printf("Is there an `a` tag containing \"Diliman\"? %v\n", found)

	w.SetContent(container.NewVBox(box))
	w.ShowAndRun() // Start Fyne's event loop
}
