package olc6502

import (
  "nes-emulator/helpers"
)

// Opcodes
func ADC(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.A + o.fetched + o.GetFlag("C"))
  o.SetFlag("C", temp > 255)
  o.SetFlag("Z", temp & 0x00FF == 0)
  o.SetFlag("Z", temp & 0x80 != 0)
  o.SetFlag("V", (uint16(o.A) ^ uint16(o.fetched)) & (uint16(o.A) ^ temp) & 0x0080 != 0)
  o.A = uint8(temp & 0x00FF)

  return 1
}
func AND(o *Olc6502) uint8 {
  o.fetch()
  o.A = o.A & o.fetched
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", (o.A & 0x80) != 0)

  return 1
}
func ASL(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.fetched) << 1
  o.SetFlag("C", temp & 0xFF00 > 0)
  o.SetFlag("Z", temp & 0x00FF == 0x00)
  o.SetFlag("N", temp & 0x80 != 0)

  if helpers.FindFunc(lookup[o.opcode].addrmode) == helpers.FindFunc(IMP) {
    o.A = uint8(temp & 0x00FF)
  } else {
    o.write(o.addr_abs, uint8(temp & 0x00FF))
  }

  return 0
}
func BCC(o *Olc6502) uint8 {
    if o.GetFlag("C") == 0 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BCS(o *Olc6502) uint8 {
  if o.GetFlag("C") == 1 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BEQ(o *Olc6502) uint8 {
    if o.GetFlag("Z") == 1 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BIT(o *Olc6502) uint8 {
  o.fetch()
  temp := o.A & o.fetched
  o.SetFlag("Z", temp & 0x00FF == 0x00)
  o.SetFlag("N", o.fetched & (1 << 7) != 0)
  o.SetFlag("V", o.fetched & (1 << 6) != 0)
  return 0
}
func BMI(o *Olc6502) uint8 {
    if o.GetFlag("N") == 1 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BNE(o *Olc6502) uint8 {
    if o.GetFlag("Z") == 0 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BPL(o *Olc6502) uint8 {
    if o.GetFlag("N") == 0 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BRK(o *Olc6502) uint8 {
  o.Pc++

  o.SetFlag("I", true)
  o.write(0x0100 + uint16(o.Stkp), uint8((o.Pc >> 8) & 0x00FF))
  o.Stkp--
  o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc & 0x00FF))
  o.Stkp--

  o.SetFlag("B", true)
  o.write(0x0100 + uint16(o.Stkp), o.Status)
  o.Stkp--
  o.SetFlag("B", false)

  o.Pc = uint16(o.read(0xFFFE)) | uint16(o.read(0xFFFF) << 8)
  return 0
}
func BVC(o *Olc6502) uint8 {
    if o.GetFlag("V") == 0 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func BVS(o *Olc6502) uint8 {
    if o.GetFlag("V") == 1 {
    o.Cycles++
    o.addr_abs = o.Pc + o.addr_rel

    if (o.addr_abs & 0xFF00) != (o.Pc & 0xFF00) {
      o.Cycles++
    }

    o.Pc = o.addr_abs
  }

  return 0
}
func CLC(o *Olc6502) uint8 {
  o.SetFlag("C", false)
  return 0
}
func CLD(o *Olc6502) uint8 {
  o.SetFlag("D", false)
  return 0
}
func CLI(o *Olc6502) uint8 {
  o.SetFlag("I", false)
  return 0
}
func CLV(o *Olc6502) uint8 {
  o.SetFlag("V", false)
  return 0
}
func CMP(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.A) - uint16(o.fetched)
  o.SetFlag("C", o.A >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1
}
func CPX(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.X) - uint16(o.fetched)
  o.SetFlag("C", o.X >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1

}
func CPY(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.Y) - uint16(o.fetched)
  o.SetFlag("C", o.Y >= o.fetched)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 1

}
func DEC(o *Olc6502) uint8 {
  o.fetch()
  temp := o.fetched - 1
  o.write(o.addr_abs, temp & 0x00FF)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 0
}
func DEX(o *Olc6502) uint8 {
  o.X--
  o.SetFlag("Z", o.X == 0x00)
  o.SetFlag("N", o.X & 0x80 != 0)
  return 0
}
func DEY(o *Olc6502) uint8 {
  o.Y--
  o.SetFlag("Z", o.Y == 0x00)
  o.SetFlag("N", o.Y & 0x80 != 0)
  return 0
}
func EOR(o *Olc6502) uint8 {
  o.fetch()
  o.A = o.A ^ o.fetched
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", o.A & 0x80 != 0)
  return 1
}
func INC(o *Olc6502) uint8 {
  o.fetch()
  temp := o.fetched + 1
  o.write(o.addr_abs, temp & 0x00FF)
  o.SetFlag("Z", temp * 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)
  return 0
}
func INX(o *Olc6502) uint8 {
  o.X++
  o.SetFlag("Z", o.X == 0x00)
  o.SetFlag("N", o.X & 0x80 != 0)
  return 0
}
func INY(o *Olc6502) uint8 {
  o.Y++
  o.SetFlag("Z", o.Y == 0x00)
  o.SetFlag("N", o.Y & 0x80 != 0)
  return 0
}
func JMP(o *Olc6502) uint8 {
  o.Pc = o.addr_abs
  return 0
}
func JSR(o *Olc6502) uint8 {
  o.Pc--

  o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc >> 8 & 0x00FF))
  o.Stkp--
  o.write(0x0100 + uint16(o.Stkp), uint8(o.Pc)  & 0x00FF)
  o.Stkp--

  o.Pc = o.addr_abs
  return 0
}
func LDA(o *Olc6502) uint8 {
  o.fetch()
  o.A = o.fetched
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", o.A & 0x80 != 0)
  return 1
}
func LDX(o *Olc6502) uint8 {
  o.fetch()
  o.X = o.fetched
  o.SetFlag("Z", o.X == 0x00)
  o.SetFlag("N", o.X & 0x80 != 0)
  return 1
}
func LDY(o *Olc6502) uint8 {
  o.fetch()
  o.Y = o.fetched
  o.SetFlag("Z", o.Y == 0x00)
  o.SetFlag("N", o.Y & 0x80 != 0)
  return 1
}
func LSR(o *Olc6502) uint8 {
  o.fetch()
  o.SetFlag("C", o.fetched & 0x0001 != 0)

  temp := o.fetched >> 1
  
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)

  if helpers.FindFunc(lookup[o.opcode].addrmode) == helpers.FindFunc(IMP) {
    o.A = temp & 0x00FF
  } else {
    o.write(o.addr_abs, temp & 0x00FF)
  }
  return 0
}
func NOP(o *Olc6502) uint8 {
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
func ORA(o *Olc6502) uint8 {
  o.fetch()
  o.A = o.A | o.fetched
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", o.A & 0x80 != 0)
  return 1
}
func PHA(o *Olc6502) uint8 {
  o.write(0x0100 + uint16(o.Stkp), o.A)
  o.Stkp--
  return 0
}
func PHP(o *Olc6502) uint8 {
  o.write(0x0100 + uint16(o.Stkp), o.Status | Flags6502["B"] | Flags6502["U"])
  o.SetFlag("B", false)
  o.SetFlag("U", false)
  o.Stkp--
  return 0
}
func PLA(o *Olc6502) uint8 {
  o.Stkp++
  o.A = o.read(0x0100 + uint16(o.Stkp))
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", (o.A & 0x80) != 0)
  return 0
}
func PLP(o *Olc6502) uint8 {
  o.Stkp++
  o.Status = o.read(0x0100 + uint16(o.Stkp))
  o.SetFlag("U", true)
  return 0
}
func ROL(o *Olc6502) uint8 {
  o.fetch()
  temp := uint16(o.fetched << 1) | uint16(o.GetFlag("C"))
  o.SetFlag("C", temp & 0xFF00 != 0)
  o.SetFlag("Z", temp & 0x00FF == 0x0000)
  o.SetFlag("N", temp & 0x0080 != 0)

  if helpers.FindFunc(lookup[o.opcode].addrmode) == helpers.FindFunc(IMP) {
    o.A = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }
  return 0
}
func ROR(o *Olc6502) uint8 {
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
  if helpers.FindFunc(lookup[o.opcode].addrmode) == helpers.FindFunc(IMP) {
    o.A = uint8(temp) & 0x00FF
  } else {
    o.write(o.addr_abs, uint8(temp) & 0x00FF)
  }

  return 0
}
func RTI(o *Olc6502) uint8 {
  o.Stkp++
  o.Status = o.read(0x0100 + uint16(o.Stkp))
  o.Status &= ^Flags6502["B"]
  o.Status &= ^Flags6502["U"]

  o.Stkp++
  o.Pc = uint16(o.read(0x0100 + uint16(o.Stkp)))
  o.Stkp++
  o.Pc |= uint16(o.read(0x0100 + uint16(o.Stkp)) << 8)
  return 0
}
func RTS(o *Olc6502) uint8 {
  o.Stkp++
  o.Pc = uint16(o.read(0x0100 + uint16(o.Stkp)))
  o.Stkp++
  o.Pc |= uint16(o.read(0x0100 + uint16(o.Stkp)) << 8)
  o.Pc++
  return 0
}
func SBC(o *Olc6502) uint8 {
  o.fetch()
  value := uint16(o.fetched ^ 0x00FF)
  temp := uint16(uint16(o.A) + value + uint16(o.GetFlag("C")))
  o.SetFlag("C", temp > 0xFF00)
  o.SetFlag("Z", (temp & 0x00FF) == 0)
  o.SetFlag("N", (temp & 0x0080) != 0)
  o.SetFlag("V", (temp ^ uint16(o.A) & (temp ^ value)) & 0x0080 != 0)

  o.A = uint8(temp & 0x00FF)

  return 1
}
func SEC(o *Olc6502) uint8 {
  o.SetFlag("C", true)
  return 0
}
func SED(o *Olc6502) uint8 {
  o.SetFlag("D", true)
  return 0
}
func SEI(o *Olc6502) uint8 {
  o.SetFlag("I", true)
  return 0
}
func STA(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.A)
  return 0
}
func STX(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.X)
  return 0
}
func STY(o *Olc6502) uint8 {
  o.write(o.addr_abs, o.Y)
  return 0
}
func TAX(o *Olc6502) uint8 {
  o.X = o.A
  o.SetFlag("Z", o.X == 0x00)
  o.SetFlag("N", o.X & 0x80 != 0)
  return 0
}
func TAY(o *Olc6502) uint8 {
  o.Y = o.A
  o.SetFlag("Z", o.Y == 0x00)
  o.SetFlag("N", o.Y & 0x80 != 0)
  return 0
}
func TSX(o *Olc6502) uint8 {
  o.X = o.Stkp
  o.SetFlag("Z", o.X == 0x00)
  o.SetFlag("N", o.X & 0x80 != 0)
  return 0
}
func TXA(o *Olc6502) uint8 {
  o.A = o.X
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", o.A & 0x80 != 0)
  return 0
}
func TXS(o *Olc6502) uint8 {
  o.Stkp = o.A
  return 0
}
func TYA(o *Olc6502) uint8 {
  o.A = o.Y
  o.SetFlag("Z", o.A == 0x00)
  o.SetFlag("N", o.A & 0x80 != 0)
  return 0
}
func XXX(o *Olc6502) uint8 {
  return 0
}
