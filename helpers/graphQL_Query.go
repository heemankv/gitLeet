package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


const BaseURL string = "https://leetcode.com/graphql/"
func call(query, cookie string) (response []byte, err error) {


  method := "POST"

  payload := strings.NewReader(query)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, BaseURL, payload)

  if err != nil {
    fmt.Println(err)
		return nil, err
  }
  req.Header.Add("authority", "leetcode.com")
  req.Header.Add("accept", "*/*")
  req.Header.Add("accept-language", "en-US,en;q=0.9")
  req.Header.Add("authorization", "")
  req.Header.Add("content-type", "application/json")
  req.Header.Add("cookie", cookie)
  req.Header.Add("sec-fetch-mode", "cors")
  req.Header.Add("sec-fetch-site", "same-origin")

  res, err := client.Do(req)
  if err != nil {
    // fmt.Println(err) TODO: bahar kro yeh
    return nil, err
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  // fmt.Println(string(body))  TODO: bahar kro yeh
	return body, nil
}

func QueryWrapper(query, cookie string)(value map[string]interface{}){
	response, err := call(query, cookie)
	if err != nil {
    fmt.Println(err)
		return // TODO: Empty List
  }
	var ql interface{} // interface for decoding arb data
	eror := json.Unmarshal(response, &ql)

	if eror != nil {
    fmt.Println(eror)
		return // TODO: Empty List
  }

	value = ql.(map[string]interface{})
	return
}

