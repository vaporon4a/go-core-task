package main

import "fmt"

const (
	initialArraySize = 5
	maxLoadFactor    = 0.7
)

type StringIntMap struct {
	array     []*bucket
	size      int
	arraySize int
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func NewHashMap() *StringIntMap {
	h := &StringIntMap{
		array:     make([]*bucket, initialArraySize),
		size:      0,
		arraySize: initialArraySize,
	}
	for i := range h.array {
		h.array[i] = &bucket{}
	}
	return h
}

// hash takes a string and an integer size as arguments and returns the hash value
// of the string modulo the given size.
//
// The hash value is calculated by summing the Unicode code points of the
// characters in the string and taking the result modulo the given size.
func hash(key string, size int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

// add inserts a new key-value pair into the bucket if the key does not already exist.
// If the key exists, it updates the value. Returns true if a new node was added,
// and false if an existing node was updated.
func (b *bucket) add(key string, value int) bool {
	if b.exist(key) {
		node := b.head
		for node != nil {
			if node.key == key {
				node.value = value
				return false
			}
			node = node.next
		}
	} else {
		node := &bucketNode{
			key:   key,
			value: value,
			next:  b.head,
		}
		b.head = node
		return true
	}
	return false
}

// remove removes the node with the given key from the bucket.
// It returns true if a node was removed, and false if the key was not found.
func (b *bucket) remove(key string) bool {
	if b.head == nil {
		return false
	}

	if b.head.key == key {
		b.head = b.head.next
		return true
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
			return true
		}
		prevNode = prevNode.next
	}
	return false
}

// exist checks if a node with the given key exists in the bucket.
// It returns true if the key exists, and false if the key does not exist.
func (b *bucket) exist(key string) bool {
	node := b.head
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

// resize doubles the size of the array in the StringIntMap and rehashes
// all existing key-value pairs into the new array. This function is
// typically called when the load factor exceeds the maximum threshold to
// maintain efficient operations. It creates a new array with double the
// current size, initializes each bucket, and redistributes the nodes from
// the old array into the new one based on the new hash indexes.
func (m *StringIntMap) resize() {
	newArraySize := m.arraySize * 2
	newArray := make([]*bucket, newArraySize)
	for i := range newArray {
		newArray[i] = &bucket{}
	}

	for i := range m.array {
		node := m.array[i].head
		for node != nil {
			index := hash(node.key, newArraySize)
			newArray[index].add(node.key, node.value)
			node = node.next
		}
	}
	m.array = newArray
	m.arraySize = newArraySize
}

// Add inserts a key-value pair into the map. If the key already exists, it updates the value.
// When a new key-value pair is added, the size of the map increases. If the load factor exceeds
// the maximum threshold, the map is resized to maintain efficient operations.
func (m *StringIntMap) Add(key string, value int) {
	index := hash(key, m.arraySize)
	if m.array[index].add(key, value) {
		m.size++
		if float64(m.size)/float64(m.arraySize) > maxLoadFactor {
			m.resize()
		}
	}
}

// Remove deletes a key-value pair from the map corresponding to the given key.
// If the key is found and removed, the size of the map is decremented.
// If the key does not exist, the map remains unchanged.
func (m *StringIntMap) Remove(key string) {
	index := hash(key, m.arraySize)
	if m.array[index].remove(key) {
		m.size--
	}

}

// Copy creates a deep copy of the StringIntMap, including all its buckets and node elements.
// It returns a new instance of StringIntMap with the same key-value pairs and structure
// as the original map, ensuring that modifications to the copy do not affect the original map.
func (m *StringIntMap) Copy() *StringIntMap {
	newMap := &StringIntMap{
		array:     make([]*bucket, m.arraySize),
		size:      m.size,
		arraySize: m.arraySize,
	}
	for i := range newMap.array {
		newMap.array[i] = m.array[i].copy()
	}
	return newMap
}

// copy creates a deep copy of the bucket, including all its node elements.
// It returns a new instance of bucket with the same key-value pairs and structure
// as the original bucket, ensuring that modifications to the copy do not affect the original bucket.
func (b *bucket) copy() *bucket {
	newBucket := &bucket{}
	node := b.head
	for node != nil {
		newBucket.add(node.key, node.value)
		node = node.next
	}
	return newBucket

}

// Exist checks if a key exists in the map.
// It returns true if the key is found, and false otherwise.
func (m *StringIntMap) Exist(key string) bool {
	index := hash(key, m.arraySize)
	return m.array[index].exist(key)
}

// Get retrieves the value associated with the given key in the map.
// It returns the value and true if the key is found, and zero and false otherwise.
func (m *StringIntMap) Get(key string) (int, bool) {
	index := hash(key, m.arraySize)
	return m.array[index].get(key)
}

// get retrieves the value associated with the given key in the bucket.
// It returns the value and true if the key is found, and zero and false otherwise.
func (b *bucket) get(key string) (int, bool) {
	node := b.head
	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}
	return 0, false
}

func main() {
	m := NewHashMap()
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)
	m.Add("key4", 4)
	m.Add("key5", 5)
	m.Add("key6", 6)
	m.Add("key7", 7)
	m.Add("key8", 8)
	m.Add("key9", 9)
	m.Add("key10", 10)
	m.Add("key11", 11)
	m.Add("key12", 12)
	m.Add("key13", 13)
	m.Add("key14", 14)
	m.Add("key15", 15)
	m.Add("key16", 16)

	fmt.Println(m.Exist("key1"))

	fmt.Println(m.Get("key1"))

	m.Remove("key1")

	fmt.Println(m.Get("key1"))

	fmt.Println(m.Copy())

}
