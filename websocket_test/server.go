package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// クライアント一覧
var clients = make(map[*websocket.Conn]bool)

// メッセージ送信用チャネル
var msgCh = make(chan string)

// index.html
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatalln(err)
	}
}

// WebSocket
func webSocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	err := websocket.Message.Send(ws, "接続しました")
	if err != nil {
		log.Fatalln(err)
	}

	// クライアント登録
	clients[ws] = true

	for {
		msg := ""
		// メッセージ受信
		err = websocket.Message.Receive(ws, &msg)
		if err != nil {
			// 失敗したらクライアント一覧から削除
			delete(clients, ws)
			break
		}
		// チャネルにメッセージ追加
		msgCh <- msg
	}
}

// メッセージ送信処理
func handleMessages() {
	for {
		// チャネルからメッセージ取り出し
		msg := <-msgCh
		fmt.Println("recieve " + msg)
		// 全クライアントにメッセージ送信
		for ws := range clients {
			err := websocket.Message.Send(ws, msg)
			if err != nil {
				// 送信失敗したクライアントは削除
				ws.Close()
				delete(clients, ws)
			}
		}
	}

}

func main() {
	// index.html
	http.HandleFunc("/", indexHandler)
	// WebSocket
	http.Handle("/ws", websocket.Handler(webSocketHandler))
	// メッセージ送信処理
	go handleMessages()

	http.ListenAndServe(":8000", nil)
}
