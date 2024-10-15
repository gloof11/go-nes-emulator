package main

type Cartridge struct {

}

func (b *Cartridge) PpuWrite(addr uint16, data uint8) {
  if (addr >= 0x0000 && addr <= 0xFFFF){
    b.cpuRam[addr & 0x07FF] = data
  }
}

func (b *Cartridge) PpuRead(addr uint16, bReadOnly bool) uint8 {
  data := uint8(0x00)
  if (addr >= 0x0000 && addr <= 0x1FFF) {
    data = b.cpuRam[addr & 0x07FF]
  }
  return data
}

func (b *Cartridge) CpuWrite(addr uint16, data uint8) {
  if (addr >= 0x0000 && addr <= 0xFFFF){
    b.cpuRam[addr & 0x07FF] = data
  }
}

func (b *Cartridge) CpuRead(addr uint16, bReadOnly bool) uint8 {
  data := uint8(0x00)
  if (addr >= 0x0000 && addr <= 0x1FFF) {
    data = b.cpuRam[addr & 0x07FF]
  }
  return data
}
