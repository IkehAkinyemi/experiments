#include <stdio.h>

void swap(int *x, int *y);

int main(void) 
{
  int y = 24;
  int x = 20;

  printf("x: %i, y: %i\n", x, y);

  swap(&x, &y);

  printf("x: %i, y: %i", x, y);
}

void swap(int *x, int *y)
{
  int tmp = *x;
  *x = *y;
  *y = tmp;
}

