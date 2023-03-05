#include <stdlib.h>
#include <stdio.h>

int main(void) 
{
  char *s;
  s = malloc(4);

  scanf(s);
  printf("string val: %s", s);
  free(s);
}