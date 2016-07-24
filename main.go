package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Author : HackerZ
// Time   : 2016/07/24

/* timer1
 * Analog logic service processing 1
 */
func timer1() {
	fmt.Println("Timer --> 1")
}

/* timer2
 * Analog logic service processing 2
 */
func timer2() {
	fmt.Println("Timer --> 2")
}

/* handleGetRes
 * Handle Get Response.
 */
func handleGetRes(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Someone In!")
	io.WriteString(rw, "hello, you are in!")
}

/* MAIN FUNCTION
 * AUTHOR HackerZ
 */
func main() {
	t1 := time.NewTimer(2 * time.Second)
	t2 := time.NewTimer(4 * time.Second)

	go func() {
		for {
			select {
			case <-t1.C:
				timer1()
				t1.Reset(2 * time.Second)
			case <-t2.C:
				timer2()
				t2.Reset(4 * time.Second)
			}
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleGetRes)
	mux.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, req *http.Request) {})
	http.ListenAndServe(":9000", mux)

}
