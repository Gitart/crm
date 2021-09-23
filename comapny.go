package main

import (
	"net/http"
	"fmt"
)

// TITLE : Карточка компании
// URl   :  
// DATA  : 
func Company_card(rw http.ResponseWriter, req *http.Request){
   
      // При выборе проекта карточки запись в куки его ид
      Id:= req.URL.Path[len("/company/card/"):]
      Cookies_write(rw, "ProjectId",  Id,  360) 

      // Справочники
      Managers  := Data("Employees")      
      Customers := Data("Companyes")    
      Prjid     := Cookies_read_str(req, "ProjectId")

      Dat:=Mst{"Dt": "Company", "Priority":Priority, "Status":Statuspr, "Manager":Managers, "Customer":Customers, "Percent":Perc, "Projectid":Prjid}
      View_report(rw, req, Dat, "company_card.html", "Карточка компании","Сведения о компании")
} 


// ************************************************************************************************
// Title : Add new company
// URL   :  /company/add/
// ************************************************************************************************
func Company_add(rw http.ResponseWriter, req *http.Request){
	var Dat Company
    req.ParseForm()
	rf      := req.FormValue
 
    pn      := rf("name")
    cn      := rf("Code")
    yp      := Times[0].Id

    if pn == "" {
       return 
    }

    Dat.Title  = pn
    Dat.name   = pn
    Dat.Code   = cn
    Dat.Date   = CTM()
    Dat.Remark = rf("Remark") 


     // Insert table and return id
     Comapnyid:=Insert_Company("Companyes", Dat)
     Cookies_write(rw, "ProjectId",  Comapnyid,  yp)
     
     fmt.Println("Добавлена новая компания ID: "+Comapnyid)

     // Redirect to journal projects
     redirect(rw,req,"/company/card/") 
}
   
   
// ************************************************************************************************
//  Title       : Журнал компаний
//  Description : Компании
//  Url         :  
//  Usage       : 
//  Author      : Savchenko Arthur
//  Create      : 29/07/2019 11:45
//  Changed     : 29/07/2019 11:45
//  Version     : 1.01.01
// ************************************************************************************************
func Companies_view(rw http.ResponseWriter, req *http.Request) {
     View_journal(rw, req, "company.html","Companyes", "Компании","Справочник компаний")
}


// ************************************************************************************************
// Delete all company
// ************************************************************************************************
func Company_delete_all(rw http.ResponseWriter, req *http.Request){
     t:=Dtb{"Main","Companyes"}
     DeleteTable(t)
     fmt.Println("Deleted All Companyes")
}

func Test_struct(rw http.ResponseWriter, req *http.Request){
	var T Сnr 
    T.Code  ="Code 0001"
    T.Title ="Name samp"
    T.Sum   = 456.56

    fmt.Println(T)
    fmt.Printf("%V",T)
    fmt.Printf("%T",T)
    fmt.Printf("%X",T)



}

