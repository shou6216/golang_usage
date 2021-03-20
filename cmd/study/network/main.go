package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Person2 struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	// json:"-" 非表示にする
	// json:"hoge,omitempty" 何も入ってない時表示しない
	Nicknames []string `json:nicknames,omitempty`
}

func (p Person2) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func main() {
	base, _ := url.Parse("http://example.com")
	reference73, _ := url.Parse("/test/a=1&b=2")
	endpoint73 := base.ResolveReference(reference73).String()
	fmt.Println(endpoint73)

	req73, _ := http.NewRequest("GET", endpoint73, nil)
	req73.Header.Add("If-None-Match", "W/wyzzy")
	q73 := req73.URL.Query()
	q73.Add("c", "3&%")
	fmt.Println(q73)
	fmt.Println(q73.Encode())

	var client73 *http.Client = &http.Client{}
	resp73, _ := client73.Do(req73)
	body73, _ := ioutil.ReadAll(resp73.Body)
	fmt.Println(string(body73))

	b74 := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p74 Person2
	if err74 := json.Unmarshal(b74, &p74); err74 != nil {
		fmt.Println(err74)
	}
	fmt.Println(p74.Name, p74.Age, p74.Nicknames)

	v74, _ := json.Marshal(p74)
	fmt.Println(string(v74))
}
