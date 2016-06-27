package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nguyenvanduocit/emoji"
)

type Post struct {
	Id int `json:"id"`
	ChapterSlug string `json:"chapter_slug"`
	PoemVi string `json:"poem_vi"`
}

type Response struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Emojis []*emoji.Emoji `json:"emojis"`
}


func FindEmoji(w http.ResponseWriter, r *http.Request){
	query := r.FormValue("q")
	response := &Response{
		Success:false,
		Message:"Unknown error!",
	}
	emojiList, err := emoji.FindEmoji(query)
	if err != nil {
		response.Message= err.Error()
	}else{
		response.Message= "Success"
		response.Success = true
		response.Emojis = emojiList
	}
	SendResponse(w, r, response)

}

func SendResponse(w http.ResponseWriter, r *http.Request, response *Response) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
	return
}

func main() {
	var ip, port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.StringVar(&port, "port", "8181", "Port")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", ip,port)
	fmt.Println("Server is listen on ", address);
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/findEmoji", FindEmoji)
	log.Fatal(http.ListenAndServe(address, router))
}
