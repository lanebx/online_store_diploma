package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB

//обьявление словарь которій хранит в себе запрос к БД.
var Queries map[string]*sql.Stmt

func connect() error {
	var e error

	db, e = sql.Open("mysql", cfg.MySQLUser + ":" + cfg.MySQLPass+"@tcp("+ cfg.MySQLHost + ":" + cfg.MySQLPort+ ")/" + cfg.MySQLBase)
	if e != nil{
		fmt.Println("Проблема db.go connect()")
	} 
	//виделить память под словарь
	Queries = make(map[string]*sql.Stmt)
	prerareQueries()

	return nil
}

func prerareQueries(){
	var e error
	Queries["Select#Category"], e = db.Prepare(`SELECT Name FROM category ORDER BY Name`)
	if e != nil{
		fmt.Println(e.Error())
	}

	///
	Queries["Select#Product"], e = db.Prepare(`SELECT * FROM product`)
	if e != nil{
		fmt.Println(e.Error())
	}

	Queries["SelectIndex#Product"], e = db.Prepare(`SELECT * FROM product ORDER BY Articute DESC LIMIT 12`)
	if e != nil{
		fmt.Println(e.Error())
	}


	Queries["Select#User"], e = db.Prepare(`SELECT Name FROM user WHERE Phone=? AND Passworld=?`)
	if e != nil{
		fmt.Println(e.Error())
	}

	Queries["Select#Manager"], e = db.Prepare(`SELECT Name, Role FROM manager WHERE Phone=? AND Passworld=?`)
	if e != nil{
		fmt.Println(e.Error())
	}

	Queries["Insert#News"], e = db.Prepare(`INSERT INTO news(Caption, Content, Image, Published, Manager ) VALUES(?, ?, ?, NOW(), "0990674737")`)
	if e != nil{
		fmt.Println(e.Error())
	}

	Queries["Insert#Product"], e = db.Prepare(`INSERT INTO product(Name, Prise, Availability, Image, Descriptions, Category ) VALUES(?, ?, "1", ?, ?, ?)`)
	if e != nil{
		fmt.Println(e.Error())
	}

	Queries["Insert#Purchase"], e = db.Prepare(`INSERT INTO purchaseone(Product, Count, telclient, Size, Message, Date ) VALUES(?, ?, ?, ?, ?, NOW())`)
	if e != nil{
		fmt.Println(e.Error())
	}

}


type InputAddPur struct{
	//с помощью этой записи, поле соотвествует полю в json
	ArticlePur string `json:"ArticlePur"` //Product
	CountPur string `json:"CountPur"`	  //Count
	TelPur string `json:"TelPur"`		  //telclient
	SizePur string `json:"SizePur"`		  //Size
	MessPur string `json:"MessPur"`		  //Message

	Rows []InputAddPur
}

func (m *InputAddPur) Insert() error{
	stmt, ok := Queries["Insert#Purchase"]
	if !ok {
		return errors.New("Не найден запрос Insert#Purchase")
	}

	_, e :=stmt.Exec(m.ArticlePur, m.CountPur, m.TelPur, m.SizePur, m.MessPur)
	if e !=nil {
		return e
	}

	return nil
}

type Product struct{
	Articute string
	Name string
	Prise string
	Availability string
	Image string
	Descriptions string
	Category string

	Rows []Product
}

func (m *Product) Select() error{
	stmt, ok := Queries["Select#Product"]
	if !ok {
		return errors.New("Не найден запрос Select#Product")
	}
	rows, e := stmt.Query()
	if e != nil{
		return e
	}
	for rows.Next(){
		e = rows.Scan(&m.Articute, &m.Name,  &m.Prise, &m.Availability, &m.Image, &m.Descriptions, &m.Category)
		if e != nil{
			fmt.Println(e.Error())
		}

		m.Rows = append(m.Rows, Product{Articute: m.Articute, Name: m.Name, Prise: m.Prise, Availability: m.Availability, Image: m.Image, Descriptions: m.Descriptions, Category: m.Category})
	}
	return nil
}

func (m *Product) SelectIndexProd() error{
	stmt, ok := Queries["SelectIndex#Product"]
	if !ok {
		return errors.New("Не найден запрос Select#Product")
	}
	rows, e := stmt.Query()
	if e != nil{
		return e
	}
	for rows.Next(){
		e = rows.Scan(&m.Articute, &m.Name,  &m.Prise, &m.Availability, &m.Image, &m.Descriptions, &m.Category)
		if e != nil{
			fmt.Println(e.Error())
		}

		m.Rows = append(m.Rows, Product{Articute: m.Articute, Name: m.Name, Prise: m.Prise, Availability: m.Availability, Image: m.Image, Descriptions: m.Descriptions, Category: m.Category})
	}
	return nil
}


//структура добавления продукта в БД 
type InputAddProduct struct{
	//с помощью этой записи, поле соотвествует полю в json
	NameProduct string `json:"NameProduct"`   	//Name
	PriseProduct string `json:"PriseProduct"`	//Prise
	CatProduct string `json:"CatProduct"`		//Category
	TextProduct string `json:"TextProduct"`		//Descriptions
	ImgProduct string `json:"ImgProduct"`			//Image

	Rows []InputAddProduct
}

func (m *InputAddProduct) Insert() error{
	stmt, ok := Queries["Insert#Product"]
	if !ok {
		return errors.New("Не найден запрос Insert#Product")
	}

	_, e :=stmt.Exec(m.NameProduct, m.PriseProduct, m.ImgProduct, m.TextProduct, m.CatProduct)
	if e !=nil {
		return e
	}

	return nil
}

//структура новости
type InputAddNews struct{
	//с помощью этой записи, поле соотвествует полю в json
	//NumbNews string `json:"NumbNews"`   //ID
	NameNews string `json:"NameNews"`   //Caption
	TextNews string `json:"TextNews"`	//Content
	ImgNews string `json:"ImgNews"`		//Image
	Published string					//Published
	Manager string						//Manager

	Rows []InputAddNews
}

//функция добавления новости в БД
func (m *InputAddNews) Insert() error{
	stmt, ok := Queries["Insert#News"]
	if !ok {
		return errors.New("Не найден запрос Insert#News")
	}

	_, e :=stmt.Exec(m.NameNews, m.TextNews, m.ImgNews)
	if e !=nil {
		return e
	}

	return nil
}

//структура категории
type category struct{
	Name string

	Rows []category
}

func (m *category) Select() error{
	stmt, ok := Queries["Select#Category"]
	if !ok {
		return errors.New("Не найден запрос Select#Category")
	}

	rows, e := stmt.Query()
	if e != nil{
		return e
	}

	for rows.Next(){
		e = rows.Scan(&m.Name)
		if e != nil{
			fmt.Println(e.Error())
		}

		m.Rows = append(m.Rows, category{Name: m.Name})
	}
	return nil
}

//структура для логина и пароля(в структуре можно давать названия произвольно, а вот 
//Login и Passworld аналогичны названиям из js, функция будет искать такие поля в json
//логин и пароль будут записаны в Login string и Pass string
type manager struct{
	Login 		string `json:"Login"`
	Passworld 	string `json:"Passworld"`
	Name 		string `json:"Name"`
	Role 		string `json:"Role"`

	Rows []manager
}

//функция пренадлежит структуре manager и возращает error
func (m *manager) Select() error{

	stmtUser, ok := Queries["Select#User"]
	if !ok {
		return errors.New("Не найден запрос Queries[Select#User]")
	}

	stmtManager, ok := Queries["Select#Manager"]
	if !ok {
		return errors.New("Не найден запрос Queries[Select#Manager]")
	}

	r := stmtUser.QueryRow(m.Login, m.Passworld)
	e := r.Scan(&m.Name)
	if e != nil{
		fmt.Println(e.Error(), "не нашло в юзерах")

		r = stmtManager.QueryRow(m.Login, m.Passworld)
		e := r.Scan(&m.Name, &m.Role)
		if e != nil{
			fmt.Println(e.Error(), "не нашло в менеджерах")

			return errors.New("Неверный логин/пароль")
		}
	}
	return nil
}