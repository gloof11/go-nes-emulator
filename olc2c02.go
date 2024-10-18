package main

type Olc2c02 struct {
  cart *Cartridge
  tblName [2][1024]uint8
  tblPalette [32]uint8
  tblPattern [2][4096]uint8 // Used for custom mappers
}

func (ppu *Olc2c02) PpuWrite(addr uint16, data uint8) {
  addr &= 0x3FFF
}

func (ppu *Olc2c02) PpuRead(addr uint16, bReadOnly bool) uint8 {
  var data uint8 = 0x00
  addr &= 0x3FFF

  return data
}

func (ppu *Olc2c02) CpuWrite(addr uint16, data uint8) {
  switch addr {
  case 0x000:
    break
  case 0x001:
    break
  case 0x002:
    break
  case 0x003:
    break
  case 0x004:
    break
  case 0x005:
    break
  case 0x006:
    break
  case 0x007:
    break
  }
}

func (ppu *Olc2c02) CpuRead(addr uint16, bReadOnly bool) uint8 {
  var data uint8 = 0x0

  switch addr {
  case 0x000:
    break
  case 0x001:
    break
  case 0x002:
    break
  case 0x003:
    break
  case 0x004:
    break
  case 0x005:
    break
  case 0x006:
    break
  case 0x007:
    break
  }

  return data
}

func (ppu *Olc2c02) ConnectCartridge(cartridge *Cartridge) {
  ppu.cart = cartridge
}

func (ppu *Olc2c02) clock() {

}
