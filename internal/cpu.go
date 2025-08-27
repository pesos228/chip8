package cpu

import (
	"fmt"
)

type Cpu struct {
	Memory    []uint8
	Registers [16]uint8
	Stack     [16]uint16
	Sp        uint8
	Pc        uint16
	I         uint16
	Dt        uint8
	St        uint8
	Config    Config
}

type Config struct {
	MemorySize   uint16
	ProgramStart uint16
}

func NewCpu(memorySize, programStart uint16) *Cpu {
	cpu := &Cpu{
		Pc:     programStart,
		Memory: make([]uint8, memorySize),
		Config: Config{
			MemorySize:   memorySize,
			ProgramStart: programStart,
		},
	}
	return cpu
}

func (c *Cpu) LoadGame(game []uint8) error {
	c.Reset()

	availableMemory := c.Config.MemorySize - c.Config.ProgramStart
	if len(game) > int(availableMemory) {
		return fmt.Errorf("game size (%d bytes) exceeds available memory (%d bytes)", len(game), availableMemory)
	}

	programStart := int(c.Config.ProgramStart)

	for i := 0; i < len(game); i++ {
		c.Memory[programStart+i] = game[i]
	}

	return nil
}

func (c *Cpu) Reset() {
	c.Memory = make([]uint8, c.Config.MemorySize)
	c.Registers = [16]uint8{}
	c.Stack = [16]uint16{}
	c.Sp = 0
	c.Pc = c.Config.ProgramStart
	c.I = 0
	c.Dt = 0
	c.St = 0
}

func (c *Cpu) Execute() {
	opcode := uint16(c.Memory[c.Pc])<<8 | uint16(c.Memory[c.Pc+1])

	for _, instr := range instructions {
		if opcode&instr.Mask == instr.Pattern {
			instr.Handler(c, opcode)
			return
		}
	}

	fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	c.Pc += 2
}
