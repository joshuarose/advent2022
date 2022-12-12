package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

type Monkey struct {
	items           Stack
	divisibleBy     int
	trueMonkey      *Monkey
	falseMonkey     *Monkey
	operation       func(f int) int
	inspectionCount int
}

func (m *Monkey) inspect() {
	// worry level modified during inspection
	itemsLen := len(m.items.items)
	for i := itemsLen - 1; i > 0; i-- {
		worryLevel := m.items.items[i]
		worryLevel = m.operation(worryLevel)
		m.inspectionCount += 1
		worryLevel = worryLevel / 3
		passMonkey := m.falseMonkey
		if worryLevel%m.divisibleBy == 0 {
			passMonkey = m.trueMonkey
		}
		passedItem := m.items.pop()
		passMonkey.items.push(passedItem)
	}
}

var monkeys []*Monkey

func init() {
	monkeys = make([]*Monkey, 8)
	monkeys[0] = &Monkey{
		items: Stack{
			items: []int{97, 81, 57, 57, 91, 61},
		},
		operation:   func(f int) int { return f * 7 },
		divisibleBy: 11,
	}

	monkeys[1] = &Monkey{
		items: Stack{
			items: []int{88, 62, 68, 90},
		},
		operation:   func(f int) int { return f * 17 },
		divisibleBy: 19,
	}

	monkeys[2] = &Monkey{
		items:       Stack{items: []int{74, 87}},
		operation:   func(f int) int { return f + 2 },
		divisibleBy: 5,
	}

	monkeys[3] = &Monkey{
		items:       Stack{items: []int{53, 81, 60, 87, 90, 99, 75}},
		operation:   func(f int) int { return f + 1 },
		divisibleBy: 2,
	}

	monkeys[4] = &Monkey{
		items:       Stack{items: []int{57}},
		operation:   func(f int) int { return f + 6 },
		divisibleBy: 13,
	}

	monkeys[5] = &Monkey{
		items:       Stack{items: []int{54, 84, 91, 55, 59, 72, 75, 70}},
		operation:   func(f int) int { return f * f },
		divisibleBy: 7,
	}

	monkeys[6] = &Monkey{
		items:       Stack{items: []int{95, 79, 79, 68, 78}},
		operation:   func(f int) int { return f + 3 },
		divisibleBy: 3,
	}

	monkeys[7] = &Monkey{
		items:       Stack{items: []int{61, 97, 67}},
		operation:   func(f int) int { return f + 4 },
		divisibleBy: 17,
	}
	// Monkey 0
	// trueMonkey:  monkeys[5],
	// falseMonkey: monkeys[6],
	linkMonkeys(monkeys[0], monkeys[5], monkeys[6])
	// Monkey 1
	// trueMonkey:  monkeys[4],
	// falseMonkey: monkeys[2],
	linkMonkeys(monkeys[1], monkeys[4], monkeys[2])
	// Monkey 2
	// trueMonkey:  monkeys[7],
	// falseMonkey: monkeys[4],
	linkMonkeys(monkeys[2], monkeys[7], monkeys[4])
	// Monkey 3
	// trueMonkey:  monkeys[2],
	// falseMonkey: monkeys[1],
	linkMonkeys(monkeys[3], monkeys[2], monkeys[1])
	// Monkey 4
	// trueMonkey:  monkeys[7],
	// falseMonkey: monkeys[0],
	linkMonkeys(monkeys[4], monkeys[7], monkeys[0])
	// Monkey 5
	// trueMonkey:  monkeys[6],
	// falseMonkey: monkeys[3],
	linkMonkeys(monkeys[5], monkeys[6], monkeys[3])
	// Monkey 6
	// trueMonkey:  monkeys[1],
	// falseMonkey: monkeys[3],
	linkMonkeys(monkeys[6], monkeys[1], monkeys[3])
	// Monkey 7
	// trueMonkey:  monkeys[0],
	// falseMonkey: monkeys[5],
	linkMonkeys(monkeys[7], monkeys[0], monkeys[5])
}

func main() {
	// 35308 too low
	// 33293 too low
	for i := 0; i < 20; i++ {
		for j := 0; j < 8; j++ {
			monkeys[j].inspect()
		}
		fmt.Printf("Round: %d completed", i+1)
	}

	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, m.inspectionCount)
	}
}

func linkMonkeys(activeMonkey, trueMonkey, falseMonkey *Monkey) {
	activeMonkey.trueMonkey = trueMonkey
	activeMonkey.falseMonkey = falseMonkey
}
