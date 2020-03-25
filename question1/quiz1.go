package main
import "fmt"

func main() {
	sum := 0
	var num int
	str2 := "Exit"
	str1 := "Please select an option:"
	str3 :=	"1 – Print Covid19 cases in Pakistan" 
	str4 :=	"2 – Print Covid19 cases in SouthKorea"
	str5 :=	"3 – Print Covid19 cases in France"
	str6 :=	"4 – Print my personalized message to Coronavirus"
	str7 :=	"0 – Exit"
	
	fmt.Println(str1)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	fmt.Scanf("%d", &num)
	
	for sum <= 4 {
		switch num {
		case 1:
			fmt.Println("Covid19 cases in Pakistan are: ")
			fmt.Println("1000")
			
		case 2:
			fmt.Println("Covid19 cases in SouthKorea are: ")
			fmt.Println("9137")

		case 3:
			fmt.Println("Covid19 cases in France are: ")
			fmt.Println("22,304")
		
		case 4:
			fmt.Println("Covid19 is a big problem for us and we need to battle it as a team and stay at home to stay safe")
		
		default:
			fmt.Printf(str2)
		}
		sum++
	}

}
