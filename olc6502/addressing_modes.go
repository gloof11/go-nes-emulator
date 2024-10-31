package olc6502

// Addressing Modes
func IMP(o *Olc6502) uint8 { o.fetched = o.A; return 0 }
func IMM(o *Olc6502) uint8 {
  o.addr_abs = o.Pc
  o.Pc++
  return 0
}
func ZP0(o *Olc6502) uint8 {
  o.addr_abs = uint16(o.read(o.Pc))
  o.Pc++
  o.addr_abs &= 0x00FF
  return 0
}
func ZPX(o *Olc6502) uint8 {
  o.addr_abs = uint16(o.read(o.Pc) + o.X)
  o.Pc++
  o.addr_abs &= 0x00FF
  return 0
}
func ZPY(o *Olc6502) uint8 {
  o.addr_abs = uint16(o.read(o.Pc) + o.Y)
  o.Pc++
  o.addr_abs &= 0x00FF
  return 0
}
func REL(o *Olc6502) uint8 {
  o.addr_rel = uint16(o.read(o.Pc))
  o.Pc++
  if ((o.addr_rel & 0x80) != 0) {  
    o.addr_rel |= 0xFF00
  }
  return 0
}
func ABS(o *Olc6502) uint8 {
  lo := uint16(o.read(o.Pc))
  o.Pc++
  hi := uint16(o.read(o.Pc))
  o.Pc++
  o.addr_abs = (hi << 8) | lo
  return 0
}
func ABX(o *Olc6502) uint8 {
  lo := uint16(o.read(o.Pc))
  o.Pc++
  hi := uint16(o.read(o.Pc))
  o.Pc++
  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.X)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
func ABY(o *Olc6502) uint8 {
  lo := uint16(o.read(o.Pc))
  o.Pc++
  hi := uint16(o.read(o.Pc))
  o.Pc++
  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.Y)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
func IND(o *Olc6502) uint8 {
  ptr_lo := uint16(o.read(o.Pc))
  o.Pc++
  ptr_hi := uint16(o.read(o.Pc))
  o.Pc++
  
  ptr := uint16(ptr_hi << 8) | ptr_lo

  if (ptr_lo == 0x00FF) {
    o.addr_abs = uint16((o.read(ptr & 0xFF00) << 8) | o.read(ptr + 0))
  } else {
    o.addr_abs = uint16((o.read(ptr + 1) << 8) | o.read(ptr + 0))
  }
  return 0
}
func IZX(o *Olc6502) uint8 {
  t := uint16(o.read(o.Pc))
  o.Pc++

  lo := uint16(o.read((t + uint16(o.X)) & 0x00FF))
  hi := uint16(o.read((t + uint16(o.X + 1)) & 0x00FF))

  o.addr_abs = (hi << 8) | lo
  return 0
}
func IZY(o *Olc6502) uint8 {
  t := uint16(o.read(o.Pc))
  o.Pc++

  lo := uint16(o.read(t & 0x00FF))
  hi := uint16(o.read((t+1) & 0x00FF))

  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.Y)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
