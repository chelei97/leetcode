package main

import "fmt"

func main() {
	fmt.Println(robot("URR", [][]int{{4,2}}, 3, 2))
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	u, r := 0, 0
	for i := 0; i < len(command); i ++ {
		if command[i] == 'R' {
			r ++
		}else {
			u ++
		}
	}
	if (x / r - y / u) > 1 || (x / r - y / u) < -1 {
		return false
	}
	for i := 0; i < len(obstacles); i ++ {
		if obstacles[i][0] > x || obstacles[i][1] > y {
			continue
		}
		times := min6(obstacles[i][0] / r, obstacles[i][1] / u)
		lastposionx := obstacles[i][0] - times * r
		lastposiony := obstacles[i][1] - times * u
		startx, starty := 0, 0
		if lastposionx == startx && lastposiony == starty {
			return false
		}
		for i := 0; i < len(command); i ++ {
			if command[i] == 'R' {
				startx ++
				if startx > lastposionx {
					break
				}else if startx == lastposionx {
					if starty == lastposiony {
						return false
					}
				}
			}else {
				starty ++
				if starty > lastposiony {
					break
				}else if starty == lastposiony {
					if startx == lastposionx {
						return false
					}
				}
			}
		}
	}
	times := min6(x / r, y / u)
	lastposionx := x - times * r
	lastposiony := y - times * u
	startx, starty := 0, 0
	if lastposionx == startx && lastposiony == starty {
		return true
	}
	for i := 0; i < len(command); i ++ {
		if command[i] == 'R' {
			startx ++
			if startx > lastposionx {
				return false
			}else if startx == lastposionx {
				if starty == lastposiony {
					return true
				}
			}
		}else {
			starty ++
			if starty > lastposiony {
				return false
			}else if starty == lastposiony {
				if startx == lastposionx {
					return true
				}
			}
		}
	}
	return true
}

func min6(a, b int) int {
	if a < b {
		return a
	}
	return b
}