package main

func CollidingAsteroids(asteroids []int) []int {
	// Write your code here.
	stack := []int{}
	for _, val := range asteroids {
		if val > 0 || len(stack) == 0 {
			stack = append(stack, val)
		} else {
			// if -ve coming
			// stack not empty, top is positive, top smaller than curr
			// keep popping stack
			for len(stack) != 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < abs(val) {
				stack = stack[:len(stack)-1]
			}

			// check equal or top is -ve
			// if equal pop (destroy)
			if len(stack) != 0 && stack[len(stack)-1] == abs(val) {
				stack = stack[:len(stack)-1]
			} else {
				// if top is -ve, just append
				if len(stack) == 0 || stack[len(stack)-1] < 0 {
					stack = append(stack, val)
				}
			}
		}
	}
	return stack
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// input := []int{-3, 5, -8, 6, 7, -4, -7}
// 	expected := []int{-3, -8, 6}
// 3 move left, no collisions
// 5 move right, colliding with -8 and destroyed
// 7 collides with -4, then 7 collide with -7 and vanished
