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
git clone [https://github.com/vinibiavatti1/go-asm-emulator.git](https://github.com/vinibiavatti1/go-asm-emulator.git)
cd go-asm-emulator
./go-asm-emulator.exe -src ./asm/fibonacci.asm
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
          mov rdi, 1
          mov [rdi], rax       ; Stores the final result into Memory[1]
```

## License

This project is open-source under the MIT License.
