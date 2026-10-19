            ; --- Calculate Factorial of 5 (5! = 120) ---
            ; 5! = 5 * 4 * 3 * 2 * 1 = 120 (0x78 in hex)

            mov rax, 1           ; rax = accumulator (1!)
            mov rcx, 5           ; rcx = number to calculate factorial for

fact_loop:
            cmp rcx, 1           ; Stop when counter reaches 1
            je fim

            ; Multiply rax by rcx using repeated addition
            push rcx             ; Save outer loop counter
            mov rbx, rax         ; rbx = base value to add
            mov rdx, rcx         ; rdx = multiplier counter
            dec rdx              ; We already have 1x in rax, need (rdx - 1) additions

mult_loop:
            cmp rdx, 0
            je mult_done
            add rax, rbx
            dec rdx
            jmp mult_loop

mult_done:
            pop rcx              ; Restore outer loop counter
            dec rcx              ; Move to next factor (e.g., 5 -> 4 -> 3)
            jmp fact_loop

fim:
            ; Save result (120 / 0x78) into Memory[1]
            mov rdi, 1
            mov [rdi], rax