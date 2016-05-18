package main

// 流程演示。
func demo(x int) {
	// 上下文对象。
	// x := -1

	// 逻辑1
	if x > 0 {
		println("1")
		x++
	} else {
		println("error")
	}

	// 逻辑2
	if x == 2 {
		println("2")
		x++
	}

	// 逻辑3
	if x == 3 {
		println("3")
		return
	}

	// 逻辑结束。
	println("over.")
}

func demo3(x int) {
	if x > 0 {
		x = logic1(x)
		
		if x == 2 {
			x = logic2(x)
		}

		if x == 3 {
			logic3(x)
			return
		}
	} else {
		println("error")
	}

	
	println("over.")
}

func logic1(x int) int {
	println(x)
	x++
	return x
}

func logic2(x int) int{
	println(x)
	x++
	return x
}

func logic3(x int) {
	println(x)
}

func main() {
	x := -1
	demo(x)
	println("-------")
	demo3(x)
}
