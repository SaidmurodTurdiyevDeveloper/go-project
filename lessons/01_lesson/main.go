package main

func main() {

	//1-masala
	// var a, P int
	// fmt.Print("kvadrat tomoni kiritilsin a=")
	// fmt.Scan(&a)
	// P = a * 4
	// fmt.Println("Peremetre = ", P)

	//2-masala
	// var a, S int
	// fmt.Println("kvadreat tomoni kiritilsin")
	// fmt.Scan(&a)
	// S = int(math.Pow(float64(a), 2))
	// fmt.Print(S)

	/*
	 fmt.Scnaf() metodini ishlatish uchun qo'llanma
	 integer %d
	 float64 %g
	 string %s
	 bool %t
	*/

	// 3- masala
	// var a, b, S, P int
	// fmt.Scanf("%d %d", &a, &b)
	// S = a * b
	// P = (a + b) * 2
	// fmt.Println("S =", S, "\nP =", P)

	/*
	 fmt.Prinf() metodini ishlatish uchun qo'llanma
	 o`zgaruvchi typini` %T
	 o`zgaruvchi qiymatini %v
	*/

	//4-masal
	// var (
	// 	d  int
	// 	pi float64 = 3.14
	// 	S  float64
	// )
	// fmt.Println("aylana dimaetri kiritilsin")
	// fmt.Scan(&d)
	// S = pi * float64(d)
	// fmt.Printf("Naitija %T da hisoblandi \nAylana uzunligi %v", S, S)

	//5-masala
	// var (
	// 	a int
	// 	V int
	// 	S int
	// )
	// fmt.Println("Kubning tomoni kiritilsin")
	// fmt.Scan(&a)
	// V = int(math.Pow(float64(a), 3))
	// S = int(6 * math.Pow(float64(a), 2))
	// fmt.Printf("Kubning hajmi %v\nKubning yuzasi %v", V, S)

	//6- masala
	// var (
	// 	a int
	// 	b int
	// 	c int
	// )
	// var (
	// 	V int
	// 	S int
	// )
	// fmt.Scanf("%d %d %d", &a, &b, &c)
	// V = a * b * c
	// S = 2 * (a*b + b*c + a*c)
	// fmt.Printf("Kubning hajmi = %v\nKubning yuzi S=%v", V, S)

	//7-masala
	// var (
	// 	r  int
	// 	pi float64 = 3.14
	// 	L  float64
	// 	S  float64
	// )
	// fmt.Scan(&r)
	// L = 2 * pi * float64(r)
	// S = pi * math.Pow(float64(r), 2)
	// fmt.Printf("Aylananing uzunligi L=%v\nAylananing yuzi S=%v", L, S)

	//8 masala
	// var (
	// 	a float64
	// 	b float64
	// 	s float64
	// )
	// fmt.Scanf("%g %g", &a, &b)
	// s = (a + b) / 2
	// fmt.Printf("O`rta arifmetik qiymati %v", s)

	//9-masala
	// var (
	// 	a float64
	// 	b float64
	// 	s float64
	// )
	// fmt.Scanf("%g %g", &a, &b)
	// s = math.Sqrt(math.Abs(a * b))
	// fmt.Printf("O`rta geometrik qiymati %v", s)

	//10-masala
	// var (
	// 	a int
	// 	b int
	// )
	// fmt.Println("Ikkita 0 ga teng bo`lmagan qiymat iriting")
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// if a != 0 && b != 0 {
	// 	fmt.Println("Yig`indisi ", a+b)
	// 	fmt.Println("Ko`paytmasi ", a*b)
	// 	fmt.Println("Kvadratlari\na=", a, "kvadrati ", math.Pow(float64(a), 2))
	// 	fmt.Println("b=", b, "kvadrati ", math.Pow(float64(b), 2))
	// } else {
	// 	fmt.Print("Error a=", a, " b=", b)
	// }

	//11-masala
	// var (
	// 	a int
	// 	b int
	// )
	// fmt.Println("Ikkita 0 ga teng bo`lmagan qiymat iriting")
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// if a != 0 && b != 0 {
	// 	fmt.Println("Yig`indisi ", a+b)
	// 	fmt.Println("Ko`paytmasi ", a*b)
	// 	fmt.Println("Modullari\na=", a, "moduli ", math.Abs(float64(a)))
	// 	fmt.Println("b=", b, "moduli ", math.Abs(float64(b)))
	// } else {
	// 	fmt.Print("Error a=", a, " b=", b)
	// }

	//12-masala
	// var (
	// 	a float64
	// 	b float64
	// 	c float64
	// 	P float64
	// )
	// fmt.Println("Uchburchak katetlarini kiriting")
	// fmt.Print("a=")
	// fmt.Scan(&a)
	// fmt.Print("b=")
	// fmt.Scan(&b)

	// c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	// P = a + b + c
	// fmt.Println("Uchburchak gipetenuzasi c =", c)
	// fmt.Println("Uchburchak peremetri P =", P)

	//13-masala
	// var (
	// 	R1 float64
	// 	R2 float64
	// 	S1 float64
	// 	S2 float64
	// 	S3 float64
	// )
	// fmt.Println("Aylanalar radiusini kiriting")
	// fmt.Print("R1=")
	// fmt.Scan(&R1)
	// fmt.Print("R2=")
	// fmt.Scan(&R2)

	// S1 = math.Pi * math.Pow(R1, 2)
	// S2 = math.Pi * math.Pow(R2, 2)
	// if R1 > R2 {
	// 	S3 = S1 - S2
	// } else {
	// 	S3 = S2 - S1
	// }
	// fmt.Print("Yuzasi S3=", S3)

	//Ferangate change celsius tempereture

	// var input = bufio.NewReader(os.Stdin)

	// inputText, errorText := input.ReadString('\n')
	// if errorText != nil {
	// 	fmt.Println(errorText)
	// 	os.Exit(1)
	// }
	// tempereture, errorTempereture := strconv.Atoi(strings.TrimSpace(inputText))
	// if errorTempereture != nil {
	// 	fmt.Println(errorTempereture)
	// 	os.Exit(1)
	// }

	// celsies := (tempereture - 32) * 5 / 9

	// fmt.Println(tempereture, "Fahrenheit equals with", celsies, "Celsius")

	//we can find this is which type of operating system
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)
	// }

	// Fizz - Buzz

	/* Write a program that prints the numbers from 1 to 100.
	   But for multiples of three print "Fizz" instead of the number and
	   for the multiples of five print "Buzz".
	   For numbers which are multiples of both three and five print "FizzBuzz".
	*/

	// for i := 0; i < 100; i++ {

	// 	switch {
	// 	case i%5 == 0 && i%3 == 0:
	// 		fmt.Println("FizzBuzz")
	// 	case i%3 == 0:
	// 		fmt.Println("Fizz")
	// 	case i%5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }
	/*
		const text string = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

		fmt.Printf("quated string: %+q", text)

	*/

}
