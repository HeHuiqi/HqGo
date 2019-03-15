package hqweb


import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"io/ioutil"
)

type msg struct {
	Num int
}

func HqWebSocket() {

	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)
	//设置静态文件目录
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./static/websocket.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	if conn != nil {
		go echo(conn)
	}

}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
			conn.Close()
			return
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}