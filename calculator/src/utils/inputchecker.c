#include "../../include/functions.h"
int checkString(char str[1000])
{
    int dot = 0, number = 0;
    for (int i = 0; i < 1000; i++)
    {
        if (str[i] == '+' || str[i] == '-' || str[i] == '/' || str[i] == '*' || str[i] == '%')
        {
            if (!number)
            {
                return 1;
            }
            dot = 0;
            number = 0;
            continue;
        }
        if (str[i] == '.')
        {
            if (!number || dot)
            {
                return 1;
            }
            dot = 1;
            number = 0;
            continue;
        }
        if (str[i] >= '0' && str[i] <= '9')
        {
            number = 1;
            continue;
        }
        if (str[i] == '\n' || str[i] == '\0')
        {
            return !number;
        }
        return 1;
    }
    return 0;
}
