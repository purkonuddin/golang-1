package main

import "fmt"  

type segitiga struct  {
	length int
	width int
	heigh int
} 

func main()  {
	const title string = "panjang"
	
	const panjang = 7
	/*
	fmt.Printf("* * * %s * * * \n", title) 
	printH(panjang)
	*/
	fmt.Println()

	fmt.Print("* * * segi tiga * * * \n")  

	var bangunDatar interface{} = &segitiga{length: panjang}
	n := bangunDatar.(*segitiga).length
	rataKanan(n)
	fmt.Println()
	rataKiri(n) 
}

func rataKanan(n int)  {

	/*
				* 	[0,0] [0,1] [0,2] [0,3] *	4 = 5 - 1 
			  * *   [1,0] [1,1] [1,2] *		*	3
			* * *   [2,0] [2,1] *	  *		*	2
		  * * * *   [3,0] *		*	  *		*	1
		* * * * *	*	  *		*	  *		*	0
	*/

	for i := 1; i <= n; i++ {
		x := n - i
		for j := 0; j < n; j++ {
			str := "*" 
			if x > j {
				str = " "
			}
			fmt.Print(str, " ")
		}
		fmt.Println()
	}
}

func rataKiri(n int)  {

	/*
		* * * * *
		* * * * 
		* * * 
		* * 
		* 
	*/

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}
}

func printH(length int)  {

	/*
		* mencetak gambar H, 
		* panjang harus ganjil
	
		** panjang **
		*  =  =  =  *
		*  =  =  =  *
		*  *  *  *  *
		*  =  =  =  *
		*  =  =  =  *
	
	*/

	n := length

	if n %2 == 0 {
		fmt.Print(" \"panjang\" harus ganjil \n")
		return
	}  

	for i := 0; i < n; i++ {
		strI := "="
		if i == n/2 {
			strI = "*"
		}
		
		for j := 1; j <= n; j++ {
			strJ := strI
			if j == 1 || j == n {
				strJ = "*"
			}
			fmt.Printf("%s %s", strJ, " ")
		}
		fmt.Println()
	}
}