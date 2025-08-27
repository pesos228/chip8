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
		{Name: "RET", Mask: 0xFFFF, Pattern: 0x00EE, Handler: handleRet},
		{Name: "ADD Vx, byte", Mask: 0xF000, Pattern: 0x7000, Handler: handleAddVxByte},
		{Name: "JP addr", Mask: 0xF000, Pattern: 0x1000, Handler: handleJumpAddr},
		{Name: "SYS addr", Mask: 0xF000, Pattern: 0x0000, Handler: handleSysAddr},
	}
}

func handleAddVxByte(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	kk := uint8(opcode & 0x00FF)
	c.Registers[x] += kk
	c.Pc += 2
}

func handleJumpAddr(c *Cpu, opcode uint16) {
	addr := opcode & 0x0FFF
	c.Pc = addr
}

func handleSysAddr(c *Cpu, opcode uint16) {
	c.Pc += 2
}

func handleRet(c *Cpu, opcode uint16) {
	if c.Sp == 0 {
		return
	}
	c.Sp--
	c.Pc = c.Stack[c.Sp]
	c.Pc += 2
}
