/**
 * This class provides a solution to the h-index problem.
 * The h-index is defined as the maximum value of h such that the given array of citations has at least h papers with h or more citations.
 */
class Solution {
    /**
     * Calculates the h-index for the given array of citations.
     * @param citations the array of citations
     * @return the h-index
     */
    public int hIndex(int[] citations) {
        Arrays.sort(citations);   // sort the array in ascending order
        System.out.println(Arrays.toString(citations));
        int h = 0;
        for (int i = citations.length - 1; i >= 0; i--) {
            if (citations[i] > h) {
                h++;
            } else {
                break;
            }
        }
        return h;
    }
}

class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] citations = {3, 0, 6, 1, 5};
        System.out.println(solution.hIndex(citations));
    }
}
