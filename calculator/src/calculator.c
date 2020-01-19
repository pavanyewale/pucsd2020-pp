#include "../include/functions.h"

float solve2(char ops[500], float number[500], int currentOps, int allOps)
{
    if (currentOps == allOps)
    {

        return number[allOps];
    }
    switch (ops[currentOps])
    {
    case '+':
        return add(number[currentOps], solve2(ops, number, currentOps + 1, allOps));
    case '-':
        number[currentOps + 1] *= -1;
        return add(number[currentOps], solve2(ops, number, currentOps + 1, allOps));

    case '/':
        number[currentOps + 1] = division(number[currentOps], number[currentOps + 1]);
        return solve2(ops, number, currentOps + 1, allOps);
    case '*':
        number[currentOps + 1] = mult(number[currentOps], number[currentOps + 1]);
        return solve2(ops, number, currentOps + 1, allOps);
    case '%':
        number[currentOps + 1] = percentage(number[currentOps], number[currentOps + 1]);
        return solve2(ops, number, currentOps + 1, allOps);
    default:
        break;
    }
    return 0;
}

float calculate(char str[1000])
{
    float numbers[500];
    char operations[500];
    int numberTop = -1, opsTop = -1;
    int start = -1;
    for (int i = 0; i < 1000; i++)
    {
        if (str[i] == '\n' || str[i] == '\0')
        {

            if (start != -1)
            {
                str[i] = '\0';
                numberTop++;
                numbers[numberTop] = convertToFloat(&str[start]);
                start = -1;
            }
            break;
        }
        if (str[i] == '+' || str[i] == '-' || str[i] == '/' || str[i] == '*' || str[i] == '%')
        {
            opsTop++;
            operations[opsTop] = str[i];
            if (start != -1)
            {
                numberTop++;
                str[i] = '\0';
                numbers[numberTop] = convertToFloat(&str[start]);
                start = -1;
            }
            continue;
        }
        if (start == -1)
        {
            start = i;
        }
    }

    return solve2(operations, numbers, 0, opsTop + 1);
}
