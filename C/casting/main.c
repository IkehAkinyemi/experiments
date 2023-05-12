#include <stdio.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len) {
  int i;
  for (i = 0; i < len; i++)
    printf("%.2x", start[i]);
  printf("\n");
}

void show_int(int x) {
  show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x) {
  show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x) {
  show_bytes((byte_pointer) &x, sizeof(void *));
}

int main() {
  int a = 12345;
  float b = 3.14159;
  int *c = &a;

  printf("Byte representation of integer %d:\n", a);
  show_int(a);

  printf("Byte representation of float %f:\n", b);
  show_float(b);

  printf("Byte representation of pointer %p:\n", (void *)c);
  show_pointer(c);

  return 0;
}
