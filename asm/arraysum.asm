            ; --- Sum Array Elements in Memory ---
            ; Array: Memory[2..5] = [10, 20, 30, 40]
            ; Expected Sum: 100 (0x64 in hex) stored into Memory[1]

            ; --- Setup Array in Memory ---
            mov rdi, 2
            mov rax, 10
            mov [rdi], rax

            mov rdi, 3
            mov rax, 20
            mov [rdi], rax

            mov rdi, 4
            mov rax, 30
            mov [rdi], rax

            mov rdi, 5
            mov rax, 40
            mov [rdi], rax

            ; --- Calculate Sum ---
            mov rax, 0           ; rax = total accumulator
            mov rcx, 4           ; rcx = array length
            mov rsi, 2           ; rsi = pointer to start of array (Memory[2])

sum_loop:
            cmp rcx, 0
            je sum_done

            mov rbx, [rsi]       ; Load current element from Memory[rsi]
            add rax, rbx         ; Accumulate sum
            inc rsi              ; Move pointer to next memory slot
            dec rcx              ; Decrement loop counter
            jmp sum_loop

sum_done:
            ; Save final total (100 / 0x64) into Memory[1]
            mov rdi, 1
            mov [rdi], rax