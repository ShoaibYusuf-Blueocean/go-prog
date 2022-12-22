package service

import( "go-prog/model"
      "fmt")

func Details() {
  
    
    details:= model.Address{}
    fmt.Println(details)
  
    a1 := model.Address{"ram", "Chennai", 3623572}
  
    fmt.Println("Address1: ", a1)
  
    
    a2 := model.Address{Name: "sam", City: "bangalore", Pincode: 277001}
  
    fmt.Println("Address2: ", a2)
  
   
    a3 := model.Address{Name: "Delhi"}
    fmt.Println("Address3: ", a3)
} 
func Person(){
      details:= model.Person{}
	  fmt.Println(details)

	  detail1:= model.Person{"yusuf",234}
      fmt.Println("Person:  ", detail1)
}