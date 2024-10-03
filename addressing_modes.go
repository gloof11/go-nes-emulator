package main

// Addressing Modes
func (o *Olc6502) IMP() uint8 { o.fetched = o.a; return 0 }
func (o *Olc6502) IMM() uint8 {
  pc := o.pc
  pc++
  o.addr_abs = pc
  return 0 
}
func (o *Olc6502) ZP0() uint8 {
  o.addr_abs = uint16(o.read(o.pc))
  o.pc++
  o.addr_abs &= 0x00FF
  return 0
}
func (o *Olc6502) ZPX() uint8 {
  o.addr_abs = uint16(o.read(o.pc) + o.x)
  o.pc++
  o.addr_abs &= 0x00FF
  return 0
}
func (o *Olc6502) ZPY() uint8 {
  o.addr_abs = uint16(o.read(o.pc) + o.y)
  o.pc++
  o.addr_abs &= 0x00FF
  return 0
}
func (o *Olc6502) REL() uint8 {
  o.addr_rel = uint16(o.read(o.pc))
  o.pc++
  if ((o.addr_rel & 0x80) != 0) {  
    o.addr_rel |= 0xFF00
  }
  return 0
}
func (o *Olc6502) ABS() uint8 {
  lo := uint16(o.read(o.pc))
  o.pc++
  hi := uint16(o.read(o.pc))
  o.pc++
  o.addr_abs = (hi << 8) | lo
  return 0
}
func (o *Olc6502) ABX() uint8 {
  lo := uint16(o.read(o.pc))
  o.pc++
  hi := uint16(o.read(o.pc))
  o.pc++
  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.x)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
func (o *Olc6502) ABY() uint8 {
  lo := uint16(o.read(o.pc))
  o.pc++
  hi := uint16(o.read(o.pc))
  o.pc++
  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.y)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
func (o *Olc6502) IND() uint8 {
  ptr_lo := uint16(o.read(o.pc))
  o.pc++
  ptr_hi := uint16(o.read(o.pc))
  o.pc++
  
  ptr := uint16(ptr_hi << 8) | ptr_lo

  if (ptr_lo == 0x00FF) {
    o.addr_abs = uint16((o.read(ptr & 0xFF00) << 8) | o.read(ptr + 0))
  } else {
    o.addr_abs = uint16((o.read(ptr + 1) << 8) | o.read(ptr + 0))
  }
  return 0
}
func (o *Olc6502) IZX() uint8 {
  t := uint16(o.read(o.pc))
  o.pc++

  lo := uint16(o.read((t + uint16(o.x)) & 0x00FF))
  hi := uint16(o.read((t + uint16(o.x + 1)) & 0x00FF))

  o.addr_abs = (hi << 8) | lo
  return 0
}
func (o *Olc6502) IZY() uint8 {
  t := uint16(o.read(o.pc))
  o.pc++

  lo := uint16(o.read(t & 0x00FF))
  hi := uint16(o.read((t+1) & 0x00FF))

  o.addr_abs = (hi << 8) | lo
  o.addr_abs += uint16(o.y)

  if ((o.addr_abs & 0xFF00) != (hi << 8)) {
    return 1
  } else {
    return 0
  }
}
