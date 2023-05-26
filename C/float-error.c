#include <stdio.h>
#include <string.h>

float sum_element(float a[], int length) {
  int i;
  float result = 0;

  for (i = 0; i <= length-1; i++)
    result += a[i];

  return result;
}
// float sum_element(float a[], unsigned length) {
//   int i;
//   float result = 0;

//   if (length != 0) {
//   for (i = 0; i <= length-1; i++)
//     result += a[i];
//   }

//   return result;
// }

// Determine whether string s is longer than string t
// WARNING: This function is buggy
int strlonger(char *s, char *t) {
  return strlen(s) - strlen(t) > 0;
}

int main() {
  float a[] = {1, 2, 3, 4};
  float sum = sum_element(a, 0);
  // printf("%.2f", sum);

  int bool = strlonger("c", "CCC");
  printf("%d", bool);
  return 0;
}