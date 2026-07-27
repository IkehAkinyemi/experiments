.globl _start

.text
_start:
    li a0, 0xff0f
    jal ra, count_bits
    ebreak

count_bits:
    li t0, 0
    li t1, 0
    li t2, 32

count_loop:
    beqz t2, count_end

    andi t3, a0, 1
    beqz t3, is_zero_bit
is_one_bit:
    addi t1, t1, 1
    j shift_next
is_zero_bit:
    addi t0, t0, 1
shift_next:
    srli a0, a0, 1
    addi t2, t2, -1
    j count_loop
count_end:
    mv a1, t1
    mv a0, t0
    ret
    