package recursion

func fib(n int) int {
	a,b :=0,1
	sum := 0
	for i := 0;i < n;i++ {
		sum = (a +b) % 1000000007
		a = b
		b = sum
	}
	 return a
 }

 func fib2 (n int) int {
	 if n <= 1 {
		 return n
	 }
	 return fib2(n -1) + fib(n-2)
 }