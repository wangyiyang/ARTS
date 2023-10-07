from typing import List

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # Check rows
        for row in board:
            if not self.isValid(row):
                return False
        
        # Check columns
        for col in zip(*board):
            if not self.isValid(col):
                return False
        
        # Check boxes
        for i in range(3):
            for j in range(3):
                box = [board[x][y] for x in range(i*3, i*3+3) for y in range(j*3, j*3+3)]
                if not self.isValid(box):
                    return False
        
        return True

    def isValid(self, nums: List[str]) -> bool:
        nums = [x for x in nums if x != '.']
        return len(nums) == len(set(nums))

if __name__ == "__main__":
    board = [["5","3",".",".","7",".",".",".","."],
             ["6",".",".","1","9","5",".",".","."],
             [".","9","8",".",".",".",".","6","."],
             ["8",".",".",".","6",".",".",".","3"],
             ["4",".",".","8",".","3",".",".","1"],
             ["7",".",".",".","2",".",".",".","6"],
             [".","6",".",".",".",".","2","8","."],
             [".",".",".","4","1","9",".",".","5"],
             [".",".",".",".","8",".",".","7","9"]]
    print(Solution().isValidSudoku(board)) # True

    board = [["8","3",".",".","7",".",".",".","."],
             ["6",".",".","1","9","5",".",".","."],
             [".","9","8",".",".",".",".","6","."],
             ["8",".",".",".","6",".",".",".","3"],
             ["4",".",".","8",".","3",".",".","1"],
             ["7",".",".",".","2",".",".",".","6"],
             [".","6",".",".",".",".","2","8","."],
             [".",".",".","4","1","9",".",".","5"],
             [".",".",".",".","8",".",".","7","9"]]
    print(Solution().isValidSudoku(board)) # False
