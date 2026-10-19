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
Clone the repository, compile the binary, and run your Assembly files using the `-src` flag:

```bash
git clone https://github.com/vinibiavatti1/go-asm-emulator.git
cd go-asm-emulator

# Run with default stack (8) and memory (8) sizes
./go-asm-emulator.exe -src ./asm/fibonacci.asm

# Run with custom stack and memory allocations
./go-asm-emulator.exe -src ./asm/fibonacci.asm -memory 16 -stack 16
```

## Example Usage

Here is a simple Fibonacci sequence runner ($F(6) = 8$) that stores the result in memory:

```asm
          mov rax, 0           ; rax = F(n-2)
          mov rbx, 1           ; rbx = F(n-1)
          mov rcx, 6           ; rcx = loop counter
fib_loop:
          cmp rcx, 0
          je end
          push rax
          push rbx
          mov rdx, rax
          add rdx, rbx
          mov rax, rbx
          mov rbx, rdx         
          pop rdx
          pop rdx
          dec rcx
          jmp fib_loop
end:
          mov rdi, 0
          mov [rdi], rax       ; Stores the final result into Memory[0]
```

Result:

```text
Registers:
rax.......: 0x00000008 (8)
rbx.......: 0x0000000D (13)
rcx.......: 0x00000000 (0)
rdx.......: 0x00000005 (5)
rsi.......: 0x00000000 (0)
rdi.......: 0x00000000 (0)
rsp.......: 0x00000000 (0)
rbp.......: 0x00000000 (0)
eflags....: 0x00000001 (1)
Memory (8):
0x00000000: 0x00000008 (8)
0x00000001: 0x00000000 (0)
0x00000002: 0x00000000 (0)
0x00000003: 0x00000000 (0)
0x00000004: 0x00000000 (0)
0x00000005: 0x00000000 (0)
0x00000006: 0x00000000 (0)
0x00000007: 0x00000000 (0)
Stack  (8):
0x00000000: 0x00000000 (0)
0x00000001: 0x00000000 (0)
0x00000002: 0x00000000 (0)
0x00000003: 0x00000000 (0)
0x00000004: 0x00000000 (0)
0x00000005: 0x00000000 (0)
0x00000006: 0x00000000 (0)
0x00000007: 0x00000000 (0)
```

## License

This project is open-source under the MIT License.
