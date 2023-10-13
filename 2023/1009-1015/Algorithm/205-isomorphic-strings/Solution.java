import java.util.HashMap;

class Solution{
    public boolean isIsomorphic(String s, String t){
        if(s.length() != t.length()){
            return false;
        }
        HashMap<Character, Character> map = new HashMap<>();
        for(int i = 0; i < s.length(); i++){
            char sChar = s.charAt(i);
            char tChar = t.charAt(i);
            if(map.containsKey(sChar)){
                if(map.get(sChar) != tChar){
                    return false;
                }
            }else{
                if(map.containsValue(tChar)){
                    return false;
                }
                map.put(sChar, tChar);
            }
        }
        return true;
    }
    
    public static void main(String[] args){
        Solution sol = new Solution();
        String s = "egg";
        String t = "add";
        System.out.println(sol.isIsomorphic(s, t));
    }
}

