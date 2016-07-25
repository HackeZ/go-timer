package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	utils "go-timer/utils"
)

// Author : HackerZ
// Time   : 2016/07/24

var (
	TimeCount int64
	wg        sync.WaitGroup
)

/* timer1
 * Analog logic service processing 1
 */
func timer1() {
	TimeCount++
	fmt.Printf("Timer --> %d\n", TimeCount)
	wg.Done()
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
	TCString := "hello, you are in!TimeCount --> " + utils.Int64ToStr(TimeCount)
	io.WriteString(rw, TCString)
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
				wg.Add(1)
				timer1()
				t1.Reset(2 * time.Second)
			case <-t2.C:
				timer2()
				t2.Reset(4 * time.Second)
			}
		}
	}()

	wg.Wait()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleGetRes)
	mux.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, req *http.Request) {})
	http.ListenAndServe(":9000", mux)
}
