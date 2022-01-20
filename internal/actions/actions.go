package actions

var store map[int64]int = make(map[int64]int)

func Set(id int64, action int) {
	store[id] = action
}

func If(id int64, action int) bool {
	return store[id] == action
}

func Clear(id int64) {
	delete(store, id)
}
