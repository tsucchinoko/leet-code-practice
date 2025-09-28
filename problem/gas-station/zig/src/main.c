#include <stdio.h>
#include <stdint.h>

/**
 * Gas Station Problem
 *
 * There are n gas stations along a circular route, where the amount of gas at
 * the ith station is gas[i]. You have a car with an unlimited gas tank and it
 * costs cost[i] of gas to travel from the ith station to its next (i + 1)th station.
 *
 * Return the starting gas station's index if you can travel around the circuit
 * once in the clockwise direction, otherwise return -1.
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
int canCompleteCircuit(int* gas, int gasSize, int* cost, int costSize) {
    if (gasSize == 0 || gasSize != costSize) {
        return -1;
    }

    int n = gasSize;
    int64_t total = 0;  // 全体の収支
    int64_t tank = 0;   // 現在のタンク
    int start = 0;      // 開始位置

    for (int i = 0; i < n; i++) {
        int64_t diff = (int64_t)gas[i] - (int64_t)cost[i];
        total += diff;
        tank += diff;

        if (tank < 0) {
            // 現在の開始位置では完走不可、次の位置から開始
            start = i + 1;
            tank = 0;
        }
    }

    // 全体の収支が負なら完走不可
    if (total < 0) {
        return -1;
    }

    return start;
}

void printArray(const char* label, int* arr, int size) {
    printf("%s[", label);
    for (int i = 0; i < size; i++) {
        if (i > 0) printf(", ");
        printf("%d", arr[i]);
    }
    printf("]");
}

int main() {
    // Example 1: Expected start = 3
    int gas1[] = {1, 2, 3, 4, 5};
    int cost1[] = {3, 4, 5, 1, 2};
    int size1 = sizeof(gas1) / sizeof(gas1[0]);

    // Example 2: Expected start = -1
    int gas2[] = {2, 3, 4};
    int cost2[] = {3, 4, 3};
    int size2 = sizeof(gas2) / sizeof(gas2[0]);

    // Test results
    int result1 = canCompleteCircuit(gas1, size1, cost1, size1);
    int result2 = canCompleteCircuit(gas2, size2, cost2, size2);

    printf("Example 1:\n");
    printArray("  gas:  ", gas1, size1);
    printf("\n");
    printArray("  cost: ", cost1, size1);
    printf("\n");
    printf("  start = %d\n\n", result1);

    printf("Example 2:\n");
    printArray("  gas:  ", gas2, size2);
    printf("\n");
    printArray("  cost: ", cost2, size2);
    printf("\n");
    printf("  start = %d\n\n", result2);

    // Algorithm explanation
    printf("Algorithm explanation:\n");
    printf("1. Track total gas-cost difference across all stations\n");
    printf("2. Track current tank level from current start position\n");
    printf("3. If tank goes negative, reset start to next station\n");
    printf("4. If total difference is negative, circuit is impossible\n");
    printf("5. Otherwise, return the final start position\n");

    return 0;
}
