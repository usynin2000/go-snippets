package main

import (
	"fmt"
	"net/http"
)

func reqUrlQuery(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header =================\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	
	body += "Query parameters =========\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	
	res.Write([]byte(body))
}

func parseFormPage(res http.ResponseWriter, req *http.Request) {
    body := fmt.Sprintf("Method: %s\r\n", req.Method)
    body += "Header ===============\r\n"
    for k, v := range req.Header {
        body += fmt.Sprintf("%s: %v\r\n", k, v)
    }
    body += "Query parameters ===============\r\n"
    if err := req.ParseForm(); err != nil {
        res.Write([]byte(err.Error()))
        return
    }
    for k, v := range req.Form {
        body += fmt.Sprintf("%s: %v\r\n", k, v)
    }
    res.Write([]byte(body))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("/"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/parse", parseFormPage)
	mux.HandleFunc("/query", reqUrlQuery)
	mux.HandleFunc("/", mainPage)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// запусти http://localhost:8080/?id=12345&name=John%20Doe&filter=town&filter=country.


// req.URL.Query()
// Берёт параметры только из URL (часть после ?).

// GET /query?name=Tom&age=20 → будут распарсены name, age.


// req.ParseForm()
// объединяет данные из query string и тела формы (
// если Content-Type подходит: application/x-www-form-urlencoded и POST.
// req.Form = URL.Query() + тело запроса.

// curl --location 'http://localhost:8080/parse?test_2=info2' \
// --header 'Content-Type: application/x-www-form-urlencoded' \
// --data-urlencode 'test=info'

//  Вот так например будет и квери и body через urlencode принимать