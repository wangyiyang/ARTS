import java.util.Stack;

class Solution {
    public String simplifyPath(String path) {
        String[] paths = path.split("/");
        Stack<String> stack = new Stack<>();
        for (String p : paths) {
            if (p.equals("..")) {
                if (!stack.isEmpty()) {
                    stack.pop();
                }
            } else if (!p.equals(".") && !p.equals("")) {
                stack.push(p);
            }
        }
        return "/" + String.join("/", stack);
    }

    public static void main(String[] args) {
        Solution sol = new Solution();
        String path = "/a/./b/../../c/";
        System.out.println("Simplified path: " + sol.simplifyPath(path));
    }
}
