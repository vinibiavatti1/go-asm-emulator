            ; --- Calculate the 6th Fibonacci term (F(6) = 8) ---
            ; F(0) = 0, F(1) = 1, F(2) = 1, F(3) = 2, F(5) = 5, F(6) = 8
            
            mov rax, 0           ; rax = F(n-2) (starts at 0)
            mov rbx, 1           ; rbx = F(n-1) (starts at 1)
            mov rcx, 6           ; rcx = loop counter (target is 6th term)
            
fib_loop:
            cmp rcx, 0           ; Check if counter has reached 0
            je end               ; If zero, jump out of the loop
            
            ; Save current state to the stack for safety
            push rax             ; Save F(n-2)
            push rbx             ; Save F(n-1)
            
            ; Calculate next term: rdx = rax + rbx
            mov rdx, rax         ; rdx = rax
            add rdx, rbx         ; rdx = rax + rbx (new term)
            
            ; Shift values for the next iteration
            ; F(n-2) becomes old F(n-1)
            ; F(n-1) becomes the newly calculated term (rdx)
            mov rax, rbx         ; rax gets old F(n-1)
            mov rbx, rdx         ; rbx gets the new term
            
            ; Clean up the stack for this iteration (2 pushes, 2 pops)
            pop rdx              ; Discard temporary stack garbage
            pop rdx              ; Keeping rsp perfectly synchronized
            
            dec rcx              ; Decrement loop counter
            jmp fib_loop         ; Jump back to the start of the loop

end:
            ; Save the final result (in rax) directly into memory slot 1
            mov rdi, 0           ; rdi points to memory index 0
            mov [rdi], rax       ; Store F(6) into Memory[0] (Alternatively: mov [0x0], rax depending on your parser syntax)