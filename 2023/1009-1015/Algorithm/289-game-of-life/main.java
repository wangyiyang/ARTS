class Solution {
    public void gameOfLift(int[][] board) {
        int m = board.length;
        int n = board[0].length;
        int[][] next = new int[m][n];
        for(int i = 0; i < m; i++) {
            for(int j = 0; j < n; j++){
                int live = 0;
                for(int x = Math.max(i - 1, 0); x <= Math.min(i + 1, m - 1); x++) {
                    for(int y = Math.max(j - 1, 0); y <= Math.min(j + 1, n - 1); y++) {
                        if(x == i && y == j) continue;
                        if(board[x][y] == 1) live++;
                    }
                }
                if(board[i][j] == 1 && (live == 2 || live == 3)) next[i][j] = 1;
                else if(board[i][j] == 0 && live == 3) next[i][j] = 1;
            }
        }
        for(int i = 0; i < m; i++) {
            for(int j = 0; j < n; j++){
                board[i][j] = next[i][j];
            }
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] board = {{0,1,0},{0,0,1},{1,1,1},{0,0,0}};
        s.gameOfLift(board);
        for(int i = 0; i < board.length; i++) {
            for(int j = 0; j < board[0].length;j++) {
                System.out.print(board[i][j] + " ");
            }
            System.out.println();
        }
    }
}
