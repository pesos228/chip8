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

	expectedPc := uint16(0x356)
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

func TestInstruction_2NNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x2345)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedPc := uint16(0x345)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}

	expectedSp := uint8(1)
	if cpu.Sp != expectedSp {
		t.Errorf("Expected SP to be %d, got %d", expectedSp, cpu.Sp)
	}

	expectedStackValue := uint16(0x102)
	if cpu.Stack[0] != expectedStackValue {
		t.Errorf("Expected Stack[0] to be 0x%X, got 0x%X", expectedStackValue, cpu.Stack[0])
	}
}

func TestInstruction_3XNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x3216)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[2] = 0x16

	cpu.Execute()

	expectedPc := uint16(0x104)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}

}

func TestInstruction_4XNN(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x4216)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[2] = 0x12

	cpu.Execute()

	expectedPc := uint16(0x104)
	if cpu.Pc != expectedPc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}

}

func TestInstruction_5XY0(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x5350)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[3] = 0x12
	cpu.Registers[5] = 0x12

	cpu.Execute()

	expectedPc := uint16(0x104)
	if expectedPc != cpu.Pc {
		t.Errorf("Expected PC to be 0x%X, got 0x%X", expectedPc, cpu.Pc)
	}
}

func TestInstruction_6XKK(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x6412)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Execute()

	expectedRegisterValue := uint8(0x12)
	if expectedRegisterValue != cpu.Registers[4] {
		t.Errorf("Expected Register[4] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[4])
	}
}

func TestInstruction_8XY0(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8340)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x3

	cpu.Execute()

	expectedRegisterValue := uint8(0x3)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}
}

func TestInstruction_8XY1(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8341)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x3
	cpu.Registers[3] = 0x10

	cpu.Execute()

	expectedRegisterValue := uint8(0x13)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}
}

func TestInstruction_8XY2(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8342)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x3
	cpu.Registers[3] = 0x13

	cpu.Execute()

	expectedRegisterValue := uint8(0x3)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}
}

func TestInstruction_8XY3(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8343)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x3
	cpu.Registers[3] = 0x13

	cpu.Execute()

	expectedRegisterValue := uint8(0x10)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}
}

func TestInstruction_8XY4_no_carry(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8344)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x3
	cpu.Registers[3] = 0x13

	cpu.Execute()

	expectedRegisterValue := uint8(0x16)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x0)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY4_with_carry(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8344)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x03
	cpu.Registers[3] = 0xFF

	cpu.Execute()

	expectedRegisterValue := uint8(0x02)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x1)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY5_no_borrow(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8345)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x0A
	cpu.Registers[3] = 0x14

	cpu.Execute()

	expectedRegisterValue := uint8(0x0A)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x1)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY5_with_borrow(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8345)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x14
	cpu.Registers[3] = 0x0A

	cpu.Execute()

	expectedRegisterValue := uint8(0xF6)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x0)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY6_LSB_is_zero(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8346)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[3] = 0x0A

	cpu.Execute()

	expectedRegisterValue := uint8(0x05)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x0)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY6_LSB_is_one(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8346)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[3] = 0x0B

	cpu.Execute()

	expectedRegisterValue := uint8(0x05)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x1)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY7_no_borrow(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8347)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x0A
	cpu.Registers[3] = 0x05

	cpu.Execute()

	expectedRegisterValue := uint8(0x05)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x1)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}

func TestInstruction_8XY7_with_borrow(t *testing.T) {
	cpu := NewCpu(512, 0x100)

	opcode := uint16(0x8347)

	cpu.Memory[0x100] = uint8(opcode >> 8)
	cpu.Memory[0x101] = uint8(opcode & 0x00FF)

	cpu.Registers[4] = 0x05
	cpu.Registers[3] = 0x0A

	cpu.Execute()

	expectedRegisterValue := uint8(0xFB)
	if expectedRegisterValue != cpu.Registers[3] {
		t.Errorf("Expected Register[3] to be 0x%X, got 0x%X", expectedRegisterValue, cpu.Registers[3])
	}

	expectedVFValue := uint8(0x0)
	if expectedVFValue != cpu.Registers[15] {
		t.Errorf("Expected Register[15] to be 0x%X, got 0x%X", expectedVFValue, cpu.Registers[15])
	}
}
