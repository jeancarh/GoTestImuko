package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Client struct {
	ClientId int     `json:"clientId"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	Tdc      string  `json:"tdc"`
	Monto    float64 `json:"monto"`
	Date     string  `json:"date"`
}

type CalculationRS struct {
	Total         float64 `json:"total"`
	ComprasPorTDC compras `json:"comprasPorTDC"`
	Nocompraron   int     `json:"nocompraron"`
	CompraMasAlta float64 `json:"compraMasAlta"`
}

type compras struct {
	Oro  float64 `json:"oro"`
	Amex float64 `json:"amex"`
}

func yourHandlerTs(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	days := v.Get("dias")
	if r.Method == "GET" {
		params := mux.Vars(r)
		date := params["date"]
		getDataApi(date, days, w)

	}
}

func main() {
	r := mux.NewRouter()
	r.Path("/resumen/{date}").Queries("dias", "{dias}").HandlerFunc(yourHandlerTs).Name("YourHandler").Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func getDataApi(date string, days string, w http.ResponseWriter) {
	const layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	intdays, _ := strconv.Atoi(days)
	var response []Client
	parsedDate = parsedDate.AddDate(0, 0, 0)
	var cli []Client
	for i := 0; i < intdays; i++ {
		resp, err := http.Get("https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/" + parsedDate.Format("2006-01-02"))
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		error := json.Unmarshal(body, &cli)
		if error != nil {
			log.Fatalln(error)
			fmt.Println(error)
		}
		for _, value := range cli {
			response = append(response, value)
		}
		parsedDate = parsedDate.AddDate(0, 0, 1)
	}
	data := returnCalculations(response)
	fmt.Fprintf(w, "%v\n", string(data))

}

func returnCalculations(resApi []Client) []byte {
	var suma float64
	var gold float64
	var amex float64
	ncomp := 0
	var max float64
	for _, value := range resApi {
		if value.Monto > 0 {
			suma += value.Monto
		}
		if value.Tdc == "amex" || value.Tdc == "amex corp" {
			amex = amex + 1
		}
		if value.Tdc == "gold" || value.Tdc == "visa gold" || value.Tdc == "master gold" {
			gold = gold + 1
		}
		if value.Compro == false {
			ncomp = ncomp + 1
		}
		if max < value.Monto {
			max = value.Monto
		}
	}
	compras := &compras{Oro: gold, Amex: amex}
	dataRS := &CalculationRS{
		Total:         suma,
		ComprasPorTDC: *compras,
		Nocompraron:   ncomp,
		CompraMasAlta: max}
	final, _ := json.Marshal(dataRS)
	return final
}
