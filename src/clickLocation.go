package main

func clickLocation(imgName string, x int, y int, text string, args ...func()) {
	fc := args[0]

	fc()

}
