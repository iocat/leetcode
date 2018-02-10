package com.iocat.leetcode.medium;

public class UniquePath62{
        public int uniquePaths(int m, int n){
                return dpUniquePaths(m, n);
        }

        private int dpUniquePaths(int m, int n){
                Integer [][] cache = new Integer[m][n]; // count how many unique ways can reach the end
                int res = countUnique(0, 0, m, n, cache);
                return res ;
        }

        private int countUnique(int i, int j, int m, int n, Integer[][] cache){
                if (i>m-1 && j>n-1){
                        return 0;
                }else if (i == m - 1 && j == n -1 ){
                        return 1;
                }
                if(cache[i][j] != null){
                        return cache[i][j];
                }
                int toCache;
                if (i == m - 1){
                        toCache =  countUnique(i, j+1, m, n ,cache);
                }else if (j == n -1){
                        toCache =  countUnique(i + 1, j, m, n, cache);
                }else {
                        toCache = countUnique(i+1, j, m,n,cache) + countUnique(i, j+1, m, n, cache);
                }
                cache[i][j]=toCache;
                return toCache;

        }

        private int backtrackUniquePaths(int m, int n){
                IntWrapper res = new IntWrapper(0);
                backtrack(res, 0, 0, m, n);
                return res.val;
        }

        private class IntWrapper{
                public IntWrapper(int val){
                        this.val = val;
                }
                int val;
        }

        private void backtrack(IntWrapper count, int i, int j, int m, int n){

                if( i == m - 1 && j == n -1){
                        count.val++;
                        return;
                }
                if(i == m - 1 && j < n) {
                        backtrack(count, i, j + 1, m, n);
                }else if ( j == n -1 && i < m){
                        backtrack(count, i + 1, j, m, n);
                }else{
                        backtrack(count, i + 1, j, m, n);
                        backtrack(count, i, j + 1, m, n);
                }
        }
}