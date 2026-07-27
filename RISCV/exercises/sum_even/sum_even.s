.globl _start

.data
array_data: .word 1, 2, 3, 4, 5, 6, 7, 8, 9, 10;

.text 
_start:
    la a0, array_data
    li a1, 10
    jal ra, sum_even_numbers
    ebreak

sum_even_numbers:
    mv t0, a0
    add a0, zero, zero
loop:
    beq a1, zero, return

    lw t1, 0(t0)
    andi t2, t1, 1
    bnez t2, advance_counter
    add a0, a0, t1
advance_counter:
    addi a1, a1, -1
    addi t0, t0, 4
    j loop
return:
    ret
