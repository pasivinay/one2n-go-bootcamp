package iterator

func Iterator(a string, n int) string {
    iterate := a
    for i:=1;i<n;i++ {
        iterate += a
    }

    return iterate
} 