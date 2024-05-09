package main

func constructLL(arr []int) {
	list := &LinkedList{}

	for _, v := range arr {
		list.Insert(v)
	}

	list.Display()
}
