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
		t.Errorf("Expected PC to be %d, got %d", expectedPc, cpu.Pc)
	}
}
