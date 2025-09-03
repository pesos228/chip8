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
		{Name: "RET (00EE)", Mask: 0xFFFF, Pattern: 0x00EE, Handler: handleRet},
		{Name: "LD Vx, Vy (8xy0)", Mask: 0xF00F, Pattern: 0x8000, Handler: handleStoreValFromReg},
		{Name: "OR Vx, Vy (8xy1)", Mask: 0xF00F, Pattern: 0x8001, Handler: handleBitwiseOr},
		{Name: "AND Vx, Vy (8xy2)", Mask: 0xF00F, Pattern: 0x8002, Handler: handleBitwiseAnd},
		{Name: "XOR Vx, Vy (8xy3)", Mask: 0xF00F, Pattern: 0x8003, Handler: handleBitwiseXor},
		{Name: "SYS addr (0nnn)", Mask: 0xF000, Pattern: 0x0000, Handler: handleSysAddr},
		{Name: "JP addr (1nnn)", Mask: 0xF000, Pattern: 0x1000, Handler: handleJumpAddr},
		{Name: "CALL addr (2nnn)", Mask: 0xF000, Pattern: 0x2000, Handler: handleCallAddr},
		{Name: "SE Vx, byte (3xkk)", Mask: 0xF000, Pattern: 0x3000, Handler: handleSkipIfEqual},
		{Name: "SNE Vx, byte (4xkk)", Mask: 0xF000, Pattern: 0x4000, Handler: handleSkipIfNotEqual},
		{Name: "SE Vx, Vy (5xy0)", Mask: 0xF000, Pattern: 0x5000, Handler: handleSkipIfRegEqual},
		{Name: "LD Vx, byte (6xkk)", Mask: 0xF000, Pattern: 0x6000, Handler: handlePutValueInReg},
		{Name: "ADD Vx, byte (7xkk)", Mask: 0xF000, Pattern: 0x7000, Handler: handleAddVxByte},
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
}

func handleCallAddr(c *Cpu, opcode uint16) {
	addr := opcode & 0x0FFF
	c.Stack[c.Sp] = c.Pc + 2
	c.Sp++
	c.Pc = addr
}

func handleSkipIfEqual(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	kk := uint8(opcode & 0x00FF)
	if c.Registers[x] == kk {
		c.Pc += 4
	} else {
		c.Pc += 2
	}
}

func handleSkipIfNotEqual(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	kk := uint8(opcode & 0x00FF)
	if c.Registers[x] != kk {
		c.Pc += 4
	} else {
		c.Pc += 2
	}
}

func handleSkipIfRegEqual(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	if c.Registers[x] == c.Registers[y] {
		c.Pc += 4
	} else {
		c.Pc += 2
	}
}

func handlePutValueInReg(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	kk := uint8(opcode & 0x00FF)

	c.Registers[x] = kk
	c.Pc += 2
}

func handleStoreValFromReg(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.Registers[x] = c.Registers[y]

	c.Pc += 2
}

func handleBitwiseOr(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.Registers[x] = c.Registers[x] | c.Registers[y]

	c.Pc += 2
}

func handleBitwiseAnd(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.Registers[x] = c.Registers[x] & c.Registers[y]
	c.Pc += 2
}

func handleBitwiseXor(c *Cpu, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.Registers[x] = c.Registers[x] ^ c.Registers[y]
	c.Pc += 2
}
