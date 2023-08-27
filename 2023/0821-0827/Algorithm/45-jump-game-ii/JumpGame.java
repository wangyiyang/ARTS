class Solution {
    public int jump(int[] nums) {
        int minJump = 0;
        int curEnd = 0;
        int curFarthest = 0;
        
        for(int i=0; i<nums.length-1; i++) {
            curFarthest = Math.max(curFarthest, i+nums[i]);
            if(i == curEnd) {
                minJump++;
                curEnd = curFarthest;
            }
        }
        
        return minJump;
    }
}

class JumpGame {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int[] nums = {2,3,1,1,4};
        int minJump = sol.jump(nums);
        System.out.println(minJump);
    }
}