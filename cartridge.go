package main

import (
  "os"
  "encoding/binary"
)

type Cartridge struct {
  vPRGMemory []uint8
  vCHRMemory []uint8
  nMapperID uint8
  nPRGBanks uint8
  nCHRBanks uint8
  pMapper MapperInterface
}

func NewCartridge(filename string) *Cartridge {
  cart := new(Cartridge)

  cart.nMapperID = 0
  cart.nPRGBanks = 0
  cart.nCHRBanks = 0

  type Header struct {
    name [4]string
    prg_rom_chunks uint8
    chr_rom_chunks uint8
    mapper1 uint8
    mapper2 uint8
    prg_ram_size uint8
    tv_system1 uint8
    unused [5]string
  }

  header := Header{}

  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  err = binary.Read(f, binary.LittleEndian, header)
  if err != nil {
    panic(err)
  }

  if (header.mapper1 & 0x04) != 0 {
    f.Seek(512, 1)
  }

  cart.nMapperID = ((header.mapper2 >> 4) << 4 | (header.mapper1 >> 4) )

  var nFileType uint8 = 1
  
  if nFileType == 0 {

  }
  if nFileType == 1 {
    cart.nPRGBanks = header.prg_rom_chunks
    cart.vPRGMemory = make([]uint8, 0, int(cart.nPRGBanks) * 16384)
    f.Read(cart.vPRGMemory)

    cart.nCHRBanks = header.chr_rom_chunks
    cart.vCHRMemory = make([]uint8, 0, int(cart.nCHRBanks) * 8192)
    f.Read(cart.vCHRMemory)
  }

  switch cart.nMapperID {
  case 0:
    cart.pMapper = NewMapper0(cart.nPRGBanks, cart.nCHRBanks)
  }

  return cart
}

func (cart *Cartridge) CpuRead(addr uint16, data uint8) bool {
  // var mapped_addr uint32 = 0
  
  // if cart.pMapper.CpuMapRead(addr, mapped_addr) {
  //   data = cart.vPRGMemory[mapped_addr]
  //   return true
  // } else {
  //   return false
  // }

  return false
}

func (cart *Cartridge) CpuWrite(addr uint16, data uint8) bool {
  // var mapped_addr uint32 = 0
  //
  // if cart.pMapper.CpuMapWrite(addr, mapped_addr) {
  //   data = cart.vPRGMemory[mapped_addr]
  //   return true
  // } else {
  //   return false
  // }

  return false
}

func (cart *Cartridge) PpuRead(addr uint16, data uint8) bool {
  // var mapped_addr uint32 = 0
  //
  // if cart.pMapper.PpuMapRead(addr, mapped_addr) {
  //   data = cart.vCHRMemory[mapped_addr]
  //   return true
  // } else {
  //   return false
  // }
  
  return false
}

func (cart *Cartridge) PpuWrite(addr uint16, data uint8) bool {
  // var mapped_addr uint32 = 0
  //
  // if cart.pMapper.PpuMapWrite(addr, mapped_addr) {
  //   data = cart.vCHRMemory[mapped_addr]
  //   return true
  // } else {
  //   return false
  // }

  return false
}

