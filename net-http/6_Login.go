package main

import (
	"io"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == "admin" && password == "pass"
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Welcome")
		} else {
			http.Error(w, "Bad login or password", http.StatusUnauthorized)
		}
		return 
	} else {
		io.WriteString(w, form)
	}
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(mainPage))
	if err != nil {
		panic(err)
	}
}