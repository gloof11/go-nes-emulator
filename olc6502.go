package main

import (
	"reflect"
)

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
  operate func(o *Olc6502) uint8
  addrmode func(o *Olc6502) uint8
  cycles uint8
}

func NewOlc6502(b *Bus) *Olc6502 {
  o := new(Olc6502)
  o.bus = b

  o.FLAGS6502 = map[string] uint8{ 
    "C": 1 << 0, // 1
    "Z": 1 << 1, // 2
    "I": 1 << 2, // 4
    "D": 1 << 3, // 8
    "B": 1 << 4, // 16
    "U": 1 << 5, // 32
    "V": 1 << 6, // 64
    "N": 1 << 7, // 128
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
    {"BRK",o.BRK,o.IMM,7},{"ORA",o.ORA,o.IZX,6},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,3},{"ORA",o.ORA,o.ZP0,3},{"ASL",o.ASL,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"PHP",o.PHP,o.IMP,3},{"ORA",o.ORA,o.IMM,2},{"ASL",o.ASL,o.IMP,2},{"???",o.XXX,o.IMP,2},{"???",o.NOP,o.IMP,4},{"ORA",o.ORA,o.ABS,4},{"ASL",o.ASL,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BPL",o.BPL,o.REL,2},{"ORA",o.ORA,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"ORA",o.ORA,o.ZPX,4},{"ASL",o.ASL,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"CLC",o.CLC,o.IMP,2},{"ORA",o.ORA,o.ABY,4},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"ORA",o.ORA,o.ABX,4},{"ASL",o.ASL,o.ABX,7},{"???",o.XXX,o.IMP,7},
		{"JSR",o.JSR,o.ABS,6},{"AND",o.AND,o.IZX,6},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"BIT",o.BIT,o.ZP0,3},{"AND",o.AND,o.ZP0,3},{"ROL",o.ROL,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"PLP",o.PLP,o.IMP,4},{"AND",o.AND,o.IMM,2},{"ROL",o.ROL,o.IMP,2},{"???",o.XXX,o.IMP,2},{"BIT",o.BIT,o.ABS,4},{"AND",o.AND,o.ABS,4},{"ROL",o.ROL,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BMI",o.BMI,o.REL,2},{"AND",o.AND,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"AND",o.AND,o.ZPX,4},{"ROL",o.ROL,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"SEC",o.SEC,o.IMP,2},{"AND",o.AND,o.ABY,4},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"AND",o.AND,o.ABX,4},{"ROL",o.ROL,o.ABX,7},{"???",o.XXX,o.IMP,7},
		{"RTI",o.RTI,o.IMP,6},{"EOR",o.EOR,o.IZX,6},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,3},{"EOR",o.EOR,o.ZP0,3},{"LSR",o.LSR,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"PHA",o.PHA,o.IMP,3},{"EOR",o.EOR,o.IMM,2},{"LSR",o.LSR,o.IMP,2},{"???",o.XXX,o.IMP,2},{"JMP",o.JMP,o.ABS,3},{"EOR",o.EOR,o.ABS,4},{"LSR",o.LSR,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BVC",o.BVC,o.REL,2},{"EOR",o.EOR,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"EOR",o.EOR,o.ZPX,4},{"LSR",o.LSR,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"CLI",o.CLI,o.IMP,2},{"EOR",o.EOR,o.ABY,4},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"EOR",o.EOR,o.ABX,4},{"LSR",o.LSR,o.ABX,7},{"???",o.XXX,o.IMP,7},
		{"RTS",o.RTS,o.IMP,6},{"ADC",o.ADC,o.IZX,6},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,3},{"ADC",o.ADC,o.ZP0,3},{"ROR",o.ROR,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"PLA",o.PLA,o.IMP,4},{"ADC",o.ADC,o.IMM,2},{"ROR",o.ROR,o.IMP,2},{"???",o.XXX,o.IMP,2},{"JMP",o.JMP,o.IND,5},{"ADC",o.ADC,o.ABS,4},{"ROR",o.ROR,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BVS",o.BVS,o.REL,2},{"ADC",o.ADC,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"ADC",o.ADC,o.ZPX,4},{"ROR",o.ROR,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"SEI",o.SEI,o.IMP,2},{"ADC",o.ADC,o.ABY,4},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"ADC",o.ADC,o.ABX,4},{"ROR",o.ROR,o.ABX,7},{"???",o.XXX,o.IMP,7},
		{"???",o.NOP,o.IMP,2},{"STA",o.STA,o.IZX,6},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,6},{"STY",o.STY,o.ZP0,3},{"STA",o.STA,o.ZP0,3},{"STX",o.STX,o.ZP0,3},{"???",o.XXX,o.IMP,3},{"DEY",o.DEY,o.IMP,2},{"???",o.NOP,o.IMP,2},{"TXA",o.TXA,o.IMP,2},{"???",o.XXX,o.IMP,2},{"STY",o.STY,o.ABS,4},{"STA",o.STA,o.ABS,4},{"STX",o.STX,o.ABS,4},{"???",o.XXX,o.IMP,4},
		{"BCC",o.BCC,o.REL,2},{"STA",o.STA,o.IZY,6},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,6},{"STY",o.STY,o.ZPX,4},{"STA",o.STA,o.ZPX,4},{"STX",o.STX,o.ZPY,4},{"???",o.XXX,o.IMP,4},{"TYA",o.TYA,o.IMP,2},{"STA",o.STA,o.ABY,5},{"TXS",o.TXS,o.IMP,2},{"???",o.XXX,o.IMP,5},{"???",o.NOP,o.IMP,5},{"STA",o.STA,o.ABX,5},{"???",o.XXX,o.IMP,5},{"???",o.XXX,o.IMP,5},
		{"LDY",o.LDY,o.IMM,2},{"LDA",o.LDA,o.IZX,6},{"LDX",o.LDX,o.IMM,2},{"???",o.XXX,o.IMP,6},{"LDY",o.LDY,o.ZP0,3},{"LDA",o.LDA,o.ZP0,3},{"LDX",o.LDX,o.ZP0,3},{"???",o.XXX,o.IMP,3},{"TAY",o.TAY,o.IMP,2},{"LDA",o.LDA,o.IMM,2},{"TAX",o.TAX,o.IMP,2},{"???",o.XXX,o.IMP,2},{"LDY",o.LDY,o.ABS,4},{"LDA",o.LDA,o.ABS,4},{"LDX",o.LDX,o.ABS,4},{"???",o.XXX,o.IMP,4},
		{"BCS",o.BCS,o.REL,2},{"LDA",o.LDA,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,5},{"LDY",o.LDY,o.ZPX,4},{"LDA",o.LDA,o.ZPX,4},{"LDX",o.LDX,o.ZPY,4},{"???",o.XXX,o.IMP,4},{"CLV",o.CLV,o.IMP,2},{"LDA",o.LDA,o.ABY,4},{"TSX",o.TSX,o.IMP,2},{"???",o.XXX,o.IMP,4},{"LDY",o.LDY,o.ABX,4},{"LDA",o.LDA,o.ABX,4},{"LDX",o.LDX,o.ABY,4},{"???",o.XXX,o.IMP,4},
		{"CPY",o.CPY,o.IMM,2},{"CMP",o.CMP,o.IZX,6},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,8},{"CPY",o.CPY,o.ZP0,3},{"CMP",o.CMP,o.ZP0,3},{"DEC",o.DEC,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"INY",o.INY,o.IMP,2},{"CMP",o.CMP,o.IMM,2},{"DEX",o.DEX,o.IMP,2},{"???",o.XXX,o.IMP,2},{"CPY",o.CPY,o.ABS,4},{"CMP",o.CMP,o.ABS,4},{"DEC",o.DEC,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BNE",o.BNE,o.REL,2},{"CMP",o.CMP,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"CMP",o.CMP,o.ZPX,4},{"DEC",o.DEC,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"CLD",o.CLD,o.IMP,2},{"CMP",o.CMP,o.ABY,4},{"NOP",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"CMP",o.CMP,o.ABX,4},{"DEC",o.DEC,o.ABX,7},{"???",o.XXX,o.IMP,7},
		{"CPX",o.CPX,o.IMM,2},{"SBC",o.SBC,o.IZX,6},{"???",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,8},{"CPX",o.CPX,o.ZP0,3},{"SBC",o.SBC,o.ZP0,3},{"INC",o.INC,o.ZP0,5},{"???",o.XXX,o.IMP,5},{"INX",o.INX,o.IMP,2},{"SBC",o.SBC,o.IMM,2},{"NOP",o.NOP,o.IMP,2},{"???",o.SBC,o.IMP,2},{"CPX",o.CPX,o.ABS,4},{"SBC",o.SBC,o.ABS,4},{"INC",o.INC,o.ABS,6},{"???",o.XXX,o.IMP,6},
		{"BEQ",o.BEQ,o.REL,2},{"SBC",o.SBC,o.IZY,5},{"???",o.XXX,o.IMP,2},{"???",o.XXX,o.IMP,8},{"???",o.NOP,o.IMP,4},{"SBC",o.SBC,o.ZPX,4},{"INC",o.INC,o.ZPX,6},{"???",o.XXX,o.IMP,6},{"SED",o.SED,o.IMP,2},{"SBC",o.SBC,o.ABY,4},{"NOP",o.NOP,o.IMP,2},{"???",o.XXX,o.IMP,7},{"???",o.NOP,o.IMP,4},{"SBC",o.SBC,o.ABX,4},{"INC",o.INC,o.ABX,7},{"???",o.XXX,o.IMP,7},
  }

  return o
}

func (o *Olc6502) read(a uint16) uint8 {
  return o.bus.read(a, false)
}

func (o *Olc6502) write(a uint16, d uint8) {
  o.bus.write(a, d)
}

func (o *Olc6502) GetFlag(f string) uint8 {
  if (o.status & o.FLAGS6502[f]) > 0 {
    return 1
  } else {
    return 0
  }
}

func (o *Olc6502) SetFlag(f string, v bool) {
  if v {
    o.status |= o.FLAGS6502[f]
  } else {
    o.status &= ^o.FLAGS6502[f]
  } 
}

// Necessary functions
func (o *Olc6502) complete() bool {
  return o.cycles == 0
}


func (o *Olc6502) clock() {
  if (o.cycles == 0) {
    o.opcode = o.read(o.pc)
    o.SetFlag("U", true)
    o.pc++

    // Get starting number of cycles
    o.cycles = o.lookup[o.opcode].cycles

    additional_cycle1 := o.lookup[o.opcode].addrmode(o)
    additional_cycle2 := o.lookup[o.opcode].operate(o)  

    o.cycles += (additional_cycle1 & additional_cycle2)
  }

  o.cycles--
}
func (o *Olc6502) reset() {
  o.a = 0
  o.x = 0 
  o.y = 0
  o.stkp = 0xFD
  o.status = 0x00 | o.FLAGS6502["U"]

  o.addr_abs = 0xFFFC
  lo := uint16(o.read(o.addr_abs + 0))
  hi := uint16(o.read(o.addr_abs + 1))

  o.pc = (hi << 8) | lo

  o.addr_rel = 0x0000
  o.addr_abs = 0x0000
  o.fetched = 0x00

  o.cycles = 8
}
func (o *Olc6502) irq() {
  if o.GetFlag("T") == 0 {
    o.write(0x0100 + uint16(o.stkp), uint8((o.pc >> 8) & 0x00FF))
    o.stkp--
    o.write(0x0100 + uint16(o.stkp), uint8(o.pc >> 8 & 0x00FF))
    o.stkp--

    o.SetFlag("B", false)
    o.SetFlag("U", true)
    o.SetFlag("I", true)
    o.write(0x0100 + uint16(o.stkp), o.status)
    o.stkp--

    o.addr_abs = 0xFFFE
    lo := uint16(o.read(o.addr_abs + 0))
    hi := uint16(o.read(o.addr_abs + 1))
    o.pc = (hi << 8) | lo

    o.cycles = 7
  }
}
func (o *Olc6502) nmi() {
    o.write(0x0100 + uint16(o.stkp), uint8(o.pc >> 8 & 0x00FF))
    o.stkp--
    o.write(0x0100 + uint16(o.stkp), uint8(o.pc >> 8 & 0x00FF))
    o.stkp--

    o.SetFlag("B", false)
    o.SetFlag("U", true)
    o.SetFlag("I", true)
    o.write(0x0100 + uint16(o.stkp), o.status)
    o.stkp--

    o.addr_abs = 0xFFFA
    lo := uint16(o.read(o.addr_abs + 0))
    hi := uint16(o.read(o.addr_abs + 1))
    o.pc = (hi << 8) | lo

    o.cycles = 8
}

// Fetching data
func (o *Olc6502) fetch() uint8 {
  if (reflect.ValueOf(o.lookup[o.opcode].addrmode).Pointer() != reflect.ValueOf(o.IMP).Pointer()) {
    o.fetched = o.read(o.addr_abs)
  }
  return o.fetched
}

func (o *Olc6502) disassemble(nStart uint16, nStop uint16) map[uint16] string {
  addr := uint32(nStart)
  var value, lo, hi uint8 = 0x00, 0x00, 0x00
  var mapLines = make(map[uint16]string)
  var line_addr uint16 = 0

  for addr <= uint32(nStop) {
    line_addr = uint16(addr)

    // Prefix line with instruction address
    sInst := "$" + hex(addr, 4) + ": "

    // Read the instruction
    opcode := uint8(o.bus.read(uint16(addr), true))
    addr++
    sInst += o.lookup[opcode].name + " "

    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.IMP) {
      sInst += " {IMP}"
    } 
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.IMM) {
      value = o.bus.read(uint16(addr), true); addr++
      sInst += "#$" + hex(value, 2) + " {IMM}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ZP0) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + hex(lo, 2) + " {ZP0}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ZPX) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + hex(lo, 2) + ", X {ZPX}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ZPY) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + hex(lo, 2) + ", Y {ZPY}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.IZX) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "($" + hex(lo, 2) + ", X) {IZX}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.IZY) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "($" + hex(lo, 2) + ", Y) {IZY}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ABS) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = o.bus.read(uint16(addr), true); addr++
      sInst += "$" + hex(((hi << 8) | lo), 4) + " {ABS}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ABX) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = o.bus.read(uint16(addr), true); addr++
      sInst += "$" + hex(((hi << 8) | lo), 4) + ", X {ABX}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.ABY) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = o.bus.read(uint16(addr), true); addr++
      sInst += "$" + hex(((hi << 8) | lo), 4) + ", Y {ABY}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.IND) {
      lo = o.bus.read(uint16(addr), true); addr++
      hi = o.bus.read(uint16(addr), true); addr++
      sInst += "($" + hex(((hi << 8) | lo), 4) + ") {IND}"
    }
    if findFunc(o.lookup[opcode].addrmode) == findFunc(o.REL) {
      value = o.bus.read(uint16(addr), true); addr++
      sInst += "$" + hex(value, 2) + " [$" + hex(uint8(addr) + value, 4) + "] {REL}"
    }

    mapLines[line_addr] = sInst + "\n"
  }

  return mapLines
}
