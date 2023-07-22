package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/http"
)

func main() {
	ipEntry := widget.NewEntry()
	portEntry := widget.NewEntry()
	var ip string
	var port string
	button := widget.NewButton("Submit", func() {

		ip = ipEntry.Text
		port = portEntry.Text
	})
	a := app.New()
	w := a.NewWindow("go music")
	button1 := widget.NewButtonWithIcon("", theme.VolumeUpIcon(), func() {
		http.Get("http://" + ip + ":" + port + "/volumeUp")
	})

	button2 := widget.NewButtonWithIcon("", theme.VolumeDownIcon(), func() {
		http.Get("http://" + ip + ":" + port + "/volumeDown")
	})

	button3 := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		http.Get("http://" + ip + ":" + port + "/nextSong")
	})
	button4 := widget.NewButton("like", func() {
		http.Get("http://" + ip + ":" + port + "/like")
	})
	button5 := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		http.Get("http://" + ip + ":" + port + "/pause")
	})

	content := container.NewVBox(
		container.NewGridWithColumns(3,
			ipEntry,
			portEntry,
			button),
		container.NewGridWithColumns(2,
			button2,
			button1),
		container.NewGridWithColumns(2,
			button5,
			button3),
		button4,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
