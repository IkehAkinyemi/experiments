.globl _start
.section .text.bios

# NS16550A UART Memory-Mapped Registers (QEMU 'virt' machine)
.equ UART_BASE,    0x10000000
.equ UART_RBR,     0x00        # Receiver Buffer Register (Read)
.equ UART_THR,     0x00        # Transmitter Holding Register (Write)
.equ UART_LSR,     0x05        # Line Status Register
.equ LSR_RX_READY, 0x01        # Bit 0: Data Ready (DR)
.equ LSR_TX_READY, 0x20        # Bit 5: Transmitter Holding Register Empty (THRE)

_start:
    li s0, UART_BASE               # Initialize UART base address in s0

    # Print initialization prompt
    la a0, prompt_msg
    jal ra, print_string

main_loop:
    # 1. Wait for RX data
wait_rx:
    lb t0, UART_LSR(s0)            # Load Line Status Register
    andi t1, t0, LSR_RX_READY      # Mask bit 0
    beqz t1, wait_rx               # If 0, data is not ready; loop

    # 2. Read character from hardware
    lb s1, UART_RBR(s0)            # Load byte from Receiver Buffer

    # 3. Process character (Toggle case for alphabetic characters)
    # Check if 'a' <= char <= 'z'
    li t0, 'a'
    blt s1, check_upper
    li t0, 'z'
    bgt s1, tx_char
    # It is lowercase: subtract 32 to make it uppercase
    addi s1, s1, -32
    j tx_char

check_upper:
    # Check if 'A' <= char <= 'Z'
    li t0, 'A'
    blt s1, tx_char
    li t0, 'Z'
    bgt s1, tx_char
    # It is uppercase: add 32 to make it lowercase
    addi s1, s1, 32

tx_char:
    # 4. Wait for TX ready
wait_tx:
    lb t0, UART_LSR(s0)            # Load Line Status Register
    andi t1, t0, LSR_TX_READY      # Mask bit 5
    beqz t1, wait_tx               # If 0, transmitter is busy; loop

    # 5. Write processed character to hardware
    sb s1, UART_THR(s0)

    # 6. Handle Carriage Return (\r) by injecting a Newline (\n)
    li t0, '\r'
    bne s1, t0, main_loop
wait_tx_nl:
    lb t0, UART_LSR(s0)
    andi t1, t0, LSR_TX_READY
    beqz t1, wait_tx_nl
    li t2, '\n'
    sb t2, UART_THR(s0)

    j main_loop                    # Return to start of loop

# Subroutine: Print a null-terminated string
# Arguments: a0 = address of string
print_string:
    mv t2, a0                      # Copy string pointer to t2
print_loop:
    lb t3, 0(t2)                   # Load current byte
    beqz t3, print_end             # If null terminator, exit subroutine

wait_tx_str:
    lb t0, UART_LSR(s0)            # Read LSR
    andi t1, t0, LSR_TX_READY      # Check THRE bit
    beqz t1, wait_tx_str           # Wait until ready

    sb t3, UART_THR(s0)            # Transmit byte
    addi t2, t2, 1                 # Increment string pointer
    j print_loop
print_end:
    ret                            # Return to caller

    .section .rodata
prompt_msg:
    .asciz "Interactive UART initialized. Type characters to see their case toggled:\r\n> "