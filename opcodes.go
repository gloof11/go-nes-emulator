package main

import (
  "reflect"
)

// Opcodes
func (o *Olc6502) ADC() uint8 {
  o.fetch()
  temp := uint16(o.a + o.fetched + o.GetFlag("C"))
  if temp > 255 {
    o.SetFlag("C", true)
  }
  if (temp & 0x00FF) == 0 {
    o.SetFlag("Z", true)
  }
  if (temp & 0x80) != 0 {
    o.SetFlag("N", true)
  }
  if (uint16(o.a) ^ uint16(o.fetched)) & (uint16(o.a) ^ temp) & 0x0080 != 0 {
    o.SetFlag("V", true)
  }

  o.a = uint8(temp & 0x00FF)

  return 1
}
func (o *Olc6502) AND() uint8 {
  o.fetch()
  o.a = o.a & o.fetched
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if (o.a & 0x80) != 0 {
    o.SetFlag("N", true)
  }

  return 1
}
func (o *Olc6502) ASL() uint8 {
  o.fetch()
  temp := uint16(o.fetched) << 1
  if temp & 0xFF00 > 0 {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x00 {
    o.SetFlag("Z", true)
  }
  if temp & 0x80 != 0 {
    o.SetFlag("N", true)
  }

  if reflect.ValueOf(o.lookup[o.opcode].addrmode).Pointer() == reflect.ValueOf(o.IMP).Pointer() {
    o.a = uint8(temp & 0x00FF)
  } else {
    o.write(o.addr_abs, uint8(temp & 0x00FF))
  }

  return 0
}
func (o *Olc6502) BCC() uint8 {
    if o.GetFlag("C") == 0 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BCS() uint8 {
  if o.GetFlag("C") == 1 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BEQ() uint8 {
    if o.GetFlag("Z") == 1 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BIT() uint8 {
  o.fetch()
  temp := o.a & o.fetched
  if temp & 0x00FF == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.fetched & (1 << 7) != 0 {
    o.SetFlag("N", true)
  }
  if o.fetched & (1 << 6) != 0 {
    o.SetFlag("V", true)
  }
  return 0
}
func (o *Olc6502) BMI() uint8 {
    if o.GetFlag("N") == 1 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BNE() uint8 {
    if o.GetFlag("Z") == 0 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BPL() uint8 {
    if o.GetFlag("N") == 0 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BRK() uint8 {
  o.pc++

  o.SetFlag("I", true)
  o.write(0x0100 + uint16(o.stkp), uint8((o.pc >> 8) & 0x00FF))
  o.stkp--
  o.write(0x0100 + uint16(o.stkp), uint8(o.pc & 0x00FF))
  o.stkp--

  o.SetFlag("B", true)
  o.write(0x0100 + uint16(o.stkp), o.status)
  o.stkp--
  o.SetFlag("B", false)

  o.pc = uint16(o.read(0xFFFE)) | uint16(o.read(0xFFFF) << 8)
  return 0
}
func (o *Olc6502) BVC() uint8 {
    if o.GetFlag("V") == 0 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) BVS() uint8 {
    if o.GetFlag("V") == 1 {
    o.cycles++
    o.addr_abs = o.pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.pc & 0xFF00) {
      o.cycles++
    }

    o.pc = o.addr_abs
  }

  return 0
}
func (o *Olc6502) CLC() uint8 {
  o.SetFlag("C", false)
  return 0
}
func (o *Olc6502) CLD() uint8 {
  o.SetFlag("D", false)
  return 0
}
func (o *Olc6502) CLI() uint8 {
  o.SetFlag("I", false)
  return 0
}
func (o *Olc6502) CLV() uint8 {
  o.SetFlag("V", false)
  return 0
}
func (o *Olc6502) CMP() uint8 {
  o.fetch()
  temp := uint16(o.a) - uint16(o.fetched)
  if o.a >= o.fetched {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) CPX() uint8 {
  o.fetch()
  temp := uint16(o.x) - uint16(o.fetched)
  if o.x >= o.fetched {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  return 1

}
func (o *Olc6502) CPY() uint8 {
  o.fetch()
  temp := uint16(o.y) - uint16(o.fetched)
  if o.y >= o.fetched {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  return 1

}
func (o *Olc6502) DEC() uint8 {
  o.fetch()
  temp := o.fetched - 1
  o.write(o.addr_abs, temp & 0x00FF)
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) DEX() uint8 {
  o.x--
  if o.x == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.x & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) DEY() uint8 {
  o.y--
  if o.y == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.y & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) EOR() uint8 {
  o.fetch()
  o.a = o.a ^ o.fetched
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.a & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) INC() uint8 {
  o.fetch()
  temp := o.fetched + 1
  o.write(o.addr_abs, temp & 0x00FF)
  if temp * 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) INX() uint8 {
  o.x++
  if o.x == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.x & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) INY() uint8 {
  o.y++
  if o.y == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.y & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) JMP() uint8 {
  o.pc = o.addr_abs
  return 0
}
func (o *Olc6502) JSR() uint8 {
  o.pc--

  o.write(0x0100 + uint16(o.stkp), uint8(o.pc >> 8 & 0x00FF))
  o.stkp--
  o.write(0x0100 + uint16(o.stkp), uint8(o.pc)  & 0x00FF)
  o.stkp--

  o.pc = o.addr_abs
  return 0
}
func (o *Olc6502) LDA() uint8 {
  o.fetch()
  o.a = o.fetched
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.a & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) LDX() uint8 {
  o.fetch()
  o.x = o.fetched
  if o.x == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.x & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) LDY() uint8 {
  o.fetch()
  o.y = o.fetched
  if o.y == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.y & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) LSR() uint8 {
  o.fetch()
  if o.fetched & 0x0001 != 0 {
    o.SetFlag("C", true)
  }

  temp := o.fetched >> 1
  
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }

  if reflect.ValueOf(o.lookup[o.opcode].addrmode).Pointer() == reflect.ValueOf(o.IMP).Pointer() {
    o.a = temp & 0x00FF
  } else {
    o.write(o.addr_abs, temp & 0x00FF)
  }
  return 0
}
func (o *Olc6502) NOP() uint8 {
  switch(o.opcode) {
    case 0x1C:
    case 0x3C:
    case 0x5C:
    case 0x7C:
    case 0xDC:
    case 0xFC:
      return 1
  }
  return 0
}
func (o *Olc6502) ORA() uint8 {
  o.fetch()
  o.a = o.a | o.fetched
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.a & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 1
}
func (o *Olc6502) PHA() uint8 {
  o.write(0x0100 + uint16(o.stkp), o.a)
  o.stkp--
  return 0
}
func (o *Olc6502) PHP() uint8 {
  o.write(0x0100 + uint16(o.stkp), o.status | o.FLAGS6502["B"] | o.FLAGS6502["U"])
  o.SetFlag("B", false)
  o.SetFlag("U", false)
  o.stkp--
  return 0
}
func (o *Olc6502) PLA() uint8 {
  o.stkp++
  o.a = o.read(0x0100 + uint16(o.stkp))
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if (o.a & 0x80) != 0 {
    o.SetFlag("N", true)
  } 
  return 0
}
func (o *Olc6502) PLP() uint8 {
  o.stkp++
  o.status = o.read(0x0100 + uint16(o.stkp))
  o.SetFlag("U", true)
  return 0
}
func (o *Olc6502) ROL() uint8 {
  o.fetch()
  temp := uint16(o.fetched << 1) | uint16(o.GetFlag("C"))
  if temp & 0xFF00 != 0 {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x0000 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }

  if reflect.ValueOf(o.lookup[o.opcode].addrmode).Pointer() == reflect.ValueOf(o.IMP).Pointer() {
    o.a = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }
  return 0
}
func (o *Olc6502) ROR() uint8 {
  o.fetch()
  temp := uint16(o.GetFlag("C") << 7) | (uint16(o.fetched) >> 1)
  if o.fetched & 0x01 != 0 {
    o.SetFlag("C", true)
  }
  if temp & 0x00FF == 0x00 {
    o.SetFlag("Z", true)
  }
  if temp & 0x0080 != 0 {
    o.SetFlag("N", true)
  }
  if reflect.ValueOf(o.lookup[o.opcode].addrmode).Pointer() == reflect.ValueOf(o.IMP).Pointer() {
    o.a = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }

  return 0
}
func (o *Olc6502) RTI() uint8 {
  o.stkp++
  o.status = o.read(0x0100 + uint16(o.stkp))
  o.status &= ^o.FLAGS6502["B"]
  o.status &= ^o.FLAGS6502["U"]

  o.stkp++
  o.pc = uint16(o.read(0x0100 + uint16(o.stkp)))
  o.stkp++
  o.pc |= uint16(o.read(0x0100 + uint16(o.stkp)) << 8)
  return 0
}
func (o *Olc6502) RTS() uint8 {
  o.stkp++
  o.pc = uint16(o.read(0x0100 + uint16(o.stkp)))
  o.stkp++
  o.pc |= uint16(o.read(0x0100 + uint16(o.stkp)) << 8)
  o.pc++
  return 0
}
func (o *Olc6502) SBC() uint8 {
  o.fetch()
  value := uint16(o.fetched ^ 0x00FF)
  temp := uint16(uint16(o.a) + value + uint16(o.GetFlag("C")))
  if temp > 0xFF00 {
    o.SetFlag("C", true)
  }
  if (temp & 0x00FF) == 0 {
    o.SetFlag("Z", true)
  }
  if (temp & 0x0080) != 0 {
    o.SetFlag("N", true)
  }
  if (temp ^ uint16(o.a) & (temp ^ value)) & 0x0080 != 0 {
    o.SetFlag("V", true)
  }

  o.a = uint8(temp & 0x00FF)

  return 1
}
func (o *Olc6502) SEC() uint8 {
  o.SetFlag("C", true)
  return 0
}
func (o *Olc6502) SED() uint8 {
  o.SetFlag("D", true)
  return 0
}
func (o *Olc6502) SEI() uint8 {
  o.SetFlag("I", true)
  return 0
}
func (o *Olc6502) STA() uint8 {
  o.write(o.addr_abs, o.a)
  return 0
}
func (o *Olc6502) STX() uint8 {
  o.write(o.addr_abs, o.x)
  return 0
}
func (o *Olc6502) STY() uint8 {
  o.write(o.addr_abs, o.y)
  return 0
}
func (o *Olc6502) TAX() uint8 {
  o.x = o.a
  if o.x == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.x & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) TAY() uint8 {
  o.y = o.a
  if o.y == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.y & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) TSX() uint8 {
  o.x = o.stkp
  if o.x == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.x & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) TXA() uint8 {
  o.a = o.x
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.a & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) TXS() uint8 {
  o.stkp = o.a
  return 0
}
func (o *Olc6502) TYA() uint8 {
  o.a = o.y
  if o.a == 0x00 {
    o.SetFlag("Z", true)
  }
  if o.a & 0x80 != 0 {
    o.SetFlag("N", true)
  }
  return 0
}
func (o *Olc6502) XXX() uint8 {
  return 0
}
