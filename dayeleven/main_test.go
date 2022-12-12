package main

import "testing"

func TestMain(t *testing.T) {
	if 1+1 != 2 {
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	testStack := Stack{}
	testStack.push(2)
	testStack.push(3)
	testStack.push(4)

	if len(testStack.items) != 3 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	testStack := Stack{}
	testStack.push(2)
	testStack.push(3)
	testStack.push(4)

	result := testStack.pop()

	if len(testStack.items) != 2 {
		t.Fail()
	}
	if result != 4 {
		t.Fail()
	}
}

func TestExample(t *testing.T) {
	monkeys = make([]*Monkey, 4)
	monkeys = []*Monkey{
		{
			items:       Stack{items: []int{79, 98}},
			operation:   func(f int) int { return f * 19 },
			divisibleBy: 23,
		},
		{
			items:       Stack{items: []int{54, 65, 75, 74}},
			operation:   func(f int) int { return f + 6 },
			divisibleBy: 19,
		},
		{
			items:       Stack{items: []int{79, 60, 97}},
			operation:   func(f int) int { return f * f },
			divisibleBy: 13,
		},
		{
			items:       Stack{items: []int{74}},
			operation:   func(f int) int { return f + 3 },
			divisibleBy: 17,
		},
	}

	linkMonkeys(monkeys[0], monkeys[2], monkeys[3])
	linkMonkeys(monkeys[1], monkeys[2], monkeys[0])
	linkMonkeys(monkeys[2], monkeys[1], monkeys[3])
	linkMonkeys(monkeys[3], monkeys[0], monkeys[1])
	main()
	expected := 101
	if monkeys[0].inspectionCount != expected {
		t.Errorf("expected %d but received %d", expected, monkeys[0].inspectionCount)
	}
}
