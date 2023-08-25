#include <stdio.h>

int maxProfit(int* prices, int pricesSize){
    int profit = 0;
    for (int i = 1; i < pricesSize; i++) {
        if (prices[i] > prices[i-1]) {
            profit += prices[i] - prices[i-1];
        }
    }
    return profit;
}

int main(int argc, char *argv[]) {
    int prices[] = {7,1,5,3,6,4};
    int pricesSize = sizeof(prices) / sizeof(prices[0]);
    int profit = maxProfit(prices, pricesSize);
    printf("profit = %d\n", profit);
    return 0;
}