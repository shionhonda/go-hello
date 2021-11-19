package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}
	db = database{item: dollars(price)}
	fmt.Fprintf(w, "created %s: %s\n", item, db[item])
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "Please select from below:\n")
		db.list(w, req)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

// e.g., http://localhost:8000/update?item=fuga&price=3.14
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "updated %s: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	delete(db, item)
	fmt.Fprintf(w, "deleted %s: %s\n", item, db[item])
}
