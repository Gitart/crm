package main

// Customer
type Customer struct {
     FirstName  string                                   // Имя
     LastName   string                                   // Фамилия
     Email      string                                   // Почта рабочая
     EmailExt   string                                   // Почта внешняя 
     Phone      string                                   // Телефон рабочий
     Mob        string                                   // Телефон мобильный
     Position   string                                   // Должность 
     Login      string                                   // Логин
     Password   string                                   // Пароль
     Hash       string                                   // Хеш сумма
     Dateinput  string                                   // Дата ввода
     Role       string                                   // Role
     InOut      bool                                     // External Internal - сотрудник внешний внутернний 
     Company    string                                   // Company id
     Block      bool                                     // Заблокирован
     Remark     string                                   // Примечание 
}

// Фазы этап проекта
type Phase struct {
     ProjectId       string                              // Ид проекта
     Name            string                              // Наименование фазы проекта  
     Description     string                              // Описание 
     Start           string                              // Начало
     Finish          string                              // Окончание 
     Status          string                              // Статус (P-План, W-Работа, D-done Выпролнено)
     Percent         int                                 // Процент выполнения
     Sum             float64                             // Сумма
     Created         string                              // Дата создания 
     Teams         []string                              // Команда проекта
     Days            int                                 // Количество дней на єтап
}

// База - таблица
type Dtb struct {
     Db     string
     Table  string
 }

// Задача
type Task struct {
     Num        string                                    // Номер задачи 
     Project_id string                                    // ID Projects - проект
     Stage_id   string                                    // ID Stage - этапа
     Code       string                                    // Код задачи
     Name       string                                    // Наименование задачи
}

// Связь задач
type Reps struct {
     Name         string  
     Pers         int 
     Datetime     string
     Description  string 
     Tasks        []Task 
}

// Задачи
type Tasks struct {
     id          string
     Id          string                                  // Уникальный код задачи
     Num         string                                  // Номер задачи 
     Project_id  string                                  // ID Projects - проект
     Stage_id    string                                  // ID Stage - этапа
     Phase       string                                  // Фаза выполнения
     Code        string                                  // Код задачи
     Name        string                                  // Наименование задачи короткое
     Title       string                                  // Наименование задачи
     Remark      string                                  // Примечание к задаче 
     Created     string                                  // Дата создания задачи
     Updated     string                                  // Дата обновления задачи
     Start       string                                  // Старт - когда планируется начаться задача
     Finish      string                                  // Финиш - когда планируется завершиться
     Due         string                                  // Когда должна завершиться
     Manager     string                                  // Кто назначил залачу - ответсвенный инженер (менеджер) за задачу (справочник Users)
     Managerid   int64                                   // Кто назначил залачу ID
     Executor    string                                  // Исполнитель
     Executor_id string                                  // Исполнитель ID
     Status      string                                  // Стаутс задачи
     HoursPlan   float64                                 // Расчетные плановые часы по задаче
     HoursFact   float64                                 // Фактически потраченные часы по задаче
     HoursOver   float64                                 // Потраченное сверхурочное время на выполнение задачи в часах
     Hours       float64                                 // Потраченное время на выполнение задачи в часах
     DaysPlan    float64                                 // Запланированно дней
     DaysFact    float64                                 // Фактически потраченно дней
     Percent     int                                     // Процент выполенения
     Difficutly  int                                     // Сложность задачи (от 10 до 1000)
     Type        string                                  // Тип задачи (работа, услуги, консультации)
     Priority    string                                  // Приоритет (высокий, средний, низкий)
     Speed       string                                  // Скорсть (срочность) исполнения (h-высокая, m-средний, l-низкий)
     Sum         float64                                 // Стоимсоть задачи кол. часов * стоимость (роли) в проекте
     Hidden      bool                                    // Скрытие задачи 
     Deleted     bool                                    // Удаленная логически задача 
     Etag        string                                  // Електронный тег
     Kind        string                                  // Ссылки на другие задачи
     Notes       string                                  // Примечание для себя
     Approve     string                                  // Дата подтверждения выполненной задачи
     Act         string                                  // Попадаение в акт выполенных работ - номер акта
     Completed   bool                                    // Выполненна и закрыта задача   
     Link        string                                  // Линк на документ на сервер SharePoint или в облако
     Tags        string                                  // Теги для поиска задачи 
     Parent      string                                  // Родительская задача если она существует  
     Position    int64                                   // Позиция в списке  
     Flag        string                                  // Системный флаг   
     Color       string                                  // Цвет подсветки фона задачи (Red - если просрочена)
     Ico         string                                  // Ico name  
}


// Расчет дней из введенных часов
func (t *Tasks) WorkDays() {
      t.DaysPlan = t.HoursPlan * 8
      t.DaysFact = t.HoursFact * 8
}


// Основные поля для справочника
type Directory struct {
	Id         string                                    // Уникальный код задачи
    Name       string                                    // Наименование задачи
}

// Основные поля для справочника
type Directoryint struct {
    Id         int                                       // Уникальный код задачи
    Name       string                                    // Наименование задачи
}

// Основные поля для справочника
type DirectoryIntStr struct {
    Id         int                                       // Уникальный код задачи
    Code       string                                    // Код позиции в справочнике
    Name       string                                    // Наименование задачи
}

// Код
func (t *DirectoryIntStr) So() string {
     return "code : " + t.Code + "Name : " +t.Name
}

// Перечень дирректорий
type DirTxt struct {
    ddd         DirectoryIntStr
    Num        string
}

// Users Пользователи
type Users struct {
     Id         string                    // Уникальный код
     Name       string                    // Имя пользователя
     Fam        string                    // Фамилия пользователя
     Position   string                    // Должность
     Login      string                    // Логин
     Hash       string                    // Hash
     Email      string                    // Email
     Emails     string                    // Email второй
     Telephone  string                    // Телефон
     Mob        string                    // Мобильный    
     Pass       string                    // Пароль
     Time       string                    // Время регистрации
     Block      bool                      // Блокировка
     Remark     string                    // Примечание
     Type       string                    // Тип (out-внешний, in-внутренний)
     Company    string                    // Компания 
     Depart     string                    // Отдел
     Boss       string                    // Фамилия Имя начальника 
     BossId     string                    // Id начальника 
     Load       int                       // Загрузка 
     Num        string                    // Табельный номер  
}

// Лог
type Logview struct {
     Id         string                    // Уникальный код
     Title      string                    // Операция
     Datetime   string                    // Дата операции
     Module     string                    // Модуль (System, Project, Tasks, User, Login)
}     

// Карточка проекта
type ProjectData   struct {
    Id             string                 // Уникальный код 
    Code           string                 // Код внутренний в компании
    ExtCode        string                 // Код внешний у контаргента если нам известно
    Key            int64                  // Уникальный код 
    Name           string                 // Имя проекта
    Description    string                 // Описание   
    Active         int                    // Активный = 1 Закрытый = 0       
    Datetime       string                 // Дата создания
    DateStart      string                 // Дата открытия
    DateEnd        string                 // Дата закрытия
    DeadLine       string                 // Дата сдачи проекта по договору 
    Customer       string                 // Наименование заказчика
    Manager        string                 // Менеджер проекта
    Note           string                 // Краткое описание
    Num            string                 // Номер 
    Percent        int                    // Процент выполнения
    Status         string                 // Текущий статус
    Summ           float64                // Cумма проекта 
}

// Участники в проекте
type ProjectComands    struct {
     Id                string             // Уникальный код 
     ProjectId         int                // Код проекта
     RoleId            int                // Код роли
     UserId            int                // Код пользователя
     Name              string             // Фамилия + Имя 
     Position          string             // Должность
     CompanyId         string             // Id компании
     Company           string             // Название компании  
     Type              string             // Тип: 1-наш, 2-подрядчик, 3-заказчик
     Remark            string             // Примечание
     Tags              string             // Теги 
 }

// Справочник ролей
type RoleSpr    struct {
    Id          string                    // Уникальный код 
    Title       string                    // Роль
    Description string                    // Описание роли
}

// Справочник ролей
type Roles      struct {
    Id          string                    // Уникальный код 
    Title       string                    // Роль
    Description string                    // Описание роли
    Name        string                    // Имя проекта
    Price       float64                   // Стоимость человеко-часов
}      

// Команда в проекте
type Teams      struct {
     ProjectId  string                    // Id проекта
     RoleId     string                    // Роль в проекте
     Role       string                    // Наименование роли
     UserId     string                    // Id пользователя
     User       string                    // Имя + Фамилия пользователя 
     Created    string                    // Дата добавления участника в команду
     Remark     string                    // Краткое примечание об участнике
     InOut      string                    // Ext Int - внешний участник или внутренний
     Price      float64                   // Стоимость роли уастия в проекте
     Currency   string                    // Валюта 
 }

// Календарь
type  Calendars struct {
     Date       string                    // Дата 
     Title      string                    // Событие
 }

// Компания
type Company struct{
     Title    string                      // Имя компании
     Code     string                      // Код компании 
     Date     string                      // Дата вввода
     Address  string                      // Улица и дом
     Email    string                      // Почта
     Phone    string                      // Рабочий телефон 
     Mob      string                      // Мобильный телефон
     Postcode string                      // Почтовый код
     Remark   string                      // Примечание
     name     string                      // Наименование
}

// Базовый тип 
type Base struct {
     Title         string                 // Имя 
     Code          string                 // Код 
     Date          string                 // Дата вввода
     Created       string                 // Дата создания
     Updated       string                 // Дата обновления 
     CreatedAt     string                 // Кто создал
     UpdatedAt     string                 // Кто обновил
     Remark        string                 // Примечание
     Note          string                 // Примечение системный
     Description   string                 // Описание  
     Name          string                  `gorethink:"name"` 
}

// К
type Сnr struct{
     Base
     Sum   float64
}

// Проектный план
type ProjectPlan struct {
     Id          string                  `gorethink:"Id"` 
     DateStart   string                  // Дата старта
     DateFinish  string                  // Дата финиш
     Name        string                  // Наименование задачи (этапа)                               
     Description string                  // Описание
     Status      string                  // Plan 
     Teams       []string                // ID команда
     Summ        float64                 // Сумма

}


// Логирование операций
type LogOp struct{
    Id          string                  // Id 
    Title       string                  // Наименование операции 
    Date        string                  // Дата старта
    User        string                  // Пользователь
    Ip          string                  // Ip   
    Type        string                  // Тип (Warn - внимание | Info - информационная | Primary-обычная основная | Danger-опасная критическая операция)    
}



