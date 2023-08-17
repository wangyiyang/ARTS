#include <stdio.h>

int search(int* nums, int numsSize, int target) {
    int left = 0, right = numsSize - 1, mid;
    while (left <= right) {
        mid = (left + right) / 2;
        if (nums[mid] == target) return mid;
        else if (nums[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}

int main() {
    int nums[] = {-1, 0, 3, 5, 9, 12};
    int target = 9;
    int numsSize = sizeof(nums) / sizeof(nums[0]);
    int ans = search(nums, numsSize, target);
    printf("%d\n", ans);
    return 0;
}
