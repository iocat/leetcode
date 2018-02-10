package com.iocat.leetcode.medium;


import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameters;
import static org.junit.Assert.assertEquals;

import java.util.Collection;
import java.util.Arrays;

@RunWith(Parameterized.class)
public class UniquePath62Test{
    public int m;
    public int n;
    public int res;

    public UniquePath62Test(int m, int n, int res){
        this.m = m;
        this.n = n;
        this.res = res;
    }

    @Parameters
    public static Collection<Object[]> data (){
        Object[][] ret = new Object[][]{{1,1,1}, {2,1,1}, {1,2,1},{2,2,2},{1,3,1},{3,1,1},{3,2,3},{2,4,4}};
        return Arrays.asList(ret);
    }
    @Test
    public void testUniquePath(){
        UniquePath62 up = new UniquePath62();
        int result = up.uniquePaths(m, n);
        assertEquals("Result", this.res, result);
    }

}