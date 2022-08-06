
package main

import (
    "fmt"
    "testing"
)

func TestMinIntBasic(t *testing.T) {
    result := MinInt(2, -2)

    if result != -2 {
        t.Errorf("MinInt(2, -2) = %d; expected -2", result)
    }
}

func TestMinIntTableDriven(t *testing.T) {
    var tests = []struct{
        a, b     int
        expected int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {3, 2, 2},
        {0, -1, -1},
    }

    for _, test := range tests {

        testname := fmt.Sprintf("%d,%d", test.a, test.b)
        t.Run(testname, func(t *testing.T) {

            result := MinInt(test.a, test.b)

            if result != test.expected {
                t.Errorf("got %d, want %d", result, test.expected)
            }

        })

    }
}

func BenchmarkMinInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MinInt(1, 2)
    }
}
