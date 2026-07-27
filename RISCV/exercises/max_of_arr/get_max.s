.globl _start

.data
array_data: .word 123, -11, 10, 431, 200

.text
_start:
    la a0, array_data
    li a1, 5
    jal ra, get_max
    ebreak

get_max:
    blez a1, handle_empty
    lw t0, 0(a0)
    addi a0, a0, 4
    addi a1, a1, -1
loop:
    beqz a1, return
    lw t1, 0(a0)
    bge t0, t1, skip_update
    mv t0, t1
skip_update:
    addi a0, a0, 4
    addi a1, a1, -1
    j loop
handle_empty:
    li t0, -1
return:
    mv a0, t0
    ret
    
    