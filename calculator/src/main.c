#include <stdio.h>
#include "../include/functions.h"
int main(void)
{
    char str[1000];
    float ans;
    int err;
    while (1)
    {
        printf("=>");
        scanf("%s", str);

        err = checkString(str);
        if (err)
        {
            printf("Invalid Input");
            printf("\n");
            continue;
        }
        ans = calculate(str);
        printf("%f", ans);
        printf("\n");
    }
    return 0;
}