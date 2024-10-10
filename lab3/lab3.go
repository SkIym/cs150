package main

import (
	"fmt"
	"image/color"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PuerkitoBio/goquery"
)

type source struct {
	label  *widget.Label
	status *canvas.Text
	url    string
}

func makeSource(label string, url string) source {

	source := source{
		widget.NewLabel(label),
		canvas.NewText("Fetching", color.RGBA{255, 127, 127, 255}),
		url,
	}

	source.status.TextSize = 60

	return source
}

const countdownDuration time.Duration = 30

const fetchTimeout int = 8

func crawler(urlChannel chan string, site *source, term string) {

	fmt.Println("Fetching", site.label.Text)

	done := make(chan struct{})

	go func() {
		select {
			case <-time.After(time.Second * time.Duration(fetchTimeout)):
				site.status.Text = "Timeout"
				fmt.Println(site.label.Text + " timed out")
				site.status.Color = color.RGBA{0, 0, 255, 255}
				urlChannel <- "timeout"
				close(done)	
			case <-done:
				return
		}
	}()

	resp, err := http.Get("https://app.requestly.io/delay/5000/" + site.url)
	if err != nil || resp.StatusCode != 200 {
		site.status.Text = "Error"
		site.status.Color = color.RGBA{0, 0, 255, 255}
		close(done)
		urlChannel <- "fail"
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		site.status.Text = "Error"
		site.status.Color = color.RGBA{0, 0, 255, 255}
		close(done)
		urlChannel <- "fail"
	}

	select {
		case <-done:
			return
		default:
			found := false
			doc.Find("li").Each(func(i int, p *goquery.Selection) {
				select {
				case <-done: // Check if the timeout has occurred
					return // Exit the loop if timeout is reached
				default:
					if strings.Contains(p.Text(), term) {
						found = true
					}
				}
			})

			if !found {
				site.status.Text = "Not Suspended"
				site.status.Color = color.RGBA{127, 127, 127, 255}
				fmt.Println("not suspended accrdng to", site.label.Text)
				close(done)
				urlChannel <- "fail"
			} else {
				site.status.Text = "Suspended"
				site.status.Color = color.RGBA{255, 0, 0, 255}
				fmt.Println("Suspended accrdng to", site.label.Text)
				close(done)
				urlChannel <- "done"
			}
	}

	// found := false
	// doc.Find("li").Each(func(i int, p *goquery.Selection) {
	// 	select {
	// 	case <-done: // Check if the timeout has occurred
	// 		return // Exit the loop if timeout is reached
	// 	default:
	// 		if strings.Contains(p.Text(), term) {
	// 			found = true
	// 		}
	// 	}
	// })

	// if !found {
	// 	select {
	// 		case <-done: // Check if the timeout has occurred
	// 			return // Exit the loop if timeout is reached
	// 		default:
	// 			site.status.Text = "Not Suspended"
	// 			site.status.Color = color.RGBA{127, 127, 127, 255}
	// 			fmt.Println("not suspended accrdng to", site.label.Text)
	// 			close(done)
	// 			urlChannel <- "fail"
	// 	}
	// } else {
	// 	select {
	// 		case <-done: // Check if the timeout has occurred
	// 			return // Exit the loop if timeout is reached
	// 		default:
	// 			site.status.Text = "Suspended"
	// 			site.status.Color = color.RGBA{255, 0, 0, 255}
	// 			fmt.Println("Suspended accrdng to", site.label.Text)
	// 			close(done)
	// 			urlChannel <- "done"
	// 	}
	// }
	site.status.Refresh()
}

func check(sources []source, searchTerm string, urlChannel chan string) {
	wg := sync.WaitGroup{}

	for i := range sources {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			crawler(urlChannel, &sources[index], searchTerm)
		}(i)
	}

	go func() {
		wg.Wait()
		close(urlChannel)
	}()
}

func main() {
	a := app.New()
	w := a.NewWindow("Fetcher")
	box := container.NewVBox()
	countdown := widget.NewLabel(strconv.Itoa(int(countdownDuration)))
	search := widget.NewEntry()
	search.SetText("Bulacan")
	box.Add(countdown)
	box.Add(search)

	raw, err := os.ReadFile("urls.csv")

	if err != nil {
		return
	}

	text := string(raw)

	lines := strings.Split(text, "\n")
	sources := []source{}

	// Decode urls.csv, initialize source structs and widgets
	for _, s := range lines {
		line := strings.Split(s, ",")
		label, url := line[0], line[1]
		site := makeSource(label, url)
		sources = append(sources, site)
		horizontal := container.NewHBox(site.label, site.status)
		box.Add(horizontal)
		// fmt.Println(label, url)
	}


	// code below this line should loop

	go func() {
		for {
			sec := time.NewTimer(time.Second * 1)
			<-sec.C
			timeRemaining, _ := strconv.Atoi(countdown.Text)
			if timeRemaining == 0 {
				countdown.SetText(strconv.Itoa(int(countdownDuration)))
			} else {
				timeRemaining -= 1
				countdown.SetText(strconv.Itoa(timeRemaining))
			}
		}
	}()

	urlChannel := make(chan string)

	check(sources, search.Text, urlChannel)

	go func() {
		for {

			results := []string{}
			for result := range urlChannel {
				// fmt.Println(countdown.Text)
				// fmt.Println(result)
				results = append(results, result)
			}

			pause := time.NewTimer(time.Second * 3)
			<-pause.C

			countdown.SetText(string(countdownDuration))
			urlChannel = make(chan string)

			for i := range sources {
				sources[i].status.Text = "Fetching"
				sources[i].status.Color = color.RGBA{255, 127, 127, 255}
			}
			check(sources, search.Text, urlChannel)
			fmt.Println(results)
		}

	}()

	w.SetContent(container.NewVBox(box))
	w.ShowAndRun()
}
