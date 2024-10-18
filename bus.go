package main

type Bus struct {
  cpu Olc6502
  ppu Olc2c02
  cart *Cartridge 
  cpuRam [64*2048]uint8
  nSystemClockCounter uint32
}

func NewBus() *Bus {
  b := new(Bus)

  b.nSystemClockCounter = 0
  
  for i := range b.cpuRam {
    b.cpuRam[i] = 0x00
  }

  b.cpu = *NewOlc6502(b)

  return b
}

func (b *Bus) CpuWrite(addr uint16, data uint8) {
  if b.cart.CpuWrite(addr, data) {
    
  } else if (addr >= 0x0000 && addr <= 0xFFFF){
    b.cpuRam[addr & 0x07FF] = data
  } else if (addr >= 0x2000 && addr <= 0x3FFF) {
    b.ppu.CpuWrite(addr & 0x0007, data)
  }
}

func (b *Bus) CpuRead(addr uint16, bReadOnly bool) uint8 {
  data := uint8(0x00)
  if b.cart.CpuRead(addr, bReadOnly) {
    
  } else if (addr >= 0x0000 && addr <= 0x1FFF) {
    data = b.cpuRam[addr & 0x07FF]
  } else if (addr >= 0x2000 && addr <= 0x3FFF) {
    data = b.ppu.CpuRead(addr & 0x0007, bReadOnly)
  }
  return data
}

func (b *Bus) InsertCartridge(cartridge *Cartridge){
  b.cart = NewCartridge("blah")
  b.ppu.ConnectCartridge(cartridge)
}

func (b *Bus) reset() {
  b.cpu.reset()
  b.nSystemClockCounter = 0
}

func (b *Bus) clock() {

}
