.globl _start

.data
.align 0
test_array: .byte 45, -12, 0, 88, -90, 3, 15

.text
.align 2

_start:
    la a0, test_array
    li a1, 7
    jal ra, bubble_sort
    ebreak

bubble_sort:
    li t0, 1
    ble a1, t0, sort_end

    # outer loop
    addi a1, a1, -1

outer_loop:
    li t0, 0
    mv t1, a0
    mv t2, a1

inner_loop:
    beqz t2, inner_end

    lb t3, 0(t1)
    lb t4, 1(t1)

    ble t3, t4, skip_swap
    sb t4, 0(t1)
    sb t3, 1(t1)
    li t0, 1

skip_swap:
    addi t1, t1, 1
    addi t2, t2, -1
    j inner_loop

inner_end:
    bnez t0, outer_loop

sort_end:
    ret
