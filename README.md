# GoAsmEmulator

A lightweight, modular x86-inspired Assembly emulator built from scratch in Go, featuring a custom instruction parser, dynamic stack management via `rsp`, independent RAM, and general-purpose registers.

## Features

- **Registers**: General-purpose registers (`rax`, `rbx`, `rcx`, `rdx`, `rsi`, `rdi`, `rsp`, `rbp`) and `eflags`.
- **Stack & RAM**: Dedicated LIFO stack managed via `rsp` and isolated memory slots.
- **Control Flow**: Supports conditional branching (`cmp`, `je`, `jne`) and loops.
- **Instruction Parsing**: Line-by-line synchronization and label handling.

## Getting Started

### Prerequisites
- Go 1.20+ installed on your machine.

### Running the Emulator
Clone the repository and run the main program:

```bash
git clone [https://github.com/YOUR_USERNAME/go-asm-emulator.git](https://github.com/YOUR_USERNAME/go-asm-emulator.git)
cd go-asm-emulator
go run main.go
