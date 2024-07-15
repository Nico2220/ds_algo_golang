package dp

func Fibdp(n int, dp []int) int {
	if n <= 1 {
		return n
	}

	dp[n] = Fibdp(n-1, dp) + Fibdp(n-2, dp)
	return dp[n]
}

func Fibdp2(n int, dp []int) int {
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func Fibdp3(n int) int {
	prev2 := 0
	prev1 := 1

	for i := 2; i <= n; i++ {
		curri := prev1 + prev2
		prev2 = prev1
		prev1 = curri
	}

	return prev1
}
