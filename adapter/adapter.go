package main

import "fmt"

// 适配器模式

type Player interface {
	PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (p MusicPlayer) PlayMusic() {
	fmt.Printf("playing music " + p.Src)
}


type VideoPlayer struct {
	Src string
}

func (v VideoPlayer) PlaySound() {
	fmt.Printf("playing sound " + v.Src)
}

type MusicAdaptor struct {
	SoundPlayer VideoPlayer
}

func (m MusicAdaptor) PlayMusic() {
	m.SoundPlayer.PlaySound()
}

func main() {
	videoPlayer := VideoPlayer{Src: "video.mid"}
	adapter := MusicAdaptor{SoundPlayer: videoPlayer}
	adapter.PlayMusic()
}
