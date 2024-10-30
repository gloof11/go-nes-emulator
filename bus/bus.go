package bus

import (
  "nes-emulator/olc6502"
)

type Bus struct {
  cpu olc6502.Olc6502
  ram [64*1024]uint8
}

func NewBus() *Bus {
  b := new(Bus)
  
  for i := range b.ram {
    b.ram[i] = 0x00
  }

  b.cpu = *olc6502.NewOlc6502(&b.ram)

  return b
}

func (b *Bus) write(addr uint16, data uint8) {
  if (addr >= 0x0000 && addr <= 0xFFFF){
    b.ram[addr] = data
  }
}

func (b *Bus) read(addr uint16, bReadOnly bool) uint8 {
  if (addr >= 0x0000 && addr <= 0xFFFF) {
    return b.ram[addr]
  }
  return 0x00
}
