package computer

import (
	"fmt"
	"strings"
)

// This will set up the structure of the computer and the methods that will be used to interact with it.

type Computer struct {
	A, B, C int   // registers
	Program []int // program instructions
	instPtr int   // index of program execution
	Output  []int // output buffer
}

func (c Computer) Print() {
	fmt.Printf("Register A: %d\n", c.A)
	fmt.Printf("Register B: %d\n", c.B)
	fmt.Printf("Register C: %d\n", c.C)
	fmt.Printf("\nProgram: ")
	for i := range c.Program {
		fmt.Printf("%d ", c.Program[i])
	}
	fmt.Printf("\n")
}

// Now the instructions. Each instruction will be in charge of increasing the instruction
// pointer. This will skip the instruction and the operand.

func (c *Computer) Adv() {
	// Opcode 0
	// division, numerator is register A, denominator is 2^(whatever the combo operand is)
	c.instPtr++

	operand := c.Program[c.instPtr]

	switch operand {
	case 0, 1, 2, 3:
		c.A = c.A >> operand
	case 4:
		c.A = c.A >> c.A
	case 5:
		c.A = c.A >> c.B
	case 6:
		c.A = c.A >> c.C
	}
	// then move the pointer
	c.instPtr++
}

func (c *Computer) Bxl() {
	// opcode 1
	// bitwise XOR of register B and instruction's literal operand.
	// stores in register B
	c.instPtr++
	operand := c.Program[c.instPtr]
	c.B = c.B ^ operand
	c.instPtr++
}

func (c *Computer) Bst() {
	// opcode 2
	// calculates the value of its combo operand modulo 8,
	// keeping lowest 3 bits, write to register B
	c.instPtr++
	operand := c.Program[c.instPtr]

	switch operand {
	case 0, 1, 2, 3:
		c.B = operand % 8
	case 4:
		c.B = c.A % 8
	case 5:
		c.B = c.B % 8
	case 6:
		c.B = c.C % 8
	}

	c.instPtr++
}

func (c *Computer) Jnz() {
	// opcode 3
	// jump not zero does nothing if A = 0, but if not zero it jumps
	// by setting instruction pointer to the value of its literal operand.
	// If this instruction jumps, instruction pointer is NOT increased by
	// 2 after this instruction
	if c.A != 0 {
		c.instPtr = c.Program[c.instPtr+1]
	} else {
		c.instPtr += 2
	}
}

func (c *Computer) Bxc() {
	// opcode 4
	// calculates bitwise XOR of register B and register C, stores in register B
	// reads an operand but ignores it
	c.instPtr++
	c.B = c.B ^ c.C
	c.instPtr++
}

func (c *Computer) Out() {
	// opcode 5
	// calculates the value of its combo operand modulo 8, outputs the value
	c.instPtr++
	operand := c.Program[c.instPtr]
	switch operand {
	case 0, 1, 2, 3:
		c.Output = append(c.Output, operand%8)
	case 4:
		c.Output = append(c.Output, c.A%8)
	case 5:
		c.Output = append(c.Output, c.B%8)
	case 6:
		c.Output = append(c.Output, c.C%8)
	}
	c.instPtr++
}

func (c *Computer) OutAndConfirm() bool {
	// opcode 5
	// This will perform the same operation as the Out() function, but it will
	// also confirm that this number is the same as the program output at each
	// step. If not, it will return false so we can bomb out in trying this
	// combination.
	idx := len(c.Output)

	c.instPtr++
	operand := c.Program[c.instPtr]
	var val int
	switch operand {
	case 0, 1, 2, 3:
		val = operand % 8
	case 4:
		val = c.A % 8
	case 5:
		val = c.B % 8
	case 6:
		val = c.C % 8
	}
	c.Output = append(c.Output, val)
	c.instPtr++
	return c.Program[idx] == val
}

func (c *Computer) Bdv() {
	// opcode 6
	// exactly like the ADV instruction, but stores the result in register B.
	// numerator is still read from register A
	c.instPtr++
	operand := c.Program[c.instPtr]
	switch operand {
	case 0, 1, 2, 3:
		c.B = c.A >> operand
	case 4:
		c.B = c.A >> c.A
	case 5:
		c.B = c.A >> c.B
	case 6:
		c.B = c.A >> c.C
	}
	c.instPtr++
}

func (c *Computer) Cdv() {
	// opcode 7
	// exactly like the ADV instruction, but stores the result in register C.
	// numerator is still read from register A
	c.instPtr++
	operand := c.Program[c.instPtr]
	switch operand {
	case 0, 1, 2, 3:
		c.C = c.A >> operand
	case 4:
		c.C = c.A >> c.A
	case 5:
		c.C = c.A >> c.B
	case 6:
		c.C = c.A >> c.C
	}
	c.instPtr++
}

func (c *Computer) PrintOutput() {
	for i := range c.Output {
		fmt.Printf("%d,", c.Output[i])
	}
	fmt.Printf("\n")
}

func (c *Computer) PrintState() {
	// concise state printing for debug purposes
	fmt.Printf("A: %d, B: %d, C: %d IP: %d Out: %s Prog: %s\n", c.A, c.B, c.C, c.instPtr, strings.Join(strings.Fields(fmt.Sprint(c.Output)), ","), strings.Join(strings.Fields(fmt.Sprint(c.Program)), ","))
}

func (c *Computer) Run() {
	// Executes the program in the order, if the instruction pointer is
	// out of bounds, the program will stop
	for {
		// fmt.Printf("Instruction pointer: %d\n", c.instPtr)
		// c.PrintState()
		if c.instPtr >= len(c.Program) {
			break
		}
		// get opcode
		opcode := c.Program[c.instPtr]
		switch opcode {
		case 0:
			c.Adv()
		case 1:
			c.Bxl()
		case 2:
			c.Bst()
		case 3:
			c.Jnz()
		case 4:
			c.Bxc()
		case 5:
			c.Out()
		case 6:
			c.Bdv()
		case 7:
			c.Cdv()
		default:
			panic("Invalid opcode")
		}
	}
}

func (c *Computer) RunAndConfirm() bool {
	// This will run the program, but it will also confirm that the output
	// is the same as the program output at each step. If not, it will return
	// false so we can bomb out in trying this combination.
	for {
		if c.instPtr >= len(c.Program) {
			break
		}
		// get opcode
		opcode := c.Program[c.instPtr]
		switch opcode {
		case 0:
			c.Adv()
		case 1:
			c.Bxl()
		case 2:
			c.Bst()
		case 3:
			c.Jnz()
		case 4:
			c.Bxc()
		case 5:
			if !c.OutAndConfirm() {
				return false
			}
		case 6:
			c.Bdv()
		case 7:
			c.Cdv()
		default:
			panic("Invalid opcode")
		}
	}
	return true
}

func (c *Computer) Initialize(aVal int) *Computer {
	return &Computer{
		A:       aVal,
		B:       0,
		C:       0,
		Program: c.Program,
		instPtr: 0,
	}
}
