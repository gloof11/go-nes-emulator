package main

import (
  "testing"
)

func TestIMP(t *testing.T) {
  nes.cpu.a = 0x42

  nes.cpu.IMP(&nes.cpu)

  if nes.cpu.fetched != nes.cpu.a {
    t.Error("Expected fetched value", nes.cpu.a, "got", nes.cpu.fetched)
  }
}

func TestIMM(t *testing.T) {
  nes.cpu.pc = 0x01
  originalPC := uint16(0x01)
  
  nes.cpu.IMM(&nes.cpu)

  // check if o.addr_abs is o.pc-1
  if nes.cpu.addr_abs != nes.cpu.pc-1 {
    t.Errorf("Expected %X, got %X", nes.cpu.pc-1, nes.cpu.addr_abs)
  }

  // check if o.pc++
  if nes.cpu.pc != (originalPC+1) {
    t.Errorf("Expected %X, got %X", originalPC+1, nes.cpu.pc)
  }
}

func TestZP0(t *testing.T) {
  // test that addr_abs reads from what's at the pc

  // test that the pc gets incremented
}

// func (nes.cpu *Olc6502) ZP0(o *Olc6502) uint8 {
//   o.addr_abs = uint16(o.read(o.pc))
//   o.pc++
//   o.addr_abs &= 0x00FF
//   return 0
// }
// func (nes.cpu *Olc6502) ZPX(o *Olc6502) uint8 {
//   o.addr_abs = uint16(o.read(o.pc) + o.x)
//   o.pc++
//   o.addr_abs &= 0x00FF
//   return 0
// }
// func (nes.cpu *Olc6502) ZPY(o *Olc6502) uint8 {
//   o.addr_abs = uint16(o.read(o.pc) + o.y)
//   o.pc++
//   o.addr_abs &= 0x00FF
//   return 0
// }
// func (nes.cpu *Olc6502) REL(o *Olc6502) uint8 {
//   o.addr_rel = uint16(o.read(o.pc))
//   o.pc++
//   if ((o.addr_rel & 0x80) != 0) {  
//     o.addr_rel |= 0xFF00
//   }
//   return 0
// }
// func (nes.cpu *Olc6502) ABS(o *Olc6502) uint8 {
//   lo := uint16(o.read(o.pc))
//   o.pc++
//   hi := uint16(o.read(o.pc))
//   o.pc++
//   o.addr_abs = (hi << 8) | lo
//   return 0
// }
// func (nes.cpu *Olc6502) ABX(o *Olc6502) uint8 {
//   lo := uint16(o.read(o.pc))
//   o.pc++
//   hi := uint16(o.read(o.pc))
//   o.pc++
//   o.addr_abs = (hi << 8) | lo
//   o.addr_abs += uint16(o.x)
//
//   if ((o.addr_abs & 0xFF00) != (hi << 8)) {
//     return 1
//   } else {
//     return 0
//   }
// }
// func (nes.cpu *Olc6502) ABY(o *Olc6502) uint8 {
//   lo := uint16(o.read(o.pc))
//   o.pc++
//   hi := uint16(o.read(o.pc))
//   o.pc++
//   o.addr_abs = (hi << 8) | lo
//   o.addr_abs += uint16(o.y)
//
//   if ((o.addr_abs & 0xFF00) != (hi << 8)) {
//     return 1
//   } else {
//     return 0
//   }
// }
// func (nes.cpu *Olc6502) IND(o *Olc6502) uint8 {
//   ptr_lo := uint16(o.read(o.pc))
//   o.pc++
//   ptr_hi := uint16(o.read(o.pc))
//   o.pc++
//
//   ptr := uint16(ptr_hi << 8) | ptr_lo
//
//   if (ptr_lo == 0x00FF) {
//     o.addr_abs = uint16((o.read(ptr & 0xFF00) << 8) | o.read(ptr + 0))
//   } else {
//     o.addr_abs = uint16((o.read(ptr + 1) << 8) | o.read(ptr + 0))
//   }
//   return 0
// }
// func (nes.cpu *Olc6502) IZX(o *Olc6502) uint8 {
//   t := uint16(o.read(o.pc))
//   o.pc++
//
//   lo := uint16(o.read((t + uint16(o.x)) & 0x00FF))
//   hi := uint16(o.read((t + uint16(o.x + 1)) & 0x00FF))
//
//   o.addr_abs = (hi << 8) | lo
//   return 0
// }
// func (nes.cpu *Olc6502) IZY(o *Olc6502) uint8 {
//   t := uint16(o.read(o.pc))
//   o.pc++
//
//   lo := uint16(o.read(t & 0x00FF))
//   hi := uint16(o.read((t+1) & 0x00FF))
//
//   o.addr_abs = (hi << 8) | lo
//   o.addr_abs += uint16(o.y)
//
//   if ((o.addr_abs & 0xFF00) != (hi << 8)) {
//     return 1
//   } else {
//     return 0
//   }
// }
