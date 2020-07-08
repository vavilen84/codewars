#include <stdio.h>
#include <regex.h>
#include <stdlib.h>

int is_ip_valid(char * ip) {
    regex_t regex;
    int err;
    char * pattern;
    pattern = "^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$";
    err = regcomp(&regex, pattern, REG_EXTENDED);
    if (err != 0) {
        char buff[512];
        regerror(err, &regex, buff, sizeof(buff));
        printf("%s", buff);
        exit(1);
    }
    regmatch_t pm;
    err = regexec(&regex, ip, 0, &pm, 0);
    if (err == 0) {
        return 1;
    }
    return 0;
}

void test_invalid() {
    int elSize, ipSize, r, i;
    elSize = 9;
    ipSize = 27;
    char invalid[elSize][ipSize] = {
            {""},
            {"invalid input"},
            {".1"},
            {"1.2.3"},
            {"1.2.3.4."},
            {"1.2.3.4.5"},
            {"123.456.78.90"},
            {"123.045.067.089"},
            {"0.2.3.4"},
    };
    for (i = 0; i < elSize; i++) {
        r = is_ip_valid(invalid[i]);
        if (r != 0) {
            printf("Failed! Invalid case");
            exit(1);
        }
    }
}

void test_valid() {
    int i, r, elSize, ipSize;
    elSize = 1;
    ipSize = 27;
    char valid[elSize][ipSize] = {
            {"1.2.3.4"},
    };
    for (i = 0; i < elSize; i++) {
        r = is_ip_valid(valid[i]);
        if (r != 1) {
            printf("Failed! Valid case");
            exit(1);
        }
    }
}

int main() {
    test_valid();
    test_invalid();
    printf("OK!");
    return 0;
}