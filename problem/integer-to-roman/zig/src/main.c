// main.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdbool.h>

// digit_to_roman: d in 0..9, one/five/ten are C-strings
// Returns heap-allocated char* (not null-terminated) of length *out_len.
// Caller must free() the returned pointer. If length is 0, returned pointer may be NULL.
static char* digit_to_roman(uint8_t d, const char* one, const char* five, const char* ten, size_t* out_len) {
    size_t len_one = strlen(one);
    size_t len_five = strlen(five);
    size_t len_ten = strlen(ten);

    if (d == 0) {
        *out_len = 0;
        return NULL;
    } else if (d <= 3) {
        *out_len = d * len_one;
        char* buf = (char*)malloc(*out_len);
        if (!buf) return NULL;
        for (uint8_t i = 0; i < d; i++) {
            memcpy(buf + i * len_one, one, len_one);
        }
        return buf;
    } else if (d == 4) {
        *out_len = len_one + len_five;
        char* buf = (char*)malloc(*out_len);
        if (!buf) return NULL;
        memcpy(buf, one, len_one);
        memcpy(buf + len_one, five, len_five);
        return buf;
    } else if (d == 5) {
        *out_len = len_five;
        char* buf = (char*)malloc(*out_len);
        if (!buf) return NULL;
        memcpy(buf, five, len_five);
        return buf;
    } else if (d <= 8) {
        uint8_t extra = d - 5;
        *out_len = len_five + extra * len_one;
        char* buf = (char*)malloc(*out_len);
        if (!buf) return NULL;
        memcpy(buf, five, len_five);
        for (uint8_t i = 0; i < extra; i++) {
            memcpy(buf + len_five + i * len_one, one, len_one);
        }
        return buf;
    } else { // d == 9
        *out_len = len_one + len_ten;
        char* buf = (char*)malloc(*out_len);
        if (!buf) return NULL;
        memcpy(buf, one, len_one);
        memcpy(buf + len_one, ten, len_ten);
        return buf;
    }
}
// LeetCode signature
char* intToRoman(int num) {
    // Arrays of Roman strings for each digit place.
    // Index = digit 0..9
    static const char* thousands[] = {"", "M", "MM", "MMM"}; // 0..3
    static const char* hundreds[]  = {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
    static const char* tens[]      = {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
    static const char* ones[]      = {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};

    // Extract digits
    int t = num / 1000;
    int h = (num % 1000) / 100;
    int te = (num % 100) / 10;
    int o = num % 10;

    const char* s_t = thousands[t];
    const char* s_h = hundreds[h];
    const char* s_te = tens[te];
    const char* s_o = ones[o];

    // Compute total length
    size_t len = strlen(s_t) + strlen(s_h) + strlen(s_te) + strlen(s_o);

    // Allocate result (+1 for null-terminator)
    char* res = (char*)malloc(len + 1);
    if (!res) return NULL;

    // Concatenate
    char* p = res;
    if (s_t[0]) { strcpy(p, s_t); p += strlen(s_t); }
    if (s_h[0]) { strcpy(p, s_h); p += strlen(s_h); }
    if (s_te[0]) { strcpy(p, s_te); p += strlen(s_te); }
    if (s_o[0]) { strcpy(p, s_o); p += strlen(s_o); }
    *p = '\0';

    return res;
}

static bool expect_equal_str(const char* got, size_t got_len, const char* want) {
    size_t want_len = strlen(want);
    if (got_len != want_len) return false;
    return (memcmp(got, want, got_len) == 0);
}

int main(void) {
    // demo printing
    uint16_t nums[] = {3749, 58, 1994, 4, 9, 40, 944};
    for (size_t i = 0; i < sizeof(nums)/sizeof(nums[0]); i++) {
        char* r = intToRoman(nums[i]);
        if (r) {
            printf("%u -> %s\n", (unsigned)nums[i], r);
            free(r);
        } else {
            printf("%u -> (error)\n", (unsigned)nums[i]);
        }
    }

    // basic tests
    struct {
        uint16_t num;
        const char* want;
    } cases[] = {
        {1, "I"},
        {3, "III"},
        {4, "IV"},
        {9, "IX"},
        {58, "LVIII"},
        {1994, "MCMXCIV"},
        {3749, "MMMDCCXLIX"},
        {40, "XL"},
        {90, "XC"},
        {400, "CD"},
        {900, "CM"},
        {3999, "MMMCMXCIX"},
    };

    size_t ncases = sizeof(cases) / sizeof(cases[0]);
    size_t failed = 0;
    for (size_t i = 0; i < ncases; i++) {
        size_t len;
        char* got = intToRoman(cases[i].num);
        if (!got) {
            printf("Test %zu: %u -> NULL (expected %s)\n", i, (unsigned)cases[i].num, cases[i].want);
            failed++;
            continue;
        }
        if (!expect_equal_str(got, len, cases[i].want)) {
            printf("Test %zu FAILED: %u -> '%s' (len=%zu), want '%s'\n", i, (unsigned)cases[i].num, got, len, cases[i].want);
            failed++;
        } else {
            // printf("Test %zu ok: %u -> %s\n", i, (unsigned)cases[i].num, got);
        }
        free(got);
    }

    if (failed == 0) {
        printf("All tests passed (%zu cases)\n", ncases);
        return 0;
    } else {
        printf("%zu tests failed out of %zu\n", failed, ncases);
        return 1;
    }
}
