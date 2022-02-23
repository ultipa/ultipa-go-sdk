package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"
)


type respStruct struct {
	Code  int
	Data struct{
		Token string
		Username string
	}
}

func RequestPost(host, url string, form url.Values, token string) (body []byte, err error) {

	client := &http.Client{}

	URL := fmt.Sprintf(`http://%v/%v`, host, url)
	req, err := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	cookie := http.Cookie{
		Name: "ultipaManger",
		Value: token,
	}
	req.AddCookie(&cookie)
	resp, err := client.Do(req)

	if err != nil {
		return nil , err
	}

	body, err = ioutil.ReadAll(resp.Body)

	return body, err
}

func TestAPI(t *testing.T) {

	form := url.Values{}
	form.Add("_host", "172.31.24.174:60061")
	form.Add("username", "root")
	form.Add("password", "root")
	resp, err := http.PostForm("http://69.230.214.186:3032/api/user/login", form)

	bs, err := ioutil.ReadAll(resp.Body)

	r := &respStruct{}
	json.Unmarshal(bs,r)


	log.Println(r, err)


	token := r.Data.Token

	form1 := url.Values{}

	form1.Add("uql", "find().nodes() return nodes limit 1")


	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {

		go func(i int) {

			defer wg.Done()
			_, err := RequestPost("69.230.214.186:3032","api/uql?_host=172.31.24.174:60061&_gn=twitter&_t=60&_nc=true", form1, token)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(i)
		}(i)
	}

	wg.Wait()
	log.Println(time.Since(start))





}