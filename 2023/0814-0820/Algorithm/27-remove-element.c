#include <stdio.h>

int removeElement(int* nums, int numsSize, int val){
    int i, j;
    for (i = 0, j = 0; i < numsSize; i++) {
        if (nums[i] != val) {
            nums[j++] = nums[i];
        }
    }
    return j;
}

int main() {
    int nums[] = {3, 2, 2, 3};
    int numsSize = sizeof(nums) / sizeof(nums[0]);
    int val = 3;
    int len = removeElement(nums, numsSize, val);
    printf("len: %d\n", len);
    for (int i = 0; i < len; i++) {
        printf("%d ", nums[i]);
    }
    printf("\n");
}

