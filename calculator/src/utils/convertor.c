#include <stdlib.h>
#include "../../include/functions.h"

float convertToFloat(char no[100])
{
    char *end;
    return strtof(no, &end);
}