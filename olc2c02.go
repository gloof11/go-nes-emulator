package main

type Olc2c02 struct {

}

func (b *Olc2c02) PpuWrite(addr uint16, data uint8) {
}

func (b *Olc2c02) PpuRead(addr uint16, bReadOnly bool) uint8 {
}

func (b *Olc2c02) CpuWrite(addr uint16, data uint8) {
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

func (b *Olc2c02) CpuRead(addr uint16, bReadOnly bool) uint8 {
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
