#include <stdio.h>
#include<stdlib.h>

char *p = NULL;
int r(char **pptr, int len) {
    p = (char*)malloc(sizeof(char)*len);
    if(p == NULL) {
        return -1;
    }
    pptr = &p;
    return 0;
}

int main()
{
    char** pptr=NULL;
    int len = 10;
    int res = r(pptr, len);
    printf("Hello world\n");
    return 0;
}

