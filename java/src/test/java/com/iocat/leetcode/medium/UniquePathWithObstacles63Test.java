package com.iocat.leetcode.medium;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;

import java.util.Arrays;
import java.util.Collection;
import static org.junit.Assert.assertEquals;
@RunWith(Parameterized.class)
public class UniquePathWithObstacles63Test {
    int[][] obstacleMap;
    int res;

    public UniquePathWithObstacles63Test(int[][] map, int res) {
        this.obstacleMap = map;
        this.res = res;
    }

    @Parameterized.Parameters
    public static Collection<Object[]> data() {
        return Arrays.asList(new Object[][]{
                {
                        new int[][]{
                                {0, 0, 0},
                                {0, 1, 0},
                                {0, 0, 0}}, 2
                },
                {
                        new int[][]{
                                {1, 1},
                                {0, 0},
                        }, 0
                },
                {
                        new int[][]{
                                {0, 0, 0, 0},
                                {0, 0, 0, 1},
                        }, 0
                },
                {
                        new int[][]{
                                {0, 0, 1, 0},
                                {1, 0, 0, 0},
                        }, 1
                },
                {
                        new int[][]{
                                {1},
                        }, 0
                },
                {
                        new int[][]{
                                {0, 0, 0},
                                {0, 0, 0},
                                {0, 0, 0},
                        }, 6
                }
        });
    }

    @Test
    public void test(){
        UniquePathWithObstacles63 upwo = new UniquePathWithObstacles63();
        assertEquals("Result",this.res,upwo.uniquePathsWithObstacles(this.obstacleMap));
    }

}
