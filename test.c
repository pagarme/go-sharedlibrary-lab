#include <stdio.h>

extern int tms_get_tables(unsigned char *payload, unsigned int length, void *write_cb, void *user_data);

int main() {
    printf("wtf\n");

    int result = tms_get_tables(0, 0, 0, 0);

    printf("PQP Q GAMBIARRA %d\n", result);

    return 1;
}
