package olc6502

import (
  "nes-emulator/helpers"
)

// Opcodes
func (cpu *Olc6502) ADC(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.a + o.fetched + o.GetFlag("C"))
  o.SetFlag("C", temp > 255)
  o.SetFlag("Z", temp & 0x00FF == 0)
  o.SetFlag("Z", temp & 0x80 != 0)
  o.SetFlag("V", (uint16(o.a) ^ uint16(o.fetched)) & (uint16(o.a) ^ temp) & 0x0080 != 0)
  o.a = uint8(temp & 0x00FF)

  return 1
}
func (cpu *Olc6502) AND(o *Olc6502) uint8 {
  o.fetch()
  o.a = o.a & o.fetched
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", (o.a & 0x80) != 0)

  return 1
}
func (cpu *Olc6502) ASL(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.fetched) << 1
  o.SetFlag("C", temp & 0xFF00 > 0)
  o.SetFlag("Z", temp & 0x00FF == 0x00)
  o.SetFlag("N", temp & 0x80 != 0)

  if helpers.FindFunc(o.lookup[o.opcode].addrmode) == helpers.FindFunc(o.IMP) {
    o.a = uint8(temp & 0x00FF)
  } else {
    o.write(o.addr_abs, uint8(temp & 0x00FF))
  }

  return 0
}
func (cpu *Olc6502) BCC(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BCS(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BEQ(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BIT(o *Olc6502) uint8 {
  o.fetch()
  temp := o.a & o.fetched
  o.SetFlag("Z", temp & 0x00FF == 0x00)
  o.SetFlag("N", o.fetched & (1 << 7) != 0)
  o.SetFlag("V", o.fetched & (1 << 6) != 0)
  return 0
}
func (cpu *Olc6502) BMI(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BNE(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BPL(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BRK(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BVC(o *Olc6502) uint8 {
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
func (cpu *Olc6502) BVS(o *Olc6502) uint8 {
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
func (cpu *Olc6502) CLC(o *Olc6502) uint8 {
  o.SetFlag("C", false)
  return 0
}
func (cpu *Olc6502) CLD(o *Olc6502) uint8 {
  o.SetFlag("D", false)
  return 0
}
func (cpu *Olc6502) CLI(o *Olc6502) uint8 {
  o.SetFlag("I", false)
  return 0
}
func (cpu *Olc6502) CLV(o *Olc6502) uint8 {
  o.SetFlag("V", false)
  return 0
}
func (cpu *Olc6502) CMP(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.a) - uint16(o.fetched)
  o.SetFlag("C", o.a >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1
}
func (cpu *Olc6502) CPX(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.x) - uint16(o.fetched)
  o.SetFlag("C", o.x >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1

}
func (cpu *Olc6502) CPY(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.y) - uint16(o.fetched)
  o.SetFlag("C", o.y >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1

}
func (cpu *Olc6502) DEC(o *Olc6502) uint8 {
  o.fetch()
  temp := o.fetched - 1
  o.write(o.addr_abs, temp & 0x00FF)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 0
}
func (cpu *Olc6502) DEX(o *Olc6502) uint8 {
  o.x--
  o.SetFlag("Z", o.x == 0x00)
  o.SetFlag("N", o.x & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) DEY(o *Olc6502) uint8 {
  o.y--
  o.SetFlag("Z", o.y == 0x00)
  o.SetFlag("N", o.y & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) EOR(o *Olc6502) uint8 {
  o.fetch()
  o.a = o.a ^ o.fetched
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", o.a & 0x80 != 0)
  return 1
}
func (cpu *Olc6502) INC(o *Olc6502) uint8 {
  o.fetch()
  temp := o.fetched + 1
  o.write(o.addr_abs, temp & 0x00FF)
  o.SetFlag("Z", temp * 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 0
}
func (cpu *Olc6502) INX(o *Olc6502) uint8 {
  o.x++
  o.SetFlag("Z", o.x == 0x00)
  o.SetFlag("N", o.x & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) INY(o *Olc6502) uint8 {
  o.y++
  o.SetFlag("Z", o.y == 0x00)
  o.SetFlag("N", o.y & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) JMP(o *Olc6502) uint8 {
  o.pc = o.addr_abs
  return 0
}
func (cpu *Olc6502) JSR(o *Olc6502) uint8 {
  o.pc--

  o.write(0x0100 + uint16(o.stkp), uint8(o.pc >> 8 & 0x00FF))
  o.stkp--
  o.write(0x0100 + uint16(o.stkp), uint8(o.pc)  & 0x00FF)
  o.stkp--

  o.pc = o.addr_abs
  return 0
}
func (cpu *Olc6502) LDA(o *Olc6502) uint8 {
  o.fetch()
  o.a = o.fetched
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", o.a & 0x80 != 0)
  return 1
}
func (cpu *Olc6502) LDX(o *Olc6502) uint8 {
  o.fetch()
  o.x = o.fetched
  o.SetFlag("Z", o.x == 0x00)
  o.SetFlag("N", o.x & 0x80 != 0)
  return 1
}
func (cpu *Olc6502) LDY(o *Olc6502) uint8 {
  o.fetch()
  o.y = o.fetched
  o.SetFlag("Z", o.y == 0x00)
  o.SetFlag("N", o.y & 0x80 != 0)
  return 1
}
func (cpu *Olc6502) LSR(o *Olc6502) uint8 {
  o.fetch()
  o.SetFlag("C", o.fetched & 0x0001 != 0)

  temp := o.fetched >> 1
  
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)

  if helpers.FindFunc(o.lookup[o.opcode].addrmode) == helpers.FindFunc(o.IMP) {
    o.a = temp & 0x00FF
  } else {
    o.write(o.addr_abs, temp & 0x00FF)
  }
  return 0
}
func (cpu *Olc6502) NOP(o *Olc6502) uint8 {
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
func (cpu *Olc6502) ORA(o *Olc6502) uint8 {
  o.fetch()
  o.a = o.a | o.fetched
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", o.a & 0x80 != 0)
  return 1
}
func (cpu *Olc6502) PHA(o *Olc6502) uint8 {
  o.write(0x0100 + uint16(o.stkp), o.a)
  o.stkp--
  return 0
}
func (cpu *Olc6502) PHP(o *Olc6502) uint8 {
  o.write(0x0100 + uint16(o.stkp), o.status | o.FLAGS6502["B"] | o.FLAGS6502["U"])
  o.SetFlag("B", false)
  o.SetFlag("U", false)
  o.stkp--
  return 0
}
func (cpu *Olc6502) PLA(o *Olc6502) uint8 {
  o.stkp++
  o.a = o.read(0x0100 + uint16(o.stkp))
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", (o.a & 0x80) != 0)
  return 0
}
func (cpu *Olc6502) PLP(o *Olc6502) uint8 {
  o.stkp++
  o.status = o.read(0x0100 + uint16(o.stkp))
  o.SetFlag("U", true)
  return 0
}
func (cpu *Olc6502) ROL(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.fetched << 1) | uint16(o.GetFlag("C"))
  o.SetFlag("C", temp & 0xFF00 != 0)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)

  if helpers.FindFunc(o.lookup[o.opcode].addrmode) == helpers.FindFunc(o.IMP) {
    o.a = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }
  return 0
}
func (cpu *Olc6502) ROR(o *Olc6502) uint8 {
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
  if helpers.FindFunc(o.lookup[o.opcode].addrmode) == helpers.FindFunc(o.IMP) {
    o.a = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }

  return 0
}
func (cpu *Olc6502) RTI(o *Olc6502) uint8 {
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
func (cpu *Olc6502) RTS(o *Olc6502) uint8 {
  o.stkp++
  o.pc = uint16(o.read(0x0100 + uint16(o.stkp)))
  o.stkp++
  o.pc |= uint16(o.read(0x0100 + uint16(o.stkp)) << 8)
  o.pc++
  return 0
}
func (cpu *Olc6502) SBC(o *Olc6502) uint8 {
  o.fetch()
  value := uint16(o.fetched ^ 0x00FF)
  temp := uint16(uint16(o.a) + value + uint16(o.GetFlag("C")))
  o.SetFlag("C", temp > 0xFF00)
  o.SetFlag("Z", (temp & 0x00FF) == 0)
  o.SetFlag("N", (temp & 0x0080) != 0)
  o.SetFlag("V", (temp ^ uint16(o.a) & (temp ^ value)) & 0x0080 != 0)

  o.a = uint8(temp & 0x00FF)

  return 1
}
func (cpu *Olc6502) SEC(o *Olc6502) uint8 {
  o.SetFlag("C", true)
  return 0
}
func (cpu *Olc6502) SED(o *Olc6502) uint8 {
  o.SetFlag("D", true)
  return 0
}
func (cpu *Olc6502) SEI(o *Olc6502) uint8 {
  o.SetFlag("I", true)
  return 0
}
func (cpu *Olc6502) STA(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.a)
  return 0
}
func (cpu *Olc6502) STX(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.x)
  return 0
}
func (cpu *Olc6502) STY(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.y)
  return 0
}
func (cpu *Olc6502) TAX(o *Olc6502) uint8 {
  o.x = o.a
  o.SetFlag("Z", o.x == 0x00)
  o.SetFlag("N", o.x & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) TAY(o *Olc6502) uint8 {
  o.y = o.a
  o.SetFlag("Z", o.y == 0x00)
  o.SetFlag("N", o.y & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) TSX(o *Olc6502) uint8 {
  o.x = o.stkp
  o.SetFlag("Z", o.x == 0x00)
  o.SetFlag("N", o.x & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) TXA(o *Olc6502) uint8 {
  o.a = o.x
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", o.a & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) TXS(o *Olc6502) uint8 {
  o.stkp = o.a
  return 0
}
func (cpu *Olc6502) TYA(o *Olc6502) uint8 {
  o.a = o.y
  o.SetFlag("Z", o.a == 0x00)
  o.SetFlag("N", o.a & 0x80 != 0)
  return 0
}
func (cpu *Olc6502) XXX(o *Olc6502) uint8 {
  return 0
}
