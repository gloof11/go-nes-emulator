package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"reflect"
	"strconv"
  "nes-emulator/bus"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

var (
  nes = bus.NewBus()
  mapAsm map[uint16] string
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

func findFunc(function any) uintptr {
  return reflect.ValueOf(function).Pointer()
}

func hex(number any, width int) string {
  format := fmt.Sprintf("%%0%dX", width)

  return fmt.Sprintf(format, number)
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   14,
	}
}

func init() {
  ss := "A20A8E0000A2038E0100AC0000A900186D010088D0FA8D0200EAEAEA"
  var nOffset uint16 = 0x8000

  for i := 0; i < len(ss); i += 2 {
    hexPair := ss[i : i+2]
    data, err := strconv.ParseUint(hexPair, 16, 8); if err != nil {
      log.Fatalf("Error parsing hex string: %v", err)
    }

    // fmt.Printf("Loading %s into memory at byte %x\n", hexPair, nOffset)
    nes.Ram[nOffset] = uint8(data) // Convert to uint8 and store in memory
    // fmt.Printf("%s loaded into memory as %x\n", hexPair, nes.Ram[nOffset])

    nOffset++
  }

  log.Println("Set the reset vector")
  nes.Ram[0xFFFC] = 0x00
  nes.Ram[0xFFFD] = 0x80

  mapAsm = nes.Cpu.Disassemble(0x0000, 0xFFFF)

  nes.Cpu.Reset()
}

func DrawRam(screen *ebiten.Image, x float64, y float64, nAddr uint16, nRows int, nColumns int) {
  nRamX, nRamY := x, y
  for row := 0; row < nRows; row ++ {
    sOffset := "$" + hex(nAddr, 4) + ":"
    for col := 0; col < nColumns; col++ {
      sOffset += " " + hex(nes.Read(nAddr, true), 2)
      nAddr += 1
    }
    {
      op := &text.DrawOptions{}
      op.GeoM.Translate(nRamX, nRamY)
      op.ColorScale.ScaleWithColor(color.White)
      op.LineSpacing = mplusNormalFace.Size
      text.Draw(screen, sOffset, mplusNormalFace, op)
    }
    nRamY += 15
  }
}

func DrawCpu(screen *ebiten.Image, x float64, y float64) {
  drawnString := "STATUS: "
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y)
    op.ColorScale.ScaleWithColor(color.White)
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "N"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 64, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["N"] == nes.Cpu.FLAGS6502["N"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "V"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 80, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["V"] == nes.Cpu.FLAGS6502["V"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "-"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 96, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["U"] == nes.Cpu.FLAGS6502["U"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "B"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 112, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["B"] == nes.Cpu.FLAGS6502["B"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "D"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 128, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["D"] == nes.Cpu.FLAGS6502["D"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "I"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 144, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["I"] == nes.Cpu.FLAGS6502["I"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "Z"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 160, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["Z"] == nes.Cpu.FLAGS6502["Z"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "C"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x + 178, y)
    if nes.Cpu.Status & nes.Cpu.FLAGS6502["C"] == nes.Cpu.FLAGS6502["C"] {
      op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
    } else {
      op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 0})
    }
    op.LineSpacing = mplusNormalFace.Size
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "PC: $" + hex(nes.Cpu.Pc, 4)
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y + 10 * 2)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "A: $" + hex(nes.Cpu.A, 2)
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y + 20 * 2)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "X: $" + hex(nes.Cpu.X, 2)
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y + 30 * 2)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "Y: $" + hex(nes.Cpu.Y, 2)
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y + 40 * 2)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
  drawnString = "Stack P: $" + hex(nes.Cpu.Stkp, 4)
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(x, y + 50 * 2)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
}

func DrawCode(screen *ebiten.Image, x float64, y float64, nLines int) {
  var nLineY int = (nLines >> 1) * 10 + int(y)

//  for i := nLines; i >= 0; i-- {
//    line, ok := mapAsm[nes.Cpu.Pc-uint16(i)]
//    if ok {
//      {
//        op := &text.DrawOptions{}
//        op.GeoM.Translate(x, float64(nLineY))
//        op.LineSpacing = mplusNormalFace.Size
//        op.ColorScale.ScaleWithColor(color.White)
//        text.Draw(screen, line, mplusNormalFace, op)
//        nLineY += 15
//      }
//    }
//  }

  for i := 0; i < nLines; i++ {
    line, ok := mapAsm[nes.Cpu.Pc+uint16(i)]
    if ok {
      if (nes.Cpu.Pc+uint16(i)) != nes.Cpu.Pc {
        {
          op := &text.DrawOptions{}
          op.GeoM.Translate(x, float64(nLineY))
          op.LineSpacing = mplusNormalFace.Size
          op.ColorScale.ScaleWithColor(color.White)
          text.Draw(screen, line, mplusNormalFace, op)
          nLineY += 15
        }
      } else {
        {
          op := &text.DrawOptions{}
          op.GeoM.Translate(x, float64(nLineY))
          op.LineSpacing = mplusNormalFace.Size
          op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 0})
          text.Draw(screen, line, mplusNormalFace, op)
          nLineY += 15
        }
      }
    }
  }
}

type Game struct{}

func (g *Game) Update() error {
  if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
    for {
      nes.Cpu.Clock()
      //fmt.Println(hex(nes.Cpu.Pc, 4))
      if nes.Cpu.Complete() {
        break
      }
    }
  }

  if inpututil.IsKeyJustPressed(ebiten.KeyR) {
    nes.Cpu.Reset()
  }
  if inpututil.IsKeyJustPressed(ebiten.KeyI) {
    nes.Cpu.Irq()
  }
  if inpututil.IsKeyJustPressed(ebiten.KeyN) {
    nes.Cpu.Nmi()
  }

  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  DrawRam(screen, 2, 2, 0x0000, 16, 16)
  DrawRam(screen, 2, 280, 0x8000, 16, 16)
  DrawCpu(screen, windowWidth/1.4, 2)
  DrawCode(screen, windowWidth/1.4, windowHeight/32, 26)

  drawnString := "SPACE = Step Instruction, R = RESET, I = IRQ, N = NMI"
  {
    op := &text.DrawOptions{}
    op.GeoM.Translate(windowWidth / 4, windowHeight - 40)
    op.LineSpacing = mplusNormalFace.Size
    op.ColorScale.ScaleWithColor(color.White)
    text.Draw(screen, drawnString, mplusNormalFace, op)
  }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return windowWidth, windowHeight
}

func main() {
  ebiten.SetWindowSize(windowWidth, windowHeight)
  ebiten.SetWindowTitle("6502 Demonstration")

  nes.Cpu.Cycles = 0x00

  if err := ebiten.RunGame(&Game{}); err != nil { 
    log.Fatal(err)
  }
}
