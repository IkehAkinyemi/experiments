.globl _start

.data
.align 2
custom_stack_base: .space 400
custom_stack_top:
custom_sp: .word custom_stack_top

.text
_start:
    li a0, 0
    li a1, 1024
    jal ra, stack_operation
    ebreak

stack_operation:
    la t0, custom_sp
    lw t1, 0(t0)

    beqz a0, execute_push
execute_pop:
    lw a0, 0(t1)
    addi t1, t1, 4
    sw t1, 0(t0)
    ret
execute_push:
    addi t1, t1, -4
    sw a1, 0(t1)
    sw t1, 0(t0)
    ret
    