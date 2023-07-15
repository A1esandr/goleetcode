func buddyStrings(s string, goal string) bool {
    srunes := []rune(s)
    for i:=0;i<len(srunes)-1;i++ {
        for j:=i+1;j<len(srunes);j++ {
            news := make([]rune, len(srunes))
            for a, sym := range srunes {
                if a == i {
                   news[a] = srunes[j]
                   continue
                }
                if a == j {
                   news[a] = srunes[i]
                   continue
                }
                news[a] = sym
            }
            if string(news) == goal {
                return true
            }
        }
    }
    return false
}
