func addDigits(num int) int {
    
    for num >= 10 {
        numToString := strconv.Itoa(num)
        var result int
        for _, c := range numToString {
        
        number, err := strconv.Atoi(string(c))
        if err != nil {
            fmt.Println(err)
        }
            result += number
    }

    num = result
    }

    return num

}
    

    

