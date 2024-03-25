// Source: https://leetcode.com/problems/happy-number/
import java.util.HashSet;
import java.util.Set;

class Solution {
    public boolean isHappy(int n) {
        Set<Integer> set = new HashSet<>();
        while (n != 1) {
            if (set.contains(n)) { // 如果n已经在set里面了，说明n已经循环了，不可能是happy number
                return false;
            }
            set.add(n);
            n = getNext(n);
        }
        return true;
    }

    private int getNext(int n) {
        int sum = 0;
        while (n > 0) {
            int digit = n % 10;
            sum += digit * digit;
            n /= 10;
        }
        return sum;
    }

    public static void main(String[] args) {
        Solution sol = new Solution();
        int n = 19;
        System.out.println("Is " + n + " a happy number? " + sol.isHappy(n));
    }

}
