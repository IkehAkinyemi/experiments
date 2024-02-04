#include <stdio.h>

// 1.3 The For Statement
// int main() {
//   for (int i = 1; i<= 5; i++) {
//     printf("%d\n", i);
//   }

//   return 0;
// }

// 1.5.1 Character Input and Output
// int main() {
//   int c;

//   c = getchar();
//   while (c != EOF)
//   {
//     // printf("%d\n", c);
//     putchar(c);
//     c = getchar();
//   }
//   printf("\n");
// }

// 1.5.2 Character Input and Output
// int main() {
//   long nc;

//   nc = 0;
//   while (getchar() != EOF)
//     ++nc;
//   printf("%ld\n", nc);
// }

// 1.5.3 Character Input and Output
// int main() {
//   int nl, c;

//   nl = 0;
//   while ((c = getchar()) != EOF) {
//     if (c == '\n')
//       ++nl;
//   }
  
//   printf("%d\n", nl);
// }

// 1.5.4 Character Input and Output
// #define IN 1 // inside a word
// #define OUT 2 // outside a word

// int main() {
//   int c, nl, nw, nc, state;

//   state = OUT;
//   nl = nw = nc = 0;
//   while ((c = getchar()) != EOF) {
//     ++nc;
//     if (c == '\n')
//       ++nl;
    
//     if (c == ' ' || c == '\n' || c == '\t')
//       state = OUT;

//     else if (state == OUT) {
//       state = IN;
//       ++nw;
//     }
//   }

//   printf("%d %d %d\n", nl, nw, nc);
// }

// #define MAX_WORD_LENGTH 20
// #define OUT 0
// #define IN 1

// int main() {
//   int c, i, j, nc, state;
//   int word_length[MAX_WORD_LENGTH];

//   state = OUT;
//   nc = 0; // number of characters in a word

//   for (i = 0; i < MAX_WORD_LENGTH; ++i) {
//     word_length[i] = 0;
//   }
  
//   // Read characters
//   while ((c = getchar()) != EOF) {
//     if (c == ' ' || c == '\n' || c == '\t') {
//       if (state == IN) {
//         if (nc < MAX_WORD_LENGTH)
//           ++word_length[nc];

//         state = OUT;
//         nc = 0;
//       }
//     } else if (state == OUT) {
//       state = IN;
//       nc = 1; // Start of new word
//     } else {
//       ++nc; // Inside a word
//     }
//   }

//   // Print the horizontal histogram
//   for (i = 1; i < MAX_WORD_LENGTH; ++i) {
//     printf("%2d: ", i);
//     for (j = 0; j < word_length[i]; ++j) {
//       putchar('*');
//     }
//     putchar('\n');
//   }
// }