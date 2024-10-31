package bus

import (
  "nes-emulator/olc6502"
)

type Bus struct {
  Cpu olc6502.Olc6502
  Ram [64*1024]uint8
}

func NewBus() *Bus {
  b := new(Bus)
  
  for i := range b.Ram {
    b.Ram[i] = 0x00
  }

  b.Cpu = *olc6502.NewOlc6502(&b.Ram)

  return b
}

func (b *Bus) Write(addr uint16, data uint8) {
  if (addr >= 0x0000 && addr <= 0xFFFF){
    b.Ram[addr] = data
  }
}

func (b *Bus) Read(addr uint16, bReadOnly bool) uint8 {
  if (addr >= 0x0000 && addr <= 0xFFFF) {
    return b.Ram[addr]
  }
  return 0x00
}
