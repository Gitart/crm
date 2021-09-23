package main

import (
    "net/http"
    "time"
    "log"
    "fmt"
    "html/template"
  r "github.com/dancannon/gorethink"
)

// ************************************************************************************************
// Карточка фазы
// URL : "/phase/card/"
// ************************************************************************************************
func Phase_card(rw http.ResponseWriter, req *http.Request) {

  // Текущая задача 
  Id:= req.URL.Path[len("/phase/card/"):]
  Cookies_write(rw, "TaskId",  Id,  360) 

  // Init data
  var Data  Mst
  
  // Ид проекта
  var Prjid string

  // Employee
  // Rk, err := r.DB("Main").Table("Employees").Filter(Mst{"Sys_flag":P}).OrderBy("Name").Run(sessionArray[0])
  Rk, err := r.DB("Main").Table("Phases").OrderBy("name").Run(sessionArray[0])
  Err(err, "Error read table.")
  defer Rk.Close()
  Rk.One(&Data)

  // // Проекты
  // Prj, errp := r.DB("Main").Table("Projects").OrderBy("name").Run(sessionArray[0])
  // Err(errp, "Error read table.")
  // defer Prj.Close()
  // Prj.All(&Prjs)

  Prjid    = Cookies_read_str(req, "ProjectId")
  nprj    := ProjectName("Projects", Prjid, "Name")
  
  // Make Data
  Dts       := Mst{
                   "Dat"         :   "Карточка задачи", 
                   "Title"       :   "Задачи ", 
                   "Descript"    :   "Журнал задач ", 
                   "ProjectId"   :    Prjid,
                   "ProjectName" :    nprj,
                   "Status"      :    Statuspr,                 
                   "Proority"    :    Priority,                
                   "User"        :    Data,                        
                   "Phases"      :    Phases,                      // Фазы проекта
                   "Type"        :    TypeTask,
                   "Datrep"      :    CTM(),
                 }

  tmpl, err := template.ParseFiles("tmp/phase_card.html", "tmp/main.html")   
  Err(err, "Error template execute.")

  erf       := tmpl.Execute(rw, Dts)
  Err(erf,  "Error template execute.")
}


// ************************************************************************************************
// Журнал фаз
// URL : "/phase/journal/"
// ************************************************************************************************
func Phase_journal(rw http.ResponseWriter, req *http.Request) {

  // Текущая фаза проекта 
  // Id:= req.URL.Path[len("/phase/journal/"):]
  // Cookies_write(rw, "PhaseId",  Id,  360) 
  // fmt.Println("---------", Id)

  // Init data
  var Data  []Mst
  
  // Ид проекта
  var Prjid string

  Prjid    = Cookies_read_str(req, "ProjectId")
  nprj    := ProjectName("Projects", Prjid, "Name")

  // Phase
  Rk, err := r.DB("Main").Table("Phases").GetAllByIndex("ProjectId", Prjid).Run(sessionArray[0])
  Err(err, "Error read table.")
  defer Rk.Close()
  Rk.All(&Data)

  // Make Data
  Dts       := Mst{
                   "Dat"         :   Data, 
                   "Title"       :   "Задачи ", 
                   "Descript"    :   "Журнал задач ", 
                   "ProjectId"   :    Prjid,
                   "ProjectName" :    nprj,
                   "Status"      :    Statuspr,                 
                   "Proority"    :    Priority,                
                   // "User"        :    ,                        
                   // "Phases"      :    ,                      // Фазы проекта
                   "Type"        :    TypeTask,
                   "Datrep"      :    CTM(),
                 }

  tmpl, err := template.ParseFiles("tmp/phases.html", "tmp/main.html")   
  Err(err, "Error template execute.")

  erf       := tmpl.Execute(rw, Dts)
  Err(erf,  "Error template execute.")
}

// ************************************************************************************************
// Добавление фазы проекта
// URL : "Phase_delete_all"
// ************************************************************************************************
func Phase_add(rw http.ResponseWriter, req *http.Request) {

     var P Phase
     rf          := req.FormValue
     tm          := time.Now().Format("02-01-2006 15:04:05")
     Prjid       := Cookies_read_str(req, "ProjectId")

     P.ProjectId  = Prjid                 // Ид проекта
     P.Name       = rf("Name")            // Наименование фазы проекта  
     P.Start      = rf("Start")           // Начало
     P.Finish     = rf("Finish")          // Окончание 
     P.Status     = rf("Status")          // Статус (P-План, W-Работа, D-done Выпролнено)
     P.Percent    = Sti(rf("Percent"))    // Процент выполнения
     P.Sum        = Stf(rf("Sum"))        // Сумма проекта
     P.Created    = tm                    // Дата создания  

    PhaseInsert(P)
    redirect(rw,req,"/phase/journal/" + Prjid) 
}


// ************************************************************************************************
// Добавление задачи в базу
// ************************************************************************************************
func PhaseInsert(Dat Phase) {
      err:=r.DB("Main").Table("Phases").Insert(Dat).Exec(sessionArray[0])
      if err!=nil{
         fmt.Println("Error Added phase")
         return
      }
      log.Println("Insert new phase")
}

// ************************************************************************************************
// Удаление всех фаз проекта
// URL : "Phase_delete_all"
// ************************************************************************************************
func Phase_delete_all(rw http.ResponseWriter, req *http.Request) {}

// ************************************************************************************************
// Удаление одной фаз проекта
// URL : "Phase_delete_one"
// ************************************************************************************************
func Phase_delete_one(rw http.ResponseWriter, req *http.Request) {}

// ************************************************************************************************
// Корректировка крточки фаз проекта
// URL : "/phase/edit/"
// ************************************************************************************************
func Phase_edit(rw http.ResponseWriter, req *http.Request) {

}
 
 