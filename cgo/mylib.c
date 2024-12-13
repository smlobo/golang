// C code
#include "mylib.h"
#include <stddef.h>

Result myFunction() {
    Result result;
    result.value1 = 10;
    result.value2 = 3.14;
    result.v3 = NULL;
    result.v4 = &result.value1;
    return result;
}
