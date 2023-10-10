class Solution{
    public void rotate(int[][] matrix){
        int n = matrix.length;
        for(int i = 0; i < n/2; i++){
            for(int j = 0; j < (n+1)/2; j++){
                int temp = matrix[i][j];
                matrix[i][j] = matrix[n-j-1][i]; // 1 -> 2
                matrix[n-j-1][i] = matrix[n-i-1][n-j-1]; // 2 -> 3
                matrix[n-i-1][n-j-1] = matrix[j][n-i-1]; // 3 -> 4
                matrix[j][n-i-1] = temp; // 4 -> 1
            }
        }
    }

    public void print(int[][] matrix){
        for(int i = 0; i < matrix.length; i++){
            for(int j = 0; j < matrix.length; j++)
                System.out.print(matrix[i][j] + " ");
            System.out.println();
        }
    }

    public static void main(String[] args){
        Solution s = new Solution();
        int[][] matrix = {{1,2,3},{4,5,6},{7,8,9}};
        s.rotate(matrix);
        s.print(matrix);
    }

}