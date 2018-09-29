package HqSMP

import (
	"fmt"
	"time"
	"strings"
)

type Player interface {
	 Play(source string)
}

func Play(source string,mtype string)  {
	mtype = strings.ToUpper(mtype)

	var pl Player
	switch mtype {
	case "MP3":
		pl = &MP3Player{}
	case "AVI":
		pl = &AVIPlayer{}
	default:
		fmt.Println("Unsupported music tupe",mtype)
		return
	}
	pl.Play(source)
}

type MP3Player struct {
	state int
	progress int
}

func (p *MP3Player) Play(source string)  {
	fmt.Println("Playing MP3",source)
	p.progress = 0
	for p.progress < 1000 {
		time.Sleep(100*time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing",source)
}

type AVIPlayer struct {
	state int
	progress int
}

func (p *AVIPlayer) Play(source string)  {
	fmt.Println("Playing AVI",source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100*time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing",source)
}