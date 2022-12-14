package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {

	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {

		fmt.Println(err.Error())
		return 
	}

	if resp.StatusCode != 200 {

		fmt.Println("Not Sucess", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var resposta []User
	err = json.Unmarshal(body, &resposta)
	if err != nil {
		fmt.Println("Erro ao recuperar dados", err.Error())
		return
	
	}

	fmt.Println(resposta)
}