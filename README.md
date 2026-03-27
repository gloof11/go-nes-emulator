# GoNES Emulator

## Description

A NES emulator built in Go, inspired by javidx9's NES emulator. It focuses on providing an educational view of the 6502 CPU and NES hardware.

## Features

*   Accurate 6502 CPU emulation (based on javidx9's work).
*   Basic memory and CPU register visualization using Ebiten.
*   Currently, it does not load external ROMs; it executes a hardcoded program.

## Getting Started

### Prerequisites

*   Go 1.22.4 or later

### Running the Emulator

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd go-nes-emulator
    ```
2.  Run the emulator:
    ```bash
    go run .
    ```

### Interaction

*   Press **SPACE** to step through instructions.
*   Press **R** to reset the emulator.
*   Press **I** to trigger an IRQ.
*   Press **N** to trigger an NMI.

## Development

### Building

To build the emulator from source, run:

```bash
go build
```
