

class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) return s;
        StringBuilder[] sbs = new StringBuilder[numRows];
        for (int i = 0; i < numRows; i++) {
            sbs[i] = new StringBuilder();
        }
        int i = 0;
        int flag = -1;
        for (char c : s.toCharArray()) {
            sbs[i].append(c);
            if (i == 0 || i == numRows - 1) flag = -flag;
            i += flag;
        }
        StringBuilder sb = new StringBuilder();
        for (StringBuilder sbsb : sbs) {
            sb.append(sbsb);
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        String str = "PAYPALISHIRING";
        System.out.println(s.convert(str, 3));
    }
}
