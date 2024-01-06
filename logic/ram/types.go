package ram

import "fmt"

type TypeRAM int

const (
	REDIS TypeRAM = iota
)

type Top []TopAddr

type TopAddr struct {
	Addr  string
	Count int
}

func (t TopAddr) String() string {
	return fmt.Sprintf("Addr: %s\n", t.Addr) +
		fmt.Sprintf("Count: %d\n", t.Count)
}

type ListERC20 struct {
	Addr map[string]bool
}
