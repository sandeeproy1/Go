package main

import (
	"container/list"
)

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

var dll = *list.New()

//dll.Init()
var hmap = make(map[int]int)
var count = 0

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