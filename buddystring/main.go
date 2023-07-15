func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    srunes := []rune(s)
    grunes := []rune(goal)
    sruneslen := len([]rune(s))
    if s == goal {
        for i:=0;i<sruneslen-1;i++ {
            for j:=i+1;j<sruneslen;j++ {
                if srunes[i] == srunes[j] {
                    return true
                }
            }
        }
        return false
    }

    prev := srunes[0]
    for i:=0;i<len(srunes);i++ {
        if i == 0 {
            continue
        }
        if srunes[i] != grunes[i] || srunes[i] != prev {
            srunes = srunes[i-1:]
            grunes = grunes[i-1:]
            break
        }
        prev = srunes[i]
    }

    prev = srunes[len(srunes)-1]
    for i:=len(srunes)-1;i>=0;i-- {
        if i == len(srunes)-1 {
            continue
        }
        if srunes[i] != grunes[i] || srunes[i] != prev {
            srunes = srunes[:i+2]
            grunes = grunes[:i+2]
            break
        }
        prev = srunes[i]
    }

    sruneslen = len(srunes)
    for i:=0;i<sruneslen-1;i++ {
        for j:=i+1;j<sruneslen;j++ {
            if srunes[i] != grunes[j] || srunes[j] != grunes[i] {
                continue
            }
            if string(srunes[:i]) != string(grunes[:i]) {
                continue
            }
            if string(srunes[j+1:]) != string(grunes[j+1:]) {
                continue
            }
            if string(srunes[i+1:j]) != string(grunes[i+1:j]) {
                continue
            }
            return true
        }
    }
    return false
}
