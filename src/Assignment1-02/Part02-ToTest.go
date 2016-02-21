package main

import (
	"container/list"
	"fmt"
)

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

var dll = *list.New()

//dll.Init()
var hmap = make(map[int]int)
var count = 0

func main() {
	//assert := assert.New(t)
	fmt.Println("Lab 1 - Part II \n---- LRU Cache ----")
	// Populates entries into the Cache
	Set(1, 10)
	Set(2, 20)
	Set(3, 30)
	fmt.Println(Get(1))
	fmt.Println(Get(2))
	fmt.Println(Get(3))
	Set(4, 40)
	fmt.Println(Get(4))
	fmt.Println(Get(1))
	// Checks Cache Invalidation
	//assert.Equal(-1, Get(1), "Get(1) should return -1 as it was the least recently used entry")
}

func Set(key int, value int) {
	// TODO: add your code here!
	if dll.Len() <= 0{
	dll.Init()
	}	
	//Check if that value is there in the hash map
	_, ok := hmap[key]
	//If present, update the hash map and keep that node at first in DLL
	if ok == true {
		// Iterate through list and move that element to front
		for e := dll.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				dll.MoveToFront(e)
			}
		}
	} else {
		//If not present, check if cache capacity is full
		if count < CACHE_SIZE {
			//If it is not full, add the new item to first
			dll.PushFront(key)
			count = count + 1
		} else {
			//If full, remove the last item from the DLL and add the new item to first.
			dll.PushFront(key)

			e := dll.Back()
			delete(hmap, e.Value.(int))

			dll.Remove(dll.Back())
		}
	}
	//Add the new item value to the hash map
	hmap[key] = value
}

func Get(key int) int {
	// TODO: add your code here!
	//If the item is present, return the value, else return -1
	_, ok := hmap[key]
	var val int
	if ok == true {
		for e := dll.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				dll.MoveToFront(e)
				break
			}
		}
		val = hmap[key]
		return val
	} else {
		return -1
	}
}