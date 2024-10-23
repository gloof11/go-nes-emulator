package main

type Mapper0 struct {
  Mapper
}

func NewMapper0(prgBanks uint8, chrBanks uint8) *Mapper0 {
  mapper0 := new(Mapper0)

  return mapper0
}

func (mapper *Mapper0) CpuMapRead(addr uint16, mapped_addr uint32) bool {
  if addr >= 0x8000 && addr <= 0xFFFF {
    // mapped_addr := func() uint32 {
    //   if mapper.nPRGBanks > 1 {
    //     return uint32(addr) & 0x7FFF
    //   } else {
    //     return uint32(addr) & 0x3FFF
    //   }
    // }
    return true
  }
  return false
}
func (mapper *Mapper0) CpuMapWrite(addr uint16, mapped_addr uint32) bool {
  if addr >= 0x8000 && addr <= 0xFFFF {
    return true
  }
  return false
}
func (mapper *Mapper0) PpuMapRead(addr uint16, mapped_addr uint32) bool {
  if addr >= 0x0000 && addr <= 0x1FFF {
    mapped_addr = uint32(addr)
    return true
  }
  return false
}
// func (mapper *Mapper0) PpuMapWrite(addr uint16, mapped_addr uint32) bool {
//   // if addr >= 0x0000 && addr <= 0x1FFF {
//   //   return true
//   // }
//   return false
// }
