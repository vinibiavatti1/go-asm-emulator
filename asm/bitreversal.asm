            ; --- Check if a Number is a Power of Two ---
            ; Number: 16 (0b00010000)
            ; Bitwise trick: N & (N - 1) should equal 0 for powers of 2.
            ; Result: 1 (True) stored into Memory[1]

            mov rax, 16          ; Value to test (16)
            mov rbx, rax         ; Copy to rbx
            dec rbx              ; rbx = 15 (0b00001111)

            and rbx, rax         ; rbx = 16 & 15

            cmp rbx, 0
            je is_power_of_two

            ; Not a power of two
            mov rax, 0
            jmp store_result

is_power_of_two:
            mov rax, 1

store_result:
            ; Save result (1 = True, 0 = False) into Memory[1]
            mov rdi, 1
            mov [rdi], rax