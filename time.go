package main
import ("fmt"
        "time")
func main() {

	t1 := time.Now()
	t2 := t1.Add(60)
	fmt.Println("Time.Now():", t1)

	fmt.Println(t2.Format(time.RFC822))   // 08 Feb 17 22:28 EST
	fmt.Println(t2.Format(time.RFC822Z))  // 08 Feb 17 22:28 -0500
	fmt.Println(t2.Format(time.RFC850))   // Wednesday, 08-Feb-17 22:29:25 EST
	fmt.Println(t2.Format(time.RFC1123))  // Wed, 08 Feb 2017 22:29:44 EST
	fmt.Println(t2.Format(time.RFC1123Z)) // Wed, 08 Feb 2017 22:30:04 -0500
	fmt.Println(t2.Format(time.RFC3339))  // 2017-02-08T22:30:37-05:00
	fmt.Println(t2.Format(time.RFC3339Nano)) // 2017-02-08T22:31:09.00749996-05:00
	fmt.Println(t2.Format(time.ANSIC))
	time.Sleep(time.Second * 5)
	fmt.Println("time elapsed: ", time.Since(t1))
	time.Unix(1452512222555,22)

}
