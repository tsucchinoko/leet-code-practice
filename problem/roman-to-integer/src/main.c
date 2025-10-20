#include <stdio.h>
#include <string.h>
#include <stddef.h>

#include <stddef.h>

static int value(char c) {
    if (c == 'I') return 1;
    else if (c == 'V') return 5;
    else if (c == 'X') return 10;
    else if (c == 'L') return 50;
    else if (c == 'C') return 100;
    else if (c == 'D') return 500;
    else if (c == 'M') return 1000;
    return 0;
}

int romanToInt(char* s) {
    if (s == NULL) return 0;
    int sum = 0;
    for (size_t i = 0; s[i] != '\0'; ) {
        int cur = value((char)s[i]);
        if (s[i + 1] != '\0') {
            int next = value((char)s[i + 1]);
            if (cur < next) {
                sum += next - cur;
                i += 2;
                continue;
            }
        }
        sum += cur;
        i += 1;
    }
    return sum;
}

int main(void) {
    const char *examples[] = {"III", "LVIII", "MCMXCIV", "I", "MMMCMXCIX", NULL};
    for (size_t i = 0; examples[i] != NULL; ++i) {
        char buf[32];
        strncpy(buf, examples[i], sizeof(buf));
        buf[sizeof(buf) - 1] = '\0';
        printf("%s -> %d\n", buf, romanToInt(buf));
    }

    // stdin から読み取りたい場合は以下のコメントを外してください。
    // 入力例: III\nLVIII\nMCMXCIV\n（EOF まで）
    /*
    char input[64];
    while (scanf("%63s", input) == 1) {
        printf("%s -> %d\n", input, romanToInt(input));
    }
    */

    return 0;
}
