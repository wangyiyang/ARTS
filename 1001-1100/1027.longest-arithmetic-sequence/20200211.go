func longestArithSeqLength(A []int) int {
    lenth := len(A)
    
    index := [10001][]int{}
    for i:=0;i<lenth;i++ {
        index[A[i]] = append(index[A[i]], i)
    }
    
    var dfs func(int, int)int
    dfs = func(j, diff int)int{
        next:= A[j]+diff
        if next<0 || next>10000 {
            return 0
        }
        for _, k:= range index[next]{
            if j<k {
                return 1+dfs(k, diff)
            }
        }
        return 0
    }
    
    ret := 0
    for i:=0;i<lenth;i++{
        for j:=i+1;j<lenth;j++{
            diff := A[j]-A[i]
            ret = max(ret, 2+dfs(j, diff))
        }
    }
    return ret
    
}

func max(a, b int)int {
    if a>b {
        return a
    }
    return b
}

