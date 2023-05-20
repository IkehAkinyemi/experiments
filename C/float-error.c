#include <stdio.h>

float sum_element(float a[], unsigned length) {
  int i;
  float result = 0;

  if (length != 0) {
  for (i = 0; i <= length-1; i++)
    result += a[i];
  }
  
  return result;
}

int main() {
  float a[] = {1, 2, 3, 4};
  float sum = sum_element(a, 0);
  printf("%.2f", sum);
  return 0;
}