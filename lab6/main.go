package main

import (
	"errors"
	"fmt"
)

type Node struct {
	key   string
	Value interface{}
}

type HashMap struct {
	size    int
	count   int
	buckets [][]Node
}

func (h *HashMap) getIndex(key string) int {
	return int(hash(key)) % h.size
}

// Implements the Jenkins hash function
func hash(key string) uint32 {
	var h uint32
	for _, c := range key {
		h += uint32(c)
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return h
}

func hashAlter(key string) int {
	hashValue := 0
	for _, char := range key {
		hashValue += int(char)
	}
	return hashValue % len(key)
}

func (h *HashMap) Len() int {
	return h.count
}

func (h *HashMap) Size() int {
	return h.size
}

// NewHashMap is the constructor that returns a new HashMap of a fixed size
// returns an error when a size of 0 is provided
func NewHashMap(size int) (*HashMap, error) {
	h := new(HashMap)
	if size < 1 {
		return h, errors.New("размер мапы должен быть > 1")
	}
	h.size = size
	h.count = 0
	h.buckets = make([][]Node, size)
	for i := range h.buckets {
		h.buckets[i] = make([]Node, 0)
	}
	return h, nil
}

func (h *HashMap) Get(key string) (*Node, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]
	for _, node := range chain {
		if node.key == key {
			return &node, true
		}
	}
	return nil, false
}

func (h *HashMap) Set(key string, value interface{}) bool {
	index := h.getIndex(key)
	chain := h.buckets[index]
	found := false

	for i := range chain {
		node := &chain[i]
		if node.key == key {
			node.Value = value
			found = true
		}
	}
	if found {
		return true
	}

	if h.size == h.count {
		return false
	}

	node := Node{key: key, Value: value}
	chain = append(chain, node)
	h.buckets[index] = chain
	h.count++

	return true
}

func (h *HashMap) Delete(key string) (*Node, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]

	found := false
	var location int
	var mapNode *Node

	for loc, node := range chain {
		if node.key == key {
			found = true
			location = loc
			mapNode = &node
		}
	}

	if found {
		h.count--
		N := len(chain) - 1
		chain[location], chain[N] = chain[N], chain[location]
		chain = chain[:N]
		h.buckets[index] = chain
		return mapNode, true
	}

	return nil, false
}

func (h *HashMap) Load() float32 {
	return float32(h.count) / float32(h.size)
}

func main() {
	myMap, err := NewHashMap(10)
	if err != nil {
		fmt.Println(err)
	}
	myMap.Set("ключ1", 123)
	myMap.Set("ключ1", 125)
	myMap.Set("ключ1", 126)

	fmt.Println(myMap.Get("ключ1"))

	myMap.Delete("ключ1")

	fmt.Println(myMap.Get("ключ1"))
}
