package main

import "fmt" 

func main()  {
	const title string = "panjang"
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
	const panjang = 7

	fmt.Printf("* * * %s * * * \n", title) 
	printH(panjang)

}

func printH(length int)  {
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
		
		for j := 0; j < n; j++ {
			strJ := strI
			if j == 0 || j == n-1{
				strJ = "*"
			}
			fmt.Printf("%s %s", strJ, " ")
		}
		fmt.Println()
	}
}