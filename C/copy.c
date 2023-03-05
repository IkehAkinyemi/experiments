#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void)
{
  char *c = "Hello";
  char *d = malloc(strlen(c) + 1);

  for (int i = 0; i < strlen(c); i++)
  {
    d[i] = c[i];
  }
  
  printf("%s\n", c);
  printf("%s\n", d);

  free(d);
}