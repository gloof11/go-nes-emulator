package main

type Olc6502 struct {
  bus *Bus
  FLAGS6502 map[string] uint8
  a uint8 
  x uint8
  y uint8
  stkp uint8
  pc uint16
  status uint8
  fetched uint8
  addr_abs uint16
  addr_rel uint16
  opcode uint8
  cycles uint8
  lookup []Instruction
}

type Instruction struct {
  name string
  operate func() uint8
  addrmode func() uint8
  cycles uint8
}

func NewOlc6502(b *Bus) *Olc6502 {
  o := new(Olc6502)
  o.bus = b

  o.FLAGS6502 = map[string] uint8{ 
    "C": 1 << 0,
    "Z": 1 << 1,
    "I": 1 << 2,
    "D": 1 << 3,
    "B": 1 << 4,
    "U": 1 << 5,
    "V": 1 << 6,
    "N": 1 << 7,
  }
  
  o.a = 0x00
  o.x = 0x00
  o.y = 0x00
  o.stkp = 0x00
  o.pc = 0x0000
  o.status = 0x00
  o.fetched = 0x00
  o.addr_abs = 0x000
  o.addr_rel = 0x00
  o.opcode = 0x00
  o.cycles = 0

  o.lookup = []Instruction{
    Instruction{"BRK",o.BRK,o.IMM,7},Instruction{"ORA",o.ORA,o.IZX,6},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,3},Instruction{"ORA",o.ORA,o.ZP0,3},Instruction{"ASL",o.ASL,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"PHP",o.PHP,o.IMP,3},Instruction{"ORA",o.ORA,o.IMM,2},Instruction{"ASL",o.ASL,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.NOP,o.IMP,4},Instruction{"ORA",o.ORA,o.ABS,4},Instruction{"ASL",o.ASL,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BPL",o.BPL,o.REL,2},Instruction{"ORA",o.ORA,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"ORA",o.ORA,o.ZPX,4},Instruction{"ASL",o.ASL,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"CLC",o.CLC,o.IMP,2},Instruction{"ORA",o.ORA,o.ABY,4},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"ORA",o.ORA,o.ABX,4},Instruction{"ASL",o.ASL,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
		Instruction{"JSR",o.JSR,o.ABS,6},Instruction{"AND",o.AND,o.IZX,6},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"BIT",o.BIT,o.ZP0,3},Instruction{"AND",o.AND,o.ZP0,3},Instruction{"ROL",o.ROL,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"PLP",o.PLP,o.IMP,4},Instruction{"AND",o.AND,o.IMM,2},Instruction{"ROL",o.ROL,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"BIT",o.BIT,o.ABS,4},Instruction{"AND",o.AND,o.ABS,4},Instruction{"ROL",o.ROL,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BMI",o.BMI,o.REL,2},Instruction{"AND",o.AND,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"AND",o.AND,o.ZPX,4},Instruction{"ROL",o.ROL,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"SEC",o.SEC,o.IMP,2},Instruction{"AND",o.AND,o.ABY,4},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"AND",o.AND,o.ABX,4},Instruction{"ROL",o.ROL,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
		Instruction{"RTI",o.RTI,o.IMP,6},Instruction{"EOR",o.EOR,o.IZX,6},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,3},Instruction{"EOR",o.EOR,o.ZP0,3},Instruction{"LSR",o.LSR,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"PHA",o.PHA,o.IMP,3},Instruction{"EOR",o.EOR,o.IMM,2},Instruction{"LSR",o.LSR,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"JMP",o.JMP,o.ABS,3},Instruction{"EOR",o.EOR,o.ABS,4},Instruction{"LSR",o.LSR,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BVC",o.BVC,o.REL,2},Instruction{"EOR",o.EOR,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"EOR",o.EOR,o.ZPX,4},Instruction{"LSR",o.LSR,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"CLI",o.CLI,o.IMP,2},Instruction{"EOR",o.EOR,o.ABY,4},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"EOR",o.EOR,o.ABX,4},Instruction{"LSR",o.LSR,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
		Instruction{"RTS",o.RTS,o.IMP,6},Instruction{"ADC",o.ADC,o.IZX,6},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,3},Instruction{"ADC",o.ADC,o.ZP0,3},Instruction{"ROR",o.ROR,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"PLA",o.PLA,o.IMP,4},Instruction{"ADC",o.ADC,o.IMM,2},Instruction{"ROR",o.ROR,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"JMP",o.JMP,o.IND,5},Instruction{"ADC",o.ADC,o.ABS,4},Instruction{"ROR",o.ROR,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BVS",o.BVS,o.REL,2},Instruction{"ADC",o.ADC,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"ADC",o.ADC,o.ZPX,4},Instruction{"ROR",o.ROR,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"SEI",o.SEI,o.IMP,2},Instruction{"ADC",o.ADC,o.ABY,4},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"ADC",o.ADC,o.ABX,4},Instruction{"ROR",o.ROR,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
		Instruction{"???",o.NOP,o.IMP,2},Instruction{"STA",o.STA,o.IZX,6},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,6},Instruction{"STY",o.STY,o.ZP0,3},Instruction{"STA",o.STA,o.ZP0,3},Instruction{"STX",o.STX,o.ZP0,3},Instruction{"???",o.XXX,o.IMP,3},Instruction{"DEY",o.DEY,o.IMP,2},Instruction{"???",o.NOP,o.IMP,2},Instruction{"TXA",o.TXA,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"STY",o.STY,o.ABS,4},Instruction{"STA",o.STA,o.ABS,4},Instruction{"STX",o.STX,o.ABS,4},Instruction{"???",o.XXX,o.IMP,4},
		Instruction{"BCC",o.BCC,o.REL,2},Instruction{"STA",o.STA,o.IZY,6},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,6},Instruction{"STY",o.STY,o.ZPX,4},Instruction{"STA",o.STA,o.ZPX,4},Instruction{"STX",o.STX,o.ZPY,4},Instruction{"???",o.XXX,o.IMP,4},Instruction{"TYA",o.TYA,o.IMP,2},Instruction{"STA",o.STA,o.ABY,5},Instruction{"TXS",o.TXS,o.IMP,2},Instruction{"???",o.XXX,o.IMP,5},Instruction{"???",o.NOP,o.IMP,5},Instruction{"STA",o.STA,o.ABX,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"???",o.XXX,o.IMP,5},
		Instruction{"LDY",o.LDY,o.IMM,2},Instruction{"LDA",o.LDA,o.IZX,6},Instruction{"LDX",o.LDX,o.IMM,2},Instruction{"???",o.XXX,o.IMP,6},Instruction{"LDY",o.LDY,o.ZP0,3},Instruction{"LDA",o.LDA,o.ZP0,3},Instruction{"LDX",o.LDX,o.ZP0,3},Instruction{"???",o.XXX,o.IMP,3},Instruction{"TAY",o.TAY,o.IMP,2},Instruction{"LDA",o.LDA,o.IMM,2},Instruction{"TAX",o.TAX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"LDY",o.LDY,o.ABS,4},Instruction{"LDA",o.LDA,o.ABS,4},Instruction{"LDX",o.LDX,o.ABS,4},Instruction{"???",o.XXX,o.IMP,4},
		Instruction{"BCS",o.BCS,o.REL,2},Instruction{"LDA",o.LDA,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,5},Instruction{"LDY",o.LDY,o.ZPX,4},Instruction{"LDA",o.LDA,o.ZPX,4},Instruction{"LDX",o.LDX,o.ZPY,4},Instruction{"???",o.XXX,o.IMP,4},Instruction{"CLV",o.CLV,o.IMP,2},Instruction{"LDA",o.LDA,o.ABY,4},Instruction{"TSX",o.TSX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,4},Instruction{"LDY",o.LDY,o.ABX,4},Instruction{"LDA",o.LDA,o.ABX,4},Instruction{"LDX",o.LDX,o.ABY,4},Instruction{"???",o.XXX,o.IMP,4},
		Instruction{"CPY",o.CPY,o.IMM,2},Instruction{"CMP",o.CMP,o.IZX,6},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"CPY",o.CPY,o.ZP0,3},Instruction{"CMP",o.CMP,o.ZP0,3},Instruction{"DEC",o.DEC,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"INY",o.INY,o.IMP,2},Instruction{"CMP",o.CMP,o.IMM,2},Instruction{"DEX",o.DEX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,2},Instruction{"CPY",o.CPY,o.ABS,4},Instruction{"CMP",o.CMP,o.ABS,4},Instruction{"DEC",o.DEC,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BNE",o.BNE,o.REL,2},Instruction{"CMP",o.CMP,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"CMP",o.CMP,o.ZPX,4},Instruction{"DEC",o.DEC,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"CLD",o.CLD,o.IMP,2},Instruction{"CMP",o.CMP,o.ABY,4},Instruction{"NOP",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"CMP",o.CMP,o.ABX,4},Instruction{"DEC",o.DEC,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
		Instruction{"CPX",o.CPX,o.IMM,2},Instruction{"SBC",o.SBC,o.IZX,6},Instruction{"???",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"CPX",o.CPX,o.ZP0,3},Instruction{"SBC",o.SBC,o.ZP0,3},Instruction{"INC",o.INC,o.ZP0,5},Instruction{"???",o.XXX,o.IMP,5},Instruction{"INX",o.INX,o.IMP,2},Instruction{"SBC",o.SBC,o.IMM,2},Instruction{"NOP",o.NOP,o.IMP,2},Instruction{"???",o.SBC,o.IMP,2},Instruction{"CPX",o.CPX,o.ABS,4},Instruction{"SBC",o.SBC,o.ABS,4},Instruction{"INC",o.INC,o.ABS,6},Instruction{"???",o.XXX,o.IMP,6},
		Instruction{"BEQ",o.BEQ,o.REL,2},Instruction{"SBC",o.SBC,o.IZY,5},Instruction{"???",o.XXX,o.IMP,2},Instruction{"???",o.XXX,o.IMP,8},Instruction{"???",o.NOP,o.IMP,4},Instruction{"SBC",o.SBC,o.ZPX,4},Instruction{"INC",o.INC,o.ZPX,6},Instruction{"???",o.XXX,o.IMP,6},Instruction{"SED",o.SED,o.IMP,2},Instruction{"SBC",o.SBC,o.ABY,4},Instruction{"NOP",o.NOP,o.IMP,2},Instruction{"???",o.XXX,o.IMP,7},Instruction{"???",o.NOP,o.IMP,4},Instruction{"SBC",o.SBC,o.ABX,4},Instruction{"INC",o.INC,o.ABX,7},Instruction{"???",o.XXX,o.IMP,7},
  }

  return o
}

func (o *Olc6502) read(a uint16) uint8 {
  return o.bus.read(a, false)
}

func (o *Olc6502) write(a uint16, d uint8) {
  o.bus.write(a, d)
}

func (o *Olc6502) GetFlag(f string) uint8 {}
func (o *Olc6502) SetFlag(f string, v bool) {}

// Addressing Modes
func (o *Olc6502) IMP() uint8 {}
func (o *Olc6502) IMM() uint8 {}
func (o *Olc6502) ZP0() uint8 {}
func (o *Olc6502) ZPX() uint8 {}
func (o *Olc6502) ZPY() uint8 {}
func (o *Olc6502) REL() uint8 {}
func (o *Olc6502) ABS() uint8 {}
func (o *Olc6502) ABX() uint8 {}
func (o *Olc6502) ABY() uint8 {}
func (o *Olc6502) IND() uint8 {}
func (o *Olc6502) IZX() uint8 {}
func (o *Olc6502) IZY() uint8 {}

// Opcodes
func (o *Olc6502) ADC() uint8 {}
func (o *Olc6502) AND() uint8 {}
func (o *Olc6502) ASL() uint8 {}
func (o *Olc6502) BCC() uint8 {}
func (o *Olc6502) BCS() uint8 {}
func (o *Olc6502) BEQ() uint8 {}
func (o *Olc6502) BIT() uint8 {}
func (o *Olc6502) BMI() uint8 {}
func (o *Olc6502) BNE() uint8 {}
func (o *Olc6502) BPL() uint8 {}
func (o *Olc6502) BRK() uint8 {}
func (o *Olc6502) BVC() uint8 {}
func (o *Olc6502) BVS() uint8 {}
func (o *Olc6502) CLC() uint8 {}
func (o *Olc6502) CLD() uint8 {}
func (o *Olc6502) CLI() uint8 {}
func (o *Olc6502) CLV() uint8 {}
func (o *Olc6502) CMP() uint8 {}
func (o *Olc6502) CPX() uint8 {}
func (o *Olc6502) CPY() uint8 {}
func (o *Olc6502) DEC() uint8 {}
func (o *Olc6502) DEX() uint8 {}
func (o *Olc6502) DEY() uint8 {}
func (o *Olc6502) EOR() uint8 {}
func (o *Olc6502) INC() uint8 {}
func (o *Olc6502) INX() uint8 {}
func (o *Olc6502) INY() uint8 {}
func (o *Olc6502) JMP() uint8 {}
func (o *Olc6502) JSR() uint8 {}
func (o *Olc6502) LDA() uint8 {}
func (o *Olc6502) LDX() uint8 {}
func (o *Olc6502) LDY() uint8 {}
func (o *Olc6502) LSR() uint8 {}
func (o *Olc6502) NOP() uint8 {}
func (o *Olc6502) ORA() uint8 {}
func (o *Olc6502) PHA() uint8 {}
func (o *Olc6502) PHP() uint8 {}
func (o *Olc6502) PLA() uint8 {}
func (o *Olc6502) PLP() uint8 {}
func (o *Olc6502) ROL() uint8 {}
func (o *Olc6502) ROR() uint8 {}
func (o *Olc6502) RTI() uint8 {}
func (o *Olc6502) RTS() uint8 {}
func (o *Olc6502) SBC() uint8 {}
func (o *Olc6502) SEC() uint8 {}
func (o *Olc6502) SED() uint8 {}
func (o *Olc6502) SEI() uint8 {}
func (o *Olc6502) STA() uint8 {}
func (o *Olc6502) STX() uint8 {}
func (o *Olc6502) STY() uint8 {}
func (o *Olc6502) TAX() uint8 {}
func (o *Olc6502) TAY() uint8 {}
func (o *Olc6502) TSX() uint8 {}
func (o *Olc6502) TXA() uint8 {}
func (o *Olc6502) TXS() uint8 {}
func (o *Olc6502) TYA() uint8 {}

func (o *Olc6502) XXX() uint8 {}
// Necessary functions
func (o *Olc6502) clock() {
  if (o.cycles == 0) {
    o.opcode = o.read(o.pc)
    o.pc++
  }
}
func (o *Olc6502) reset() {}
func (o *Olc6502) irq() {}
func (o *Olc6502) nmi() {}

// Fetching data
func (o *Olc6502) fetch() uint8 {}
