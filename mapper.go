package main

type Mapper struct {
  nPRGBanks uint8
  nCHRBanks uint8
}

type MapperInterface interface {
  CpuMapRead(addr uint16, mapped_addr uint32) bool
  CpuMapWrite(addr uint16, mapped_addr uint32) bool
  PpuMapRead(addr uint16, mapped_addr uint32) bool
  PpuMapWrite(addr uint16, mapped_addr uint32) bool
}

// func NewMapper(prgBanks uint8, chrBanks uint8) *Mapper {
//   mapper := new(Mapper)
//
//   return mapper
// }
//
// func (mapper *Mapper) CpuMapRead(addr uint16, mapped_addr uint32) bool {
//   return false
// }
// func (mapper *Mapper) CpuMapWrite(addr uint16, mapped_addr uint32) bool {
//   return false
// }
// func (mapper *Mapper) PpuMapRead(addr uint16, mapped_addr uint32) bool {
//   return false
// }
func (mapper *Mapper) PpuMapWrite(addr uint16, mapped_addr uint32) bool {
  return false
}
