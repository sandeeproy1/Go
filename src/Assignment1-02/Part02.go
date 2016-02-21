package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Print("\n****************************************************************************\n")
	var cap int
	fmt.Print("Enter the capacity of LRU Cache: ")
	fmt.Scanf("%d\n", &cap)
	lru := new(LRUCache)
	lru.initialize(cap)

	i := 1
	var oper int
	var key string
	var val string

	for i == 1 {
		fmt.Print("\nSelect an operation 1. Get 2. Set 3. Exit : ")
		fmt.Scanf("%d\n", &oper)

		if oper == 1 {
			fmt.Print("Enter key : ")
			fmt.Scanf("%s\r\n", &key)
			fmt.Print("The value is: ")
			fmt.Print(lru.get(key))
		} else if oper == 2 {
			fmt.Print("Enter key : ")
			fmt.Scanf("%s\r\n", &key)
			fmt.Print("Enter value : ")
			fmt.Scanf("%s\n", &val)
			lru.set(key, val)

		} else if oper == 3 {
			break
		} else {
			continue
		}
		lru.displaydll()
		lru.displaymap()
		fmt.Print("\n****************************************************************************\n")
	}
	//
	//for loop
	//Select an operation 1. Get 2. Set 3. Exit :
	//1. Enter key:
	//2. Enter key:
	//2. Enter value

	//var ln int
	//fmt.Scanf("%s", &ln)

	//var lru LRUCache

	//lru.set("fname", "sandeep")
	//lru.set("lname", "roy")
	//lru.set("age", "22")
	//lru.get("fname")
	//lru.set("age", "23")
	//lru.set("place", "manteca")
	//lru.displaydll()
	//lru.displaymap()
}

type LRUCache struct {
	dll      list.List
	hmap     map[string]string
	capacity int
	count    int
}

func (lru *LRUCache) initialize(cap int) {
	lru.dll = *list.New()
	lru.dll.Init()
	lru.hmap = make(map[string]string)
	lru.capacity = cap
	lru.count = 0
}

func (lru *LRUCache) get(key string) string {
	//If the item is present, return the value, else return -1
	_, ok := lru.hmap[key]
	var val string
	if ok == true {
		for e := lru.dll.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				lru.dll.MoveToFront(e)
				break
			}
		}
		val = lru.hmap[key]
		return val
	} else {
		return "-1"
	}
}

func (lru *LRUCache) set(key string, val string) {
	//Check if that value is there in the hash map
	_, ok := lru.hmap[key]
	//If present, update the hash map and keep that node at first in DLL
	if ok == true {
		// Iterate through list and move that element to front
		for e := lru.dll.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				lru.dll.MoveToFront(e)
			}
		}
	} else {
		//If not present, check if cache capacity is full
		if lru.count < lru.capacity {
			//If it is not full, add the new item to first
			lru.dll.PushFront(key)
			lru.count = lru.count + 1
		} else {
			//If full, remove the last item from the DLL and add the new item to first.
			lru.dll.PushFront(key)

			e := lru.dll.Back()
			delete(lru.hmap, e.Value.(string))

			lru.dll.Remove(lru.dll.Back())
		}
	}
	//Add the new item value to the hash map
	lru.hmap[key] = val
}

func (lru *LRUCache) displaymap() {
	fmt.Print("Hash Map=> ")
	for key, value := range lru.hmap {
		fmt.Print(key, ":", value)
		fmt.Print(", ")
	}
	fmt.Println("")
}
func (lru *LRUCache) displaydll() {
	// Iterate through list and print its contents.
	fmt.Print("\nDoubly Linked List=> ")
	for e := lru.dll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		fmt.Print(", ")
	}
	fmt.Println("")
}