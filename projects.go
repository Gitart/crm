package main

import (
    "net/http"
    "log"
    "fmt"
    "time"
  r "github.com/dancannon/gorethink"
)

// ************************************************************************************************
//  Title       : Журнал проектов
//  Description : Проекты открытые в компании
//  Url         :  
//  Usage       : 
//  Author      : Savchenko Arthur
//  Create      : 15/07/2019
//  Changed     : 24/07/2019
//  Version     : 2.01.01
// ************************************************************************************************
func Projects_view(rw http.ResponseWriter, req *http.Request) {
     View_journal(rw, req, "projects.html", "Projects", "Проекты", "Текущие проекты компании")
}


// ************************************************************************************************
// Карточка о проекте
// Url:  http://localhost:1234/project/card/
// ************************************************************************************************
func Project_card(rw http.ResponseWriter, req *http.Request){
   
      // При выборе проекта карточки запись в куки его ид
      Id:= req.URL.Path[len("/project/card/"):]
      Cookies_write(rw, "ProjectId",  Id,  360) 

      // Справочники
      Managers  := Data("Employees")      
      Customers := Data("Companyes")    
      Prjid     := Cookies_read_str(req, "ProjectId")

      Dat:=Mst{"Dt":"Project", "Priority":Priority, "Status":Statuspr, "Manager":Managers, "Customer":Customers, "Percent":Perc, "Projectid":Prjid}
      View_report(rw, req, Dat, "project_card.html", "Карточка проекта","Сведения о проекте")
} 


// todo num - по порядку номер 
// Note - описание проекта 
// *******************************************************
// Добавление нового проекта
// *******************************************************
func Project_add(rw http.ResponseWriter, req *http.Request){
      
rf      := req.FormValue
tm      := time.Now().Format
dt      := tm("02-01-2006 15:04:05")
pn      := rf("ProjectName")
cn      := rf("Code")

if pn == "" {
   return 
}

// fmt.Println("Percent", rf("Percent"))
Dat:= ProjectData {
            Description:   rf("Description"),
            Name:          pn,
            Code:          cn,
            Percent:       Sti(rf("Percent")),
            Manager:       rf("Manager"),
            Key:           NowUnixTime(),  
            Id:            Keygen(5, "alphanum"),   
            Customer :     rf("Customer"),
            Status:        Statuspr[0].Id,
            DeadLine:      rf("DateEnd"), //CTUS("2006-01-02") 
            Active:        1,
            Datetime:      dt,
            DateStart:     rf("DateStart"),   
            DateEnd:       rf("DateEnd"),   
      }
     
     // Year
     yp:=Times[0].Id
     
     // In time create project write cookies to local 
     // Name project to cookie
     Cookies_write(rw, "project", pn,  yp)    

     // Code project to cookie 
     Cookies_write(rw, "code",    cn ,  yp) 

     // Insert table and return id
     Prjid:=Insert_Table("Projects", Dat)
     Cookies_write(rw, "ProjectId",  Prjid,  yp)

     // Redirect to journal projects
     redirect(rw,req,"/project/card/") 
} 
 
// Удаление всех проектов физически из таблицы
// Операция позволительно только администратору
// Или владельцу проекта тому кто создал проект
// Обязательная проверка роли перед операцией удаления
func Project_delete(rw http.ResponseWriter, req *http.Request){

    if req.Method == "DELETE"{ 
       err:=r.DB("Main").Table("Projects").Delete().Exec(sessionArray[0])
     if err!=nil{
      fmt.Println("Project error for delete")
     }else{
      fmt.Println("All Projects Was Been Deleted Success.")
     }
   }else{
    fmt.Println("Wrong method")
   }
}
 
// Удаление проекта физически из таблицы
// Операция позволительно только администратору
// Или владельцу проекта тому кто создал проект
// Обязательная проверка роли перед операцией удаления
// Url : /project/delete/{:guid}
func Project_delete_one(rw http.ResponseWriter, req *http.Request){
     Id:=1
     r.DB("Main").Table("Projects").Filter(Mst{"id":Id}).Delete().Exec(sessionArray[0])
}


// ************************************************************************************************
// Проекты лист
// http.HandleFunc("/proj/list/",  Project_journal )     // Журнал проектов
// ************************************************************************************************
func Project_journal(rw http.ResponseWriter, req *http.Request) {

  var Prjs []Mst        // Projects
  var Samp []Mst        // Accumulative 
  var Samp_tasks Mst    
  
  // Get all projects
  Prj, errp := r.DB("Main").Table("Projects").Filter(Mst{"Active":1}).OrderBy("id").Run(sessionArray[0])
  Err(errp, "Error read table.")
  defer Prj.Close()
  Prj.All(&Prjs)


// Обнуление переменной
// Цикл по проектам компании и наполнение задачами для отчета
for _, prt := range Prjs {
 
       Pid := prt["id"].(string)
       Nam := prt["customer"].(string)
       Des := prt["Description"].(string)
       Tsk := Datdb(Pid)

       // Created tasks & projects
       Samp_tasks=Mst{"Title": Nam,  "Description":Des, "Id":Pid, "Tasks": Tsk}
       
       // Добавление в задачи по проекту
       Samp=append(Samp, Samp_tasks)
  }

  Dts       := Mst{
                   "Dt":          Samp,
                   "Projects":    Prjs,
                   "Title":       "Отчет о статусе проектов ", 
                   "Descript":    "Журнал задач в разрезе проектов", 
                   "Datrep":      CTM(),
                 }

   View_report(rw,req, Dts, "projects_issues.html" , "Проекты","Проектный лист" )
}



// *******************************************************************
// Обновление итоговых часов по текущему проекту
// /project/sum/
// *******************************************************************
func Project_sum(rw http.ResponseWriter, req *http.Request){

     go Hours_sum("dd3d6287-c5f9-4ae9-b3b5-10c2332efb49")
}


// Подсчет общее количенстов часов по проекту по всем задачам
// dd3d6287-c5f9-4ae9-b3b5-10c2332efb49
func Hours_sum(IDProject string){

     var response []Mst
       Prj           := make(Mst)


     res, _ := r.DB("Main").Table("Tasks").
            GetAllByIndex("Project_id", IDProject).
            Group("Project_id").
            Map(Mst{"Hp":   r.Row.Field("HoursPlan"),
                    "Ho":   r.Row.Field("HoursOver"),
                    "Hf":   r.Row.Field("HoursFact"),
                  }).
            Ungroup().
            Map(Mst{ "Name":     "Summ Hours by Project : " + IDProject,
                     "Grphour":   r.Row.Field("group"),                                 // группировка
                     "Hpr":       r.Row.Field("reduction").Field("Hp").Sum().Default(0),           // Количество медикаментов в на складе после продажи
                     "Hor":       r.Row.Field("reduction").Field("Ho").Sum().Default(0),           // сумма
                     "Hfr":       r.Row.Field("reduction").Field("Hf").Sum().Default(0),
                   }).
            Run(sessionArray[0])

  err := res.All(&response)
  
  // Error
  if err != nil {
     panic(err)
  }

  rr  := response[0]
  hrs := rr["Hfr"].(float64)
  hrd := hrs/8
  Prj["SumHpr"]= rr["Hpr"].(float64)
  Prj["SumHor"]= rr["Hor"].(float64)
  Prj["SumHfr"]= rr["Hfr"].(float64)
  Prj["SumDay"]= hrd
   

  // fmt.Println(rr["Name"],rr["Hpr"], "Часов = (", hrs, " ) Дней =", hrd)
   
  errec:=r.DB("Main").Table("Projects").Get(IDProject).Update(Prj).Exec(sessionArray[0])
  if errec!=nil{
     log.Println("Error Update Project Hours", errec.Error())
     return
  }
}


// **********************************************************************************
// Возвращает поле по ид из выбранной таблицы
// **********************************************************************************
func ProjectName(Tab, Id, Field string) string{
  var NamePr Mst
  Prjn, errp := r.DB("Main").Table(Tab).Get(Id).Run(sessionArray[0])
  
  if errp !=nil{
     return " "
  }

  defer Prjn.Close()
  Prjn.One(&NamePr)
  
  // Если по текущему условию ничего не найдено возвращение пустого значения
  if NamePr==nil{
     return " " 
  }
  return NamePr[Field].(string)
} 






















































// ************************************************************************************************
// Журнал активных проектов
// http.HandleFunc("/proj/list/",  Project_journal )     // Журнал проектов
// ************************************************************************************************
func Project_journal_old_variant(rw http.ResponseWriter, req *http.Request) {

  var Prjs []Mst        // Projects
  var Samp []Mst        // Accumulative 
  var Samp_tasks Mst    
  
  // Приоритет
  pr:=Priority
  
  // Статус
  st:=Statuspr

  //  samp:=[]Mst{ 
  //                Mst{"Title": "Test1", 
  //                    "Tasks": []Mst{
  //                                    Mst{"Name":"N1111"},
  //                                    Mst{"Name":"N1112"},
  //                                    Mst{"Name":"N1113"},
  //                                    Mst{"Name":"N1114"},
  //                                   }, 
  //                    },

  //                  Mst{"Title": "Test2", 
  //                    "Tasks": []Mst{  
  //                                    Mst{"Name":"S-1111"},
  //                                    Mst{"Name":"S-1112"},
  //                                    Mst{"Name":"S-1113"},
  //                                    Mst{"Name":"S-1114"},
  //                                   }, 
  //                    },
  //  }

  //  samp2:= Mst{"Title": "ПВХЗ-456", 
  //                    "Tasks": []Mst{
  //                                    Mst{"Name":"ПВХЗ--N1111"},
  //                                    Mst{"Name":"ПВХЗ--N1112"},
  //                                    Mst{"Name":"ПВХЗ--N1113"},
  //                                    Mst{"Name":"ПВХЗ--N1114"},
  //                                   }, 
  //  }



  // Samp=append(samp,samp2)


  // // 
  // // Rk, err := r.DB("Main").Table("Employees").Filter(Mst{"Sys_flag":P}).OrderBy("Name").Run(sessionArray[0])
  // Rk, err := r.DB("Main").Table("Employees").OrderBy("name").Run(sessionArray[0])
  // Err(err, "Error read table.")
  // defer Rk.Close()
  // Rk.All(&Data)

  Prj, errp := r.DB("Main").Table("Projects").Filter(Mst{"Active":1}).OrderBy("Id").Run(sessionArray[0])
  Err(errp, "Error read table.")
  defer Prj.Close()
  Prj.All(&Prjs)

// Обнуление переменной
// Цикл по проектам компании и наполнение задачами для отчета
for _, prt := range Prjs {
 
       Pid := prt["id"].(string)
       Nam := prt["customer"].(string)
       Des := prt["Description"].(string)
       Tsk := Datdb(Pid)
       // fmt.Println( "N#",Pid, "\n", Nam, " * ", Nam, "\n\n")

       fmt.Println(prt["Description"])

       // fmt.Println(prt["customer"])
       // ipr:=prt["tasks"].([]map[string]interface{})
       // ipr:=prt["tasks"].([]Mst)
       // fmt.Println(ipr)
       // ipr:=prt["tasks"].([]interface{})

        // // Прокрутка задач
        // for _, prtss:=range tsk{
        //     fmt.Println( prtss["Projectid"], "***", prtss["Title"])
        // }     

        // samp_tasks:=map[string]interface{}{"Title": prt["customer"], "Id" :Prjid, "Tasks": tsk}
        Samp_tasks=Mst{"Title": Nam, "Description":Des, "Id":Pid, "Tasks": Tsk}

         // fmt.Println("==============================================================\n")
         // O_Json(Samp_tasks,rw)

        // Добавление задач в массив по проекту
        Samp=append(Samp, Samp_tasks)
  }

 // for _, ttt:=range(Samp){
 //        fmt.Println("---------------------",ttt["Id"])
 //        fmt.Println("-",ttt["Tasks"].([]Mst))

 // }

 // OJson(samp,rw)
 // fmt.Println("Wbrk...", Prjs)

  // Цикл по структуре базы
  // for Prj.Next(&Rep) {
  //     fmt.Println("Rep" )
  //     fmt.Println(Rep )
  // }

 // Ieq:=r.EqJoinOpts{Index:"Projectid"}
 //  Rt,_:=  r.DB("Main").Table("Projects").EqJoin("id", r.DB("Main").Table("Issues"), Ieq).Zip().Run(sessionArray[0])
 //  defer Rt.Close()
 //  Rt.All(&Prt)

  Dts       := Mst{
                   "Dt":          Samp,
                   "Dat":         "Карточка задачи", 
                   "Status":      st,
                   "Proority":    pr,
                   // "User":        Data,
                   "Projects":    Prjs,
                   "Title":       "Задачи ", 
                   "Descript":    "Журнал задач ", 
                   "Datrep":      CTM(),
                 }

   View_report(rw,req, Dts, "projects_issues.html" , "Проекты","Проектный лист" )
}

