package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	users := []user{
		{
			name: "小王",
			age:  19,
		},
		{
			name: "小红",
			age:  18,
		},
		{
			name: "小红1",
			age:  16,
		},
		{
			name: "小红2",
			age:  13,
		},
		{
			name: "小红3",
			age:  33,
		},
	}
	hp := NewHeap(users, func(a, b user) bool {
		return a.age < b.age
	})

	hp.Push(user{
		name: "p1",
		age:  32,
	})
	hp.Push(user{
		name: "p2",
		age:  31,
	})
	hp.Pop()
	hp.Push(user{
		name: "p3",
		age:  13,
	})
	hp.Push(user{
		name: "p4",
		age:  18,
	})
	for !hp.IsEmpty() {
		t.Log("\n GetTOP: ", hp.GetTop(), "\n")
		hp.Pop()
	}

}
