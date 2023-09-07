import java.util.Arrays;

public class Solution {
    public String reverseWords(String s) {
        String[] strs = s.trim().split("\\s+");
        System.out.println(Arrays.toString(strs));
        StringBuilder sb = new StringBuilder();
        for (int i = strs.length - 1; i >= 0; i--) {
            sb.append(strs[i]).append(" ");
        }
        return sb.toString().trim();
    }
    
    public static void main(String[] args) {
        Solution s = new Solution();
        String str = "  hello world!  ";
        System.out.println(s.reverseWords(str));
    }
}
