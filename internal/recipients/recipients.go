package recipients

var store map[int64]int64 = make(map[int64]int64)

func Set(sender int64, recipient int64) {
	store[sender] = recipient
}

func Get(sender int64) (recipient int64) {
	recipient = store[sender]
	delete(store, sender)
	return
}

func Clear(sender int64) {
	delete(store, sender)
}
