package olc6502

import (
	"reflect"
  "nes-emulator/helpers"
)

type Olc6502 struct {
  ram *[64*1024]uint8
  A uint8 
  X uint8
  Y uint8
  Stkp uint8
  Pc uint16
  Status uint8
  fetched uint8
  addr_abs uint16
  addr_rel uint16
  opcode uint8
  Cycles uint8
}

type instruction struct {
  name string
  operate func(o *Olc6502) uint8
  addrmode func(o *Olc6502) uint8
  cycles uint8
}

var lookup []instruction
var Flags6502 map[string] uint8

func init() {
  Flags6502 = map[string] uint8{ 
    "C": 1 << 0, // 1
    "Z": 1 << 1, // 2
    "I": 1 << 2, // 4
    "D": 1 << 3, // 8
    "B": 1 << 4, // 16
    "U": 1 << 5, // 32
    "V": 1 << 6, // 64
    "N": 1 << 7, // 128
  }

  lookup = []instruction{
    {"BRK",BRK,IMM,7},{"ORA",ORA,IZX,6},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,3},{"ORA",ORA,ZP0,3},{"ASL",ASL,ZP0,5},{"???",XXX,IMP,5},{"PHP",PHP,IMP,3},{"ORA",ORA,IMM,2},{"ASL",ASL,IMP,2},{"???",XXX,IMP,2},{"???",NOP,IMP,4},{"ORA",ORA,ABS,4},{"ASL",ASL,ABS,6},{"???",XXX,IMP,6},
		{"BPL",BPL,REL,2},{"ORA",ORA,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"ORA",ORA,ZPX,4},{"ASL",ASL,ZPX,6},{"???",XXX,IMP,6},{"CLC",CLC,IMP,2},{"ORA",ORA,ABY,4},{"???",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"ORA",ORA,ABX,4},{"ASL",ASL,ABX,7},{"???",XXX,IMP,7},
		{"JSR",JSR,ABS,6},{"AND",AND,IZX,6},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"BIT",BIT,ZP0,3},{"AND",AND,ZP0,3},{"ROL",ROL,ZP0,5},{"???",XXX,IMP,5},{"PLP",PLP,IMP,4},{"AND",AND,IMM,2},{"ROL",ROL,IMP,2},{"???",XXX,IMP,2},{"BIT",BIT,ABS,4},{"AND",AND,ABS,4},{"ROL",ROL,ABS,6},{"???",XXX,IMP,6},
		{"BMI",BMI,REL,2},{"AND",AND,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"AND",AND,ZPX,4},{"ROL",ROL,ZPX,6},{"???",XXX,IMP,6},{"SEC",SEC,IMP,2},{"AND",AND,ABY,4},{"???",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"AND",AND,ABX,4},{"ROL",ROL,ABX,7},{"???",XXX,IMP,7},
		{"RTI",RTI,IMP,6},{"EOR",EOR,IZX,6},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,3},{"EOR",EOR,ZP0,3},{"LSR",LSR,ZP0,5},{"???",XXX,IMP,5},{"PHA",PHA,IMP,3},{"EOR",EOR,IMM,2},{"LSR",LSR,IMP,2},{"???",XXX,IMP,2},{"JMP",JMP,ABS,3},{"EOR",EOR,ABS,4},{"LSR",LSR,ABS,6},{"???",XXX,IMP,6},
		{"BVC",BVC,REL,2},{"EOR",EOR,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"EOR",EOR,ZPX,4},{"LSR",LSR,ZPX,6},{"???",XXX,IMP,6},{"CLI",CLI,IMP,2},{"EOR",EOR,ABY,4},{"???",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"EOR",EOR,ABX,4},{"LSR",LSR,ABX,7},{"???",XXX,IMP,7},
		{"RTS",RTS,IMP,6},{"ADC",ADC,IZX,6},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,3},{"ADC",ADC,ZP0,3},{"ROR",ROR,ZP0,5},{"???",XXX,IMP,5},{"PLA",PLA,IMP,4},{"ADC",ADC,IMM,2},{"ROR",ROR,IMP,2},{"???",XXX,IMP,2},{"JMP",JMP,IND,5},{"ADC",ADC,ABS,4},{"ROR",ROR,ABS,6},{"???",XXX,IMP,6},
		{"BVS",BVS,REL,2},{"ADC",ADC,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"ADC",ADC,ZPX,4},{"ROR",ROR,ZPX,6},{"???",XXX,IMP,6},{"SEI",SEI,IMP,2},{"ADC",ADC,ABY,4},{"???",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"ADC",ADC,ABX,4},{"ROR",ROR,ABX,7},{"???",XXX,IMP,7},
		{"???",NOP,IMP,2},{"STA",STA,IZX,6},{"???",NOP,IMP,2},{"???",XXX,IMP,6},{"STY",STY,ZP0,3},{"STA",STA,ZP0,3},{"STX",STX,ZP0,3},{"???",XXX,IMP,3},{"DEY",DEY,IMP,2},{"???",NOP,IMP,2},{"TXA",TXA,IMP,2},{"???",XXX,IMP,2},{"STY",STY,ABS,4},{"STA",STA,ABS,4},{"STX",STX,ABS,4},{"???",XXX,IMP,4},
		{"BCC",BCC,REL,2},{"STA",STA,IZY,6},{"???",XXX,IMP,2},{"???",XXX,IMP,6},{"STY",STY,ZPX,4},{"STA",STA,ZPX,4},{"STX",STX,ZPY,4},{"???",XXX,IMP,4},{"TYA",TYA,IMP,2},{"STA",STA,ABY,5},{"TXS",TXS,IMP,2},{"???",XXX,IMP,5},{"???",NOP,IMP,5},{"STA",STA,ABX,5},{"???",XXX,IMP,5},{"???",XXX,IMP,5},
		{"LDY",LDY,IMM,2},{"LDA",LDA,IZX,6},{"LDX",LDX,IMM,2},{"???",XXX,IMP,6},{"LDY",LDY,ZP0,3},{"LDA",LDA,ZP0,3},{"LDX",LDX,ZP0,3},{"???",XXX,IMP,3},{"TAY",TAY,IMP,2},{"LDA",LDA,IMM,2},{"TAX",TAX,IMP,2},{"???",XXX,IMP,2},{"LDY",LDY,ABS,4},{"LDA",LDA,ABS,4},{"LDX",LDX,ABS,4},{"???",XXX,IMP,4},
		{"BCS",BCS,REL,2},{"LDA",LDA,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,5},{"LDY",LDY,ZPX,4},{"LDA",LDA,ZPX,4},{"LDX",LDX,ZPY,4},{"???",XXX,IMP,4},{"CLV",CLV,IMP,2},{"LDA",LDA,ABY,4},{"TSX",TSX,IMP,2},{"???",XXX,IMP,4},{"LDY",LDY,ABX,4},{"LDA",LDA,ABX,4},{"LDX",LDX,ABY,4},{"???",XXX,IMP,4},
		{"CPY",CPY,IMM,2},{"CMP",CMP,IZX,6},{"???",NOP,IMP,2},{"???",XXX,IMP,8},{"CPY",CPY,ZP0,3},{"CMP",CMP,ZP0,3},{"DEC",DEC,ZP0,5},{"???",XXX,IMP,5},{"INY",INY,IMP,2},{"CMP",CMP,IMM,2},{"DEX",DEX,IMP,2},{"???",XXX,IMP,2},{"CPY",CPY,ABS,4},{"CMP",CMP,ABS,4},{"DEC",DEC,ABS,6},{"???",XXX,IMP,6},
		{"BNE",BNE,REL,2},{"CMP",CMP,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"CMP",CMP,ZPX,4},{"DEC",DEC,ZPX,6},{"???",XXX,IMP,6},{"CLD",CLD,IMP,2},{"CMP",CMP,ABY,4},{"NOP",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"CMP",CMP,ABX,4},{"DEC",DEC,ABX,7},{"???",XXX,IMP,7},
		{"CPX",CPX,IMM,2},{"SBC",SBC,IZX,6},{"???",NOP,IMP,2},{"???",XXX,IMP,8},{"CPX",CPX,ZP0,3},{"SBC",SBC,ZP0,3},{"INC",INC,ZP0,5},{"???",XXX,IMP,5},{"INX",INX,IMP,2},{"SBC",SBC,IMM,2},{"NOP",NOP,IMP,2},{"???",SBC,IMP,2},{"CPX",CPX,ABS,4},{"SBC",SBC,ABS,4},{"INC",INC,ABS,6},{"???",XXX,IMP,6},
		{"BEQ",BEQ,REL,2},{"SBC",SBC,IZY,5},{"???",XXX,IMP,2},{"???",XXX,IMP,8},{"???",NOP,IMP,4},{"SBC",SBC,ZPX,4},{"INC",INC,ZPX,6},{"???",XXX,IMP,6},{"SED",SED,IMP,2},{"SBC",SBC,ABY,4},{"NOP",NOP,IMP,2},{"???",XXX,IMP,7},{"???",NOP,IMP,4},{"SBC",SBC,ABX,4},{"INC",INC,ABX,7},{"???",XXX,IMP,7},
  }
}

func NewOlc6502(ram *[64*1024]uint8) Olc6502 {
  o := Olc6502{}
  o.ram = ram

  
  o.A = 0x00
  o.X = 0x00
  o.Y = 0x00
  o.Stkp = 0x00
  o.Pc = 0x0000
  o.Status = 0x00
  o.fetched = 0x00
  o.addr_abs = 0x000
  o.addr_rel = 0x00
  o.opcode = 0x00
  o.Cycles = 0


  return o
}

func (o *Olc6502) read(a uint16, bReadOnly ...bool) uint8 {
  if (a >= 0x0000 && a <= 0xFFFF) {
    return o.ram[a]
  }
  return 0x00
}

func (o *Olc6502) write(a uint16, d uint8) {
  if (a >= 0x0000 && a <= 0xFFFF){
    o.ram[a] = d
  }
}

func (o *Olc6502) GetFlag(f string) uint8 {
  if (o.Status & Flags6502[f]) > 0 {
    return 1
  } else {
    return 0
  }
}

func (o *Olc6502) SetFlag(f string, v bool) {
  if v {
    o.Status |= Flags6502[f]
  } else {
    o.Status &= ^Flags6502[f]
  } 
}

// Necessary functions
func (o *Olc6502) Complete() bool {
  return o.Cycles == 0
}


func (o *Olc6502) Clock() {
  if (o.Cycles == 0) {
    o.opcode = o.read(o.Pc)
    o.SetFlag("U", true)
    o.Pc++

    // Get starting number of cycles
    o.Cycles = lookup[o.opcode].cycles

    additional_cycle1 := lookup[o.opcode].addrmode(o)
    additional_cycle2 := lookup[o.opcode].operate(o)  

    o.Cycles += (additional_cycle1 & additional_cycle2)
  }

  o.Cycles--
}
func (o *Olc6502) Reset() {
  o.A = 0
  o.X = 0 
  o.Y = 0
  o.Stkp = 0xFD
  o.Status = 0x00 | Flags6502["U"]

  o.addr_abs = 0xFFFC
  lo := uint16(o.read(o.addr_abs + 0))
  hi := uint16(o.read(o.addr_abs + 1))

  o.Pc = (hi << 8) | lo

  o.addr_rel = 0x0000
  o.addr_abs = 0x0000
  o.fetched = 0x00

  o.Cycles = 8
}
func (o *Olc6502) Irq() {
  if o.GetFlag("T") == 0 {
    o.write(0x0100 + uint16(o.Stkp), uint8((o.Pc >> 8) & 0x00FF))
    o.Stkp--
    o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc >> 8 & 0x00FF))
    o.Stkp--

    o.SetFlag("B", false)
    o.SetFlag("U", true)
    o.SetFlag("I", true)
    o.write(0x0100 + uint16(o.Stkp), o.Status)
    o.Stkp--

    o.addr_abs = 0xFFFE
    lo := uint16(o.read(o.addr_abs + 0))
    hi := uint16(o.read(o.addr_abs + 1))
    o.Pc = (hi << 8) | lo

    o.Cycles = 7
  }
}
func (o *Olc6502) Nmi() {
    o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc >> 8 & 0x00FF))
    o.Stkp--
    o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc >> 8 & 0x00FF))
    o.Stkp--

    o.SetFlag("B", false)
    o.SetFlag("U", true)
    o.SetFlag("I", true)
    o.write(0x0100 + uint16(o.Stkp), o.Status)
    o.Stkp--

    o.addr_abs = 0xFFFA
    lo := uint16(o.read(o.addr_abs + 0))
    hi := uint16(o.read(o.addr_abs + 1))
    o.Pc = (hi << 8) | lo

    o.Cycles = 8
}

// Fetching data
func (o *Olc6502) fetch() uint8 {
  if (reflect.ValueOf(lookup[o.opcode].addrmode).Pointer() != reflect.ValueOf(IMP).Pointer()) {
    o.fetched = o.read(o.addr_abs)
  }
  return o.fetched
}

func (o *Olc6502) Disassemble(nStart uint16, nStop uint16) map[uint16] string {
  addr := uint32(nStart)
  var value, lo, hi uint8 = 0x00, 0x00, 0x00
  var mapLines = make(map[uint16]string)
  var line_addr uint16 = 0

  for addr <= uint32(nStop) {
    line_addr = uint16(addr)

    // Prefix line with instruction address
    sInst := "$" + helpers.Hex(addr, 4) + ": "

    // Read the instruction
    opcode := uint8(o.read(uint16(addr), true))
    addr++
    sInst += lookup[opcode].name + " "

    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(IMP) {
      sInst += " {IMP}"
    } 
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(IMM) {
      value = o.read(uint16(addr), true); addr++
      sInst += "#$" + helpers.Hex(value, 2) + " {IMM}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ZP0) {
      lo = o.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + helpers.Hex(lo, 2) + " {ZP0}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ZPX) {
      lo = o.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + helpers.Hex(lo, 2) + ", X {ZPX}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ZPY) {
      lo = o.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "$" + helpers.Hex(lo, 2) + ", Y {ZPY}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(IZX) {
      lo = o.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "($" + helpers.Hex(lo, 2) + ", X) {IZX}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(IZY) {
      lo = o.read(uint16(addr), true); addr++
      hi = 0x00
      sInst += "($" + helpers.Hex(lo, 2) + ", Y) {IZY}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ABS) {
      lo = o.read(uint16(addr), true); addr++
      hi = o.read(uint16(addr), true); addr++
      sInst += "$" + helpers.Hex(((hi << 8) | lo), 4) + " {ABS}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ABX) {
      lo = o.read(uint16(addr), true); addr++
      hi = o.read(uint16(addr), true); addr++
      sInst += "$" + helpers.Hex(((hi << 8) | lo), 4) + ", X {ABX}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(ABY) {
      lo = o.read(uint16(addr), true); addr++
      hi = o.read(uint16(addr), true); addr++
      sInst += "$" + helpers.Hex(((hi << 8) | lo), 4) + ", Y {ABY}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(IND) {
      lo = o.read(uint16(addr), true); addr++
      hi = o.read(uint16(addr), true); addr++
      sInst += "($" + helpers.Hex(((hi << 8) | lo), 4) + ") {IND}"
    }
    if helpers.FindFunc(lookup[opcode].addrmode) == helpers.FindFunc(REL) {
      value = o.read(uint16(addr), true); addr++
      sInst += "$" + helpers.Hex(value, 2) + " [$" + helpers.Hex(uint8(addr) + value, 4) + "] {REL}"
    }

    mapLines[line_addr] = sInst + "\n"
  }

  return mapLines
}
