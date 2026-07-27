#ELF

.globl _start
.org 0x400

.data
binary_array: .byte 1, 0, 0, 1, 0, 1, 2;  # Example
.org 0x200

.text
_start:
    la a0, binary_array
    jal ra, binary_to_decimal
    ebreak

binary_to_decimal:
    mv t0, a0
    li a0, 0
    li t1, 2
    
process_bit:
    lbu t2, 0(t0)
    beq t2, t1, return
    
    slli a0, a0, 1
    or a0, a0, t2
    
    addi t0, t0, 1
    j process_bit
    
return:
    ret
