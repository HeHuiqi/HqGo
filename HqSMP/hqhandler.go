package HqSMP

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

var mm *MusicManager

var id int = 1
//var ctrl,signal chan int

func HandlerMMCommands(tokens []string)  {
	switch tokens[1] {
	case "list":
		for i :=0 ;i<mm.Len() ;i++  {
			mu,_ := mm.Get(i)
			fmt.Printf("%d: %s,%s,%s,%s\n",i+1,mu.Name,mu.Artist,mu.Source,mu.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			mm.Add(&MusicEntry{
				Id:strconv.Itoa(id),
				Name:tokens[2],
				Artist:tokens[3],
				Source:tokens[4],
				Type:tokens[5],
			})
		}else {
			fmt.Println("USAGE: mm add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			mm.RemoveByName(tokens[2])
		}else {
			fmt.Println("USAGE: mm remove ")

		}
	default:
		fmt.Println("Unrecognized mm command:",tokens[1])


	}
}

func HandlerPlayCommand(tokens []string)  {

	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := mm.Find(tokens[1])
	if e == nil {
		fmt.Println("The music ",tokens[1],"does not exit")
		return
	}
	 Play(e.Source,e.Type)
}

func HandlerMain()  {
	fmt.Println(`
Enter following commands to control the player:
mm list -- View the existing music lib
mm add <name><artist><source><type> -- Add a music to the music lib
mm remove <name> -- Remove the specified music from the lib
play <name> -- Play the specified music
`)
	mm = NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for  {

		fmt.Print("$>")
		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line," ")
		if tokens[0] == "play" {
			HandlerPlayCommand(tokens)
		}else if tokens[0] == "mm" {
			HandlerMMCommands(tokens)
		}else {
			fmt.Println("Unrecognized mm command:",tokens[0])
		}

	}
}