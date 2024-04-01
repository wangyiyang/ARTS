class Solution {
    public int mySqrt(int x) {
        if (x == 0) return 0;
        int left = 1, right = x;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (mid <= x / mid && (mid + 1) > x / (mid + 1)) {
                return mid;
            } else if (mid > x / mid) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left;

    }

    public static void main(String[] args) {
        Solution sol = new Solution();
        System.out.println(sol.mySqrt(4)); // 2
        System.out.println(sol.mySqrt(8)); // 2
        System.out.println(sol.mySqrt(9)); // 3
        System.out.println(sol.mySqrt(16)); // 4
        System.out.println(sol.mySqrt(25)); // 5
        System.out.println(sol.mySqrt(36)); // 6
        System.out.println(sol.mySqrt(49)); // 7
        System.out.println(sol.mySqrt(64)); // 8
        System.out.println(sol.mySqrt(81)); // 9
        System.out.println(sol.mySqrt(100)); // 10
        System.out.println(sol.mySqrt(121)); // 11
        System.out.println(sol.mySqrt(144)); // 12
        System.out.println(sol.mySqrt(169)); // 13
        System.out.println(sol.mySqrt(196)); // 14
        System.out.println(sol.mySqrt(225)); // 15
        System.out.println(sol.mySqrt(256)); // 16
        System.out.println(sol.mySqrt(289)); // 17
        System.out.println(sol.mySqrt(324)); // 18
        System.out.println(sol.mySqrt(361)); // 19
        System.out.println(sol.mySqrt(400)); // 20
        System.out.println(sol.mySqrt(441)); // 21
        System.out.println(sol.mySqrt(484)); // 22
        System.out.println(sol.mySqrt(529)); // 23
        System.out.println(sol.mySqrt(576)); // 24
        System.out.println(sol.mySqrt(625)); // 25
        System.out.println(sol.mySqrt(676)); // 26
        System.out.println(sol.mySqrt(729)); // 27
        System.out.println(sol.mySqrt(784)); // 28
        System.out.println(sol.mySqrt(841)); // 29
        System.out.println(sol.mySqrt(900)); // 30
        System.out.println(sol.mySqrt(961)); // 31
        System.out.println(sol.mySqrt(1024)); // 32
    }
}
