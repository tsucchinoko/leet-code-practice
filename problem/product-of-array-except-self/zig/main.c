#include <stdio.h>
#include <stdlib.h>

/**
 * Product of Array Except Self
 * 
 * Given an integer array nums, return an array answer such that answer[i] 
 * is equal to the product of all the elements of nums except nums[i].
 * 
 * Time Complexity: O(n)
 * Space Complexity: O(1) extra space (not counting the output array)
 */
int* productExceptSelf(int* nums, int numsSize, int* returnSize) {
    // Allocate memory for output array
    int* out = (int*)malloc(numsSize * sizeof(int));
    if (out == NULL) {
        *returnSize = 0;
        return NULL;
    }
    
    *returnSize = numsSize;
    
    // First pass: store left products in out
    int left = 1;
    for (int i = 0; i < numsSize; i++) {
        out[i] = left;
        left = left * nums[i];
    }
    
    // Second pass: multiply by right products on the fly
    int right = 1;
    for (int i = numsSize - 1; i >= 0; i--) {
        out[i] = out[i] * right;
        right = right * nums[i];
    }
    
    return out;
}

void printArray(const char* label, int* arr, int size) {
    printf("%s", label);
    for (int i = 0; i < size; i++) {
        if (i != 0) printf(", ");
        printf("%d", arr[i]);
    }
    printf("\n");
}

int main() {
    // Example 1
    int nums1[] = {1, 2, 3, 4};
    int size1 = sizeof(nums1) / sizeof(nums1[0]);
    int returnSize1;
    
    int* out1 = productExceptSelf(nums1, size1, &returnSize1);
    if (out1 != NULL) {
        printArray("input: ", nums1, size1);
        printArray("output: ", out1, returnSize1);
        free(out1);
    }
    
    // Example 2
    int nums2[] = {-1, 1, 0, -3, 3};
    int size2 = sizeof(nums2) / sizeof(nums2[0]);
    int returnSize2;
    
    int* out2 = productExceptSelf(nums2, size2, &returnSize2);
    if (out2 != NULL) {
        printArray("input: ", nums2, size2);
        printArray("output: ", out2, returnSize2);
        free(out2);
    }
    
    return 0;
}