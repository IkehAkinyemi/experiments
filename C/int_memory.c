#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int *n = malloc(3 * sizeof(int));
  
  n[0] = 12;
  n[1] = 15;
  n[2] = 16;

  printf("%i,", *(n+1));
  free(n);
}