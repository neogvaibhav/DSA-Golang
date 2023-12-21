package patterns

import (
	"fmt"
)

func Pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}

func Pattern2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern3(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()

	}
}

func Pattern4(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

func Pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}

func Pattern6(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

func Pattern7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}
}

func Pattern8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Print("*")
		}
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
