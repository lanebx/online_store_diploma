package main

import (
	_ "crypto/aes"
	"fmt"
	_ "net/http"
	_ "runtime/trace"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "html/template"
	_ "net/http"
	_ "path/filepath"

	_ "github.com/julienschmidt/httprouter"
)


var router *gin.Engine

func main()  {
	//запускает подключение к БД
	e := connect()
	if e!= nil{
		fmt.Println(e.Error())
		return
	}
	
	//переменная для работы веб сервера
	router = gin.Default()

	router.Static("/assets", cfg.Assets)

	router.Use(static.Serve("/", static.LocalFile(cfg.Data, false)))
	router.Use(static.Serve("/product/", static.LocalFile(cfg.Data, false)))
 
	word := sessions.NewCookieStore([]byte("my-private-key"))
	router.Use(sessions.Sessions("session", word))

	//агружает в память html файлы из которых будут шаблоны
	router.LoadHTMLFiles(cfg.HTML + "index.html", cfg.HTML + "contacts.html", cfg.HTML + "catalog.html", 
	cfg.HTML + "delivery.html", cfg.HTML + "return.html", cfg.HTML + "news.html",cfg.HTML + "pants.html",
	cfg.HTML + "outerwear.html",cfg.HTML + "jackets.html",cfg.HTML + "dresses.html",
	cfg.HTML + "knitwear.html", cfg.HTML + "product.html",)

	router.GET("/", index)

	router.GET("/catalog", catalog)
	router.GET("/news", news)
	router.GET("/contacts", contacts)
	router.GET("/delivery", delivery)
	router.GET("/return", returnProd)

	router.GET("/product/:id", products)

	//для категорий
	router.GET("/Брюки", pants)
	router.GET("/Верхній одяг", outerwear)
	router.GET("/Жакети", jackets)
	router.GET("/Сукні", dresses)
	router.GET("/Трикотаж", knitwear)


	// router.GET("/user/:name", func(c *gin.Context) {
    //     name := c.Param("name")
    //     c.String(http.StatusOK, "Hello %s", name)
    // })

	//Обработчик 
	router.POST("/login", login)
	router.POST("/addNews", addNews)
	router.POST("/addProduct", addProduct)
	router.POST("/addPur", addPur)

	//Запуск сервера
	router.Run(cfg.ServerHost +":"+ cfg.ServerPort)

	fmt.Println("работает все!")

}

func addPur(c *gin.Context){
	
	//в этой переменной теперь записаны ответы с формы
	var adpur InputAddPur

	e := c.BindJSON(&adpur)
	if e !=nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error":e.Error(),
		})
		return
	}
	
	e = adpur.Insert()
	if e != nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
}

func addProduct(c *gin.Context){
	
	//в этой переменной теперь записаны ответы с формы
	var adp InputAddProduct

	e := c.BindJSON(&adp)
	if e !=nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error":e.Error(),
		})
		return
	}
	e = adp.Insert()
	if e != nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
}

func addNews(c *gin.Context){
	
	//в этой переменной теперь записаны ответы с формы
	var adn InputAddNews

	e := c.BindJSON(&adn)
	if e !=nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error":e.Error(),
		})
		return
	}
	e = adn.Insert()
	if e != nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
}

//страници меню
func news(c *gin.Context){
	
	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	c.HTML(200, "news.html", gin.H{
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
func returnProd(c *gin.Context){
	
	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	c.HTML(200, "return.html", gin.H{
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
func delivery(c *gin.Context){
	
	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	c.HTML(200, "delivery.html", gin.H{
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
func catalog(c *gin.Context){
	
	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "catalog.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
func contacts(c *gin.Context){
	
	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	c.HTML(200, "contacts.html", gin.H{
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
//главная страница
func index(c *gin.Context){

	admin := false
	login := false

	s := sessions.Default(c)

	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.SelectIndexProd()
	if er != nil{
		fmt.Println(er.Error())
	}
	
	c.HTML(200, "index.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
//в этой функции нужно принять данные что были отправлены из js на сервер
func login(c *gin.Context){
	
	//в этой переменной уже записаны данные которые прилетели от пользователя
	var m manager  //переменная с типом структуры которая в файле db.go 
	
	e := c.BindJSON(&m)  
	if e != nil{
		fmt.Println(e.Error())

		//ТУТ пользователю вернятся название ошибки
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	e = m.Select()
	if e != nil{
		fmt.Println(e.Error())

		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	// сессия нужна для авторизации(добавляет в куки)
	s := sessions.Default(c)
	s.Set("MySecretKey", m.Role)

	e = s.Save()
	if e != nil{
		fmt.Println(e.Error())
	}

	c.JSON(200, gin.H{
		"Error": nil,
		"Name": m.Name,
	})
}


//функции каталога
// router.GET("/Брюки", pants)
	// router.GET("/Верхній%20одяг", outerwear)
	// router.GET("/Жакети", jackets)
	// router.GET("/Сукні", dresses)
	// router.GET("/Трикотаж", knitwear)

func pants(c *gin.Context){
	
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}

	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}

	var prod1 Product
	er := prod1.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "pants.html", gin.H{
		"Products" : prod1.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}

func outerwear(c *gin.Context){
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}
	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "outerwear.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}

func jackets(c *gin.Context){
	
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}
	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "jackets.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}

func dresses(c *gin.Context){
	
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}
	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "dresses.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}

func knitwear(c *gin.Context){
	
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}
	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "knitwear.html", gin.H{
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}

func products(c *gin.Context){

	var id = c.Param("id")
	
	admin := false
	login := false
	s := sessions.Default(c)
	//если этот интерфейс смог преобразоватся к строке то в ОК будет - тру
	role, ok := s.Get("MySecretKey").(string)
	if ok {
		if role == "Administrator"{
			admin = true
		}
		login = true
	}
	var cat category
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var prod Product
	er := prod.Select()
	if er != nil{
		fmt.Println(er.Error())
	}

	c.HTML(200, "product.html", gin.H{
		"Id" : id,
		"Products" : prod.Rows,
		"Category": cat.Rows,
		"Admin": 	admin,
		"IsLogin": login,
	})
}
