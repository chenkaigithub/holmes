package main

func Export(c chan int, holmesConfig *HolmesConfig) {
	c <- 1
}
