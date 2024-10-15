package main

type Bus struct {
  cpu Olc6502
  ppu Olc2c02
  cpuRam [64*2048]uint8
}

func NewBus() *Bus {
  b := new(Bus)
  
  for i := range b.cpuRam {
    b.cpuRam[i] = 0x00
  }

  b.cpu = *NewOlc6502(b)

  return b
}

func (b *Bus) CpuWrite(addr uint16, data uint8) {
  if (addr >= 0x0000 && addr <= 0xFFFF){
    b.cpuRam[addr & 0x07FF] = data
  }
  if (addr >= 0x2000 && addr <= 0x3FFF) {
    b.ppu.CpuWrite(addr & 0x0007, data)
  }
}

func (b *Bus) CpuRead(addr uint16, bReadOnly bool) uint8 {
  data := uint8(0x00)
  if (addr >= 0x0000 && addr <= 0x1FFF) {
    data = b.cpuRam[addr & 0x07FF]
  }
  if (addr >= 0x2000 && addr <= 0x3FFF) {
    data = b.ppu.CpuRead(addr & 0x0007, bReadOnly)
  }
  return data
}
