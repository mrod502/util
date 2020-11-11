package util

//Comparator - something with a less func
type Comparator interface {
	Key() string
}

type node struct {
	Left    *node
	Right   *node
	Current Comparator
}

func (n node) search(key string) (res *node, found bool) {
	if key == n.Current.Key() {
		return &n, true
	}
	return
}
