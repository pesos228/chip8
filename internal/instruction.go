package cpu

type Instruction struct {
	Name    string
	Mask    uint16
	Pattern uint16
	Handler func(c *Cpu, opcode uint16)
}

var instructions []Instruction

func init() {
	instructions = []Instruction{
		{Name: "ADD Vx, byte", Mask: 0xF000, Pattern: 0x7000, Handler: handleAddVxByte},
	}
}

func handleAddVxByte(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	kk := uint8(opcode & 0x00FF)
	c.Registers[x] += kk
	c.Pc += 2
}
