func fact(int n)  int {
    if(n == 0) {
        return 1
    }
    return fact(n-1) * n
}
func main() void {
    return fact(5)
}