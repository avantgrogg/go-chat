package main

import "net/http"
import "log"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
            <html>
                <head>
                    <title>Chat</title>
                </head>
                <body>
                    Lets chat!
                </body>
            </html>
        `))
	})
	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
