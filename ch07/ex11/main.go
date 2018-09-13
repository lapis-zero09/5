package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollar float64

type Item struct {
	Name  string
	Price dollar
}

type memoryDB struct {
	mu     sync.Mutex
	nextID int
	Items  map[int]*Item
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		Items: map[int]*Item{
			1: &Item{Name: "shoes", Price: 50.},
			2: &Item{Name: "socks", Price: 5.},
		},
		nextID: 3,
	}
}

func main() {
	db := newMemoryDB()
	http.HandleFunc("/items", db.ListItems)
	http.HandleFunc("/item", db.GetItem)
	http.HandleFunc("/update", db.UpdateItem)
	http.HandleFunc("/create", db.AddItem)
	http.HandleFunc("/delete", db.DeleteItem)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.Items = nil
}

func (db *memoryDB) GetItem(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()
	strID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "id param must be int", http.StatusBadRequest)
		return
	}
	item, ok := db.Items[id]
	if !ok {
		err := fmt.Sprintf("memorydb: item not found with ID %d", id)
		http.Error(w, err, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%v", *item)
}

func (db *memoryDB) ListItems(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var items []*Item
	for _, b := range db.Items {
		items = append(items, b)
	}

	for _, item := range items {
		fmt.Fprintf(w, "%v", *item)
	}
}

func (db *memoryDB) AddItem(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var item Item
	name := req.URL.Query().Get("name")
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "price param must be int", http.StatusBadRequest)
		return
	}
	item.Name = name
	item.Price = dollar(price)

	db.Items[db.nextID] = &item

	db.nextID++

	fmt.Fprintln(w, "successfully added")
}

func (db *memoryDB) DeleteItem(w http.ResponseWriter, req *http.Request) {
	strID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "id param must be int", http.StatusBadRequest)
		return
	}

	if id == 0 {
		err := "memorydb: item with unassigned ID passed into deleteItem"
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Items[id]; !ok {
		http.Error(w, "memorydb: could not delete item with ID %d, does not exist", id)
		return
	}
	delete(db.Items, id)
	fmt.Fprintln(w, "successfully deleted")
}

// UpdateItem updates the entry for a given item.
func (db *memoryDB) UpdateItem(w http.ResponseWriter, req *http.Request) {
	strID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "id param must be int", http.StatusBadRequest)
		return
	}

	if id == 0 {
		err := "memorydb: item with unassigned ID passed into updateItem"
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	var item Item
	name := req.URL.Query().Get("name")
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "price param must be int", http.StatusBadRequest)
		return
	}

	if name != "" {
		item.Name = name
	}
	if price > 0. {
		item.Price = dollar(price)
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.Items[id] = &item
	fmt.Fprintln(w, "successfully updated")
}
