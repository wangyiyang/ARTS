#include "stdio.h"
#include "stdlib.h"
#include "string.h"
#include "stdbool.h"


int* circularGameLosers(int n, int k, int* returnSize){
    bool visit[n];
    memset(visit, 0, sizeof(visit));
    for (int i = k, j = 0; !visit[j]; i += k) {
        visit[j] = true;
        j = (j + i) % n;
    }
    int *ans = (int *)malloc(sizeof(int) * n);
    int pos = 0;
    for (int i = 0; i < n; i++) {
        if (!visit[i]) {
            ans[pos++] = i + 1;
        }
    }
    *returnSize = pos;
    return ans;
}

int main(int argc, char *argv[]){
    int n = 0, k = 0;
    scanf("%d %d", &n, &k);
    int returnSize = 0;
    int *ans = circularGameLosers(n, k, &returnSize);
    for (int i = 0; i < returnSize; i++) {
        printf("%d ", ans[i]);
    }
    printf("\n");
    return 0;
}
