package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/k0kubun/pp"
)

var (
	APPSTORE_ID   string = "493470467"
	PLAYSTORE_URL string = ""
)

const (
	APPSTORE_URL string = "http://itunes.apple.com/jp/rss/customerreviews/id=%s/json"
)

func main() {
	//	if len(os.Args) == 1 {
	//		cmd := exec.Command(os.Args[0], "--child")
	//		cmd.Start()
	//	} else {
	//		tick := time.Tick(3 * time.Second)
	//		for {
	//			select {
	//			case _ = <-tick:
	//				now := time.Now()
	//				content := []byte(now.String())
	//				ioutil.WriteFile("/tmp/output", content, os.ModePerm)
	//			}
	//		}
	//	}
	GetAppStoreReview()
}

func GetAppStoreReview() {
	url := fmt.Sprintf(APPSTORE_URL, APPSTORE_ID)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Errorf("error")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("error")
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	var res ITunesResponse
	json.Unmarshal(val, &res)
	for _, value := range res.Feed.Entry {
		fmt.Print("content:")
		pp.Println(value.Title.Label)
		fmt.Print("name:")
		pp.Println(value.Author.Name.Label)
	}
	//pp.Println(res)
}

func GetPlayStoreReview() {

}

func Notify() {

}
