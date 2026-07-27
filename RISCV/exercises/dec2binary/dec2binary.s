.globl _start

.bss   
.align 2
binary_array:
    .space 33

.text
.align 2

_start:
    li a0, 98452
    la a1, binary_array
    jal ra, decimal_to_binary
    ebreak


decimal_to_binary:
    mv t0, a1
    li t1, 32

extract_bit:
    srli t2, a0, 31
    sb t2, 0(t0)
    
    slli a0, a0, 1
    addi t0, t0, 1
    addi t1, t1, -1
    
    bnez t1, extract_bit

append_terminator:
    li t2, 2
    sb t2, 0(t0)
    
    mv a0, a1
    ret
    