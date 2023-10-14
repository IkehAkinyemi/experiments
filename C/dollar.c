#include <stdio.h>

#define FIRST_AMOUNT 20
#define SECOND_AMOUNT 10
#define THIRD_AMOUNT 5
#define FOURTH_AMOUNT 1

int main()
{
  int amount;

  printf("Enter dollar amount: ");
  scanf("%d", &amount);

  if (amount >= FIRST_AMOUNT)
  {
    printf("$20 bills: %d\n", amount / FIRST_AMOUNT);
    amount = amount % FIRST_AMOUNT;
  }

  if (amount >= SECOND_AMOUNT)
  {
    printf("$10 bills: %d\n", amount / SECOND_AMOUNT);
    amount = amount % SECOND_AMOUNT;
  }

  if (amount >= THIRD_AMOUNT)
  {
    printf("$5 bills: %d\n", amount / THIRD_AMOUNT);
    amount = amount % THIRD_AMOUNT;
  }

  if (amount >= FOURTH_AMOUNT)
  {
    printf("$1 bills: %d\n", amount / FOURTH_AMOUNT);
    amount = amount % FOURTH_AMOUNT;
  }
}