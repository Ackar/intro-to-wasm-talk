package main

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"
)

type imageSearchComponent struct {
	searchText string

	lastUpdate      time.Time
	imgur           *Imgur
	searchInput     js.Value
	resultContainer js.Value
	textDiv         js.Value
}

func newImageSearchComponent(imgur *Imgur) *imageSearchComponent {
	i := &imageSearchComponent{
		imgur:           imgur,
		searchInput:     js.Global().Get("search"),
		resultContainer: js.Global().Get("result"),
		textDiv:         js.Global().Get("text"),
	}

	_ = i.searchInput.Call("addEventListener", "input", wrapFunc(i.onSearchUpdate))

	return i
}

func (i *imageSearchComponent) onSearchUpdate() {
	i.setText("")

	i.searchText = i.searchInput.Get("value").String()
	fmt.Printf("updated search: %s\n", i.searchText)

	if i.searchText == "" {
		i.resetResult()
		return
	}

	i.setText("Searching...")
	search := i.searchText
	// wait a little in case the user is still typing
	time.Sleep(500 * time.Millisecond)
	if i.searchText != search {
		return
	}

	fmt.Printf("searching '%s'\n", search)
	imgs, err := i.imgur.Search(i.searchText)
	if err != nil {
		fmt.Println(err)
		return
	}
	i.setText("")
	i.updateResults(imgs)
}

func (i *imageSearchComponent) updateResults(links []string) {
	if len(links) == 0 {
		i.setText("No results.")
	}

	var resValue []string
	for _, l := range links {
		resValue = append(resValue, fmt.Sprintf(`<div class="resimg"><a href="%s"><img src="%s" /></div></a>`, l, l))
	}
	v := strings.Join(resValue, "\n")
	i.resultContainer.Set("innerHTML", v)
}

func (i *imageSearchComponent) resetResult() {
	i.resultContainer.Set("innerHTML", "")
}

func (i *imageSearchComponent) setText(t string) {
	i.textDiv.Set("innerHTML", t)
}

func wrapFunc(f func()) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// start in a goroutine so that we can call blocking code (http request,
		// sleep...) inside the func
		go f()
		return nil
	})
}

func main() {
	imgur := NewImgur("79ae2f94f98a3c4")
	_ = newImageSearchComponent(imgur)

	select {}
}
