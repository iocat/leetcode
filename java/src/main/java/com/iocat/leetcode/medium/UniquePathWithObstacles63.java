package com.iocat.leetcode.medium;

public class UniquePathWithObstacles63 {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if (obstacleGrid.length == 0 || obstacleGrid[0].length == 0) {
            return 0;
        }
        int m = obstacleGrid.length, n = obstacleGrid[0].length;

        if (obstacleGrid[0][0] == 1 || obstacleGrid[m - 1][n - 1] == 1) {  // edge cases
            return 0;
        }

        Integer[][] cache = new Integer[m][n];
        return dpCount(obstacleGrid, 0, 0, m, n, cache);
    }

    public int dpCount(int[][] map, int i, int j, int m, int n, Integer[][] cache) {
        if (i > m - 1 || j > n - 1) { // out of bound
            return 0;
        } if (i == m -1 && j == n - 1 && map[i][j] == 0) { // reach the bottom right cell
            return 1;
        }else if (cache[i][j] != null) {    // check the cache
            return cache[i][j];
        }
        int count ;

        if (map[i][j] == 1) { // encounter an obstacle so we can't go any more
            count = 0;
        } else if (i == m - 1) { // can no longer go down
            count = dpCount(map, i, j + 1, m, n, cache);
        } else if (j == n - 1) { // can no longer go right
            count = dpCount(map, i + 1, j, m, n, cache);
        } else {
            count = dpCount(map, i, j + 1, m, n, cache) + dpCount(map, i + 1, j, m, n, cache);
        }
        cache[i][j] = new Integer(count);
        return count;
    }
}
