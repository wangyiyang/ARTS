class Solution {
    public void setZeroes(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        boolean[] row = new boolean[m];
        boolean[] col = new boolean[n];
        for(int i = 0; i < m; i++)
            for(int j = 0; j < n; j++)
                if(matrix[i][j] == 0){
                    row[i] = true;
                    col[j] = true;
                }
        for(int i = 0; i < m; i++)
            if(row[i])
                for(int j = 0; j < n; j++)
                    matrix[i][j] = 0;
        for(int j = 0; j < n; j++)
            if(col[j])
                for(int i = 0; i < m; i++)
                    matrix[i][j] = 0;

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
        int[][] matrix = {{1,1,1},{1,0,1},{1,1,1}};
        s.setZeroes(matrix);
        s.print(matrix);
    }
}