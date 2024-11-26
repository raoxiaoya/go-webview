package main

import (
	"errors"
	"log"

	webview2 "github.com/jchv/go-webview2"
)

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			Width:  800,
			Height: 600,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()

	w.Bind("test_love", Great)

	w.Navigate("file:///D:/dev/php/magook/trunk/server/go-webview/test.html")
	// w.Navigate("https://github.com/avelino/awesome-go#gui")

	w.Run()
}

func Great(param string) (result string, err error) {
	if param == "I love you" {
		return "I love you too", nil
	} else if param == "I hate you" {
		return "", errors.New("break up")
	} else {
		return "I don't know", nil
	}
}
