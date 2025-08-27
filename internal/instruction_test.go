package cpu

import (
	"testing"
)

func TestInstruction_7XNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	cpu.Registers[3] = 5

	opcode := uint16(0x7304)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedRegisterValue := uint8(9)
	if cpu.Registers[3] != expectedRegisterValue {
		t.Errorf("Expected V3 to be %d, got %d", expectedRegisterValue, cpu.Registers[3])
	}

	expectedPc := uint16(0x102)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}
}

func TestInstruction_1NNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x1234)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedPc := uint16(0x234)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}
}

func TestInstruction_00EE(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	cpu.Stack[0] = 0x356
	cpu.Sp = 1

	opcode := uint16(0x00EE)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedPc := uint16(0x356 + 2)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}

	expectedSp := uint8(0)
	if cpu.Sp != expectedSp {
		t.Errorf("Expected SP to be %d, got %d", expectedSp, cpu.Sp)
	}
}

func TestInstruction_0NNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x0123)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedPc := uint16(0x102)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}
}
