#include <stdio.h>
#include <stdlib.h>

char* get_string(char *d);

int main(void)
{
  char *s = get_string("Input name: ");
  printf("%s", s);
}

char* get_string(char *d)
{
  printf("%s", d);
  char *s = malloc(4);
  scanf("%s", s);
  return s;
}