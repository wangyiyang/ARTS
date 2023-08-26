
#include <stdbool.h>
#include <stdio.h>

bool canJump(int *nums, int numsSize) {
    int max = 0;
    for (int i = 0; i < numsSize; i++) {
        if (i > max) {
            return false;
        }
        if (i + nums[i] > max) {
            max = i + nums[i];
        }
    }
    return true;
}

int main(int argc, char *argv[]) {
    int nums[] = {2, 3, 1, 1, 4};
    int numsSize = sizeof(nums) / sizeof(nums[0]);
    bool result = canJump(nums, numsSize);
    printf("result = %s\n", result ? "true" : "false");
    return 0;
}
