package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/itchyny/volume-go"
	"net"
	"net/http"
	"strings"
)

const (
	KeyLike = "l"
)
const (
	AudioVolDown = "audio_vol_down"
	AudioVolUp   = "audio_vol_up"
	AudioPlay    = "audio_play"
	AudioNext    = "audio_next"
)

func audioGetVol() int {
	volume, err := volume.GetVolume()
	if err != nil {
		fmt.Println("vol err: ", err)
	}
	return volume
}
func audioPlay(play string) {
	robotgo.KeyTap(play)
}
func keyOperate(key string) {
	robotgo.KeyTap(key, "ctrl", "alt")
}
func handleVolumeUp(w http.ResponseWriter, r *http.Request) {
	audioPlay(AudioVolUp)
	volume := audioGetVol()
	fmt.Fprintf(w, "Volume increased. Current volume: %d", volume)
}

func handleVolumeDown(w http.ResponseWriter, r *http.Request) {
	audioPlay(AudioVolDown)
	volume := audioGetVol()
	fmt.Fprintf(w, "Volume decreased. Current volume: %d", volume)
}

func handleNextSong(w http.ResponseWriter, r *http.Request) {
	audioPlay(AudioNext)
	fmt.Fprintf(w, "Switched to the next song.")
}
func handlePause(w http.ResponseWriter, r *http.Request) {
	audioPlay(AudioPlay)
	fmt.Fprintf(w, "the song pause.")
}
func handleLike(w http.ResponseWriter, r *http.Request) {
	keyOperate(KeyLike)
	fmt.Fprintf(w, "like this song.")
}

func getOutBoundIP() string {
	connect, err := net.Dial("udp", "8.8.4.4:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	addr := connect.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]

	return ip
}
func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,this is the song control server."))
	})
	http.HandleFunc("/volumeUp", handleVolumeUp)
	http.HandleFunc("/volumeDown", handleVolumeDown)
	http.HandleFunc("/nextSong", handleNextSong)
	http.HandleFunc("/pause", handlePause)
	http.HandleFunc("/like", handleLike)
	ip := getOutBoundIP()
	fmt.Println("Server is running on " + ip + ":3000")
	http.ListenAndServe(":3000", nil)
}
