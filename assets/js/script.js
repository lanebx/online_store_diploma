/* 
<div class="modal">
    <i class="icon-cross"></i>
    <input type="text" id="UserLogin"placeholder="Логін">
    <input type="password" id="UserPassworld" placeholder="Пароль">
    <button>Увійти</button>
</div> 
    if (document.getElementsByClassName("modal") != false){
        return;
    }


    <div class="modal-add-news">
  <input type="text" placeholder="Заголовок новости">
  <textarea placeholder="Текст новости"></textarea>
  <input class="file" multiple>
  <div class="preview-class"></div>
</div>
    */
    

function OpenLoginWindow(){

    ckeckelem = document.getElementById("modal_1");
    if (ckeckelem){
        return false;
    }

    let modal  = document.createElement("div");
    modal.className = "modal";
    modal.id = "modal_1";
    
    let inputLodin = document.createElement("input");
    inputLodin.type = "text";
    inputLodin.placeholder = "Телефон без +38";

    let inputPass = document.createElement("input");
    inputPass.type = "password";
    inputPass.placeholder = "Пароль";

    let buttonL = document.createElement("button");
    buttonL.textContent = "Увійти";
    buttonL.onclick = Login.bind(buttonL, inputLodin, inputPass);

    let closeW = document.createElement("i");
    closeW.className = "icon-cross";
    closeW.onclick = function() {
        this.parentElement.remove();
    };

    modal.append(closeW, inputLodin, inputPass, buttonL);

    document.body.append(modal);
}

/** 
    Авторизация
* @param {HTMLInputElement} login
* @param {HTMLInputElement} pass
**/
function Login(login, pass){
    let isCorrect = true;
    
    if(!login || login.value.length !=10) {
        login.classList.add("incorrect");
        login.onclick = ClearIncorrect;
        isCorrect = false;
    }

    if(!pass || pass.value.length < 8) {
        pass.classList.add("incorrect");
        pass.onclick = ClearIncorrect;
        isCorrect = false;
    }

    if (!isCorrect){
        return;
    }

    //нужен для передачи данных на сервер
    let xhr = new XMLHttpRequest();

    //отправиь пост запрос по адресу /login 
    //(на сервере в папке main.go есть обработчик)
    xhr.open("POST", "/login");
    
    xhr.onload = function(event){
        try {
            let data = JSON.parse(event.target.response);
            if(data && "Error" in data && data.Error == null){
                if("Name" in data){
                    window.location.reload();
                }
            }else{
                console.log(data.Error)
            }
        } catch (error) {
            //ошибка в консоль браузера(нарим. неверный лог пароль)
            console.log(error);
        }

    };
    
    //JSON.stringify переводит обьект в строку заносит в переменную data
    //поле логин со значением login.value и пароль аналогично
    let data = JSON.stringify({
        Login:login.value,
        Passworld:pass.value,
    });

    //при авторизации в консоли браузера появятся логин и пароль под которыми вошли
    console.log(data);

    //для отправки значения логина и пароля, как строки
    xhr.send(data)
}
/**
 * Очистка неккоректного поля 
 */
function ClearIncorrect(){
    this.classList.remove("incorrect");
    this.value = "";
    this.onclick = undefined;
}


function createModalAddPurchase(article){
    ckeckelemNews = document.getElementById("modal-add-news");
    if (ckeckelemNews){
        return false;
    }
    let back  = document.createElement("div"),
      modal  = document.createElement("div"),
      close  = document.createElement("div"),
      icon_close  = document.createElement("i"),
      h2  = document.createElement("h2"),
      inputCount  = document.createElement("input"),
      inputTel = document.createElement("input"),
      inputSize = document.createElement("input"),
      textareaMess  = document.createElement("textarea"),
      buttonCreate  = document.createElement("div")
    ;
  
    back.append(modal)
    modal.append(close, h2, inputCount, inputTel, inputSize, textareaMess, buttonCreate)
    close.append(icon_close)
  
    back.className = "back";
    modal.className = "modal-add-news";
    close.className = "close";
    icon_close.className = "icon-cross";
    h2.textContent = "Заявка на замовлення";
    inputCount.type = "number";
    inputCount.step = "1";
    inputCount.placeholder = "Оберіть кількість";
    inputTel.type = "text";
    inputTel.placeholder = "Ваш телефон";
    inputSize.type = "text";
    inputSize.placeholder = "оберіть озмір: S, M, L";
    textareaMess.placeholder = "Ви можете залишити нам повідомлення...";
  
    buttonCreate.className = "button";
    buttonCreate.textContent = "Надіслати";
    
    icon_close.onclick = function() {
      this.parentElement.parentElement.parentElement.remove();
    }
    
    buttonCreate.onclick = createPur.bind(buttonCreate, article, inputCount, inputTel, inputSize, textareaMess);
    //console.log(article)
    document.body.append(back);
}

/**
 * новость
 * @param {HTMLInputElement} inputCount 
 * @param {HTMLInputElement} inputTel 
 * @param {HTMLInputElement} inputSize 
 * @param {HTMLTextAreaElement} textareaMess 
 * 
 */
 function createPur(article, purCount, purTel, purSize, purMess){
    let isCorrect = true;
  
    if(!purCount || purCount.value == ""){
        purCount.classList.add("incorrect");
        purCount.onclick = ClearIncorrect;
        isCorrect = false;
    }

    if(!purTel || purTel.value.length < 9){
      purTel.classList.add("incorrect");
      purTel.onclick = ClearIncorrect;
      isCorrect = false;
    }

    if(!purSize || purSize.value == ""){
        purSize.classList.add("incorrect");
        purSize.onclick = ClearIncorrect;
        isCorrect = false;
    }

    if(!purMess || purMess.value == ""){
        purMess.value == "-";
    }

    if (!isCorrect){
      return;
    }
    
    let xhr = new XMLHttpRequest();
  
    xhr.open("POST", "/addPur");
  
    //переводит обьект в строку
    let data2 = JSON.stringify({
      ArticlePur: article,
      CountPur: purCount.value,
      TelPur: purTel.value,
      SizePur: purSize.value,
      MessPur: purMess.value,
    });
      
    console.log(data2);
  
    xhr.send(data2);
  
    this.parentElement.parentElement.remove();
    alert("Дякую! Ваша заявка принятя, менеджер зателефонує протягом 1 робочого дня");
  
  }
  
  //Очистка incorrect
  function ClearIncorrect(){
    this.classList.remove("incorrect");
    this.value = "";
    this.onclick = undefined;
  }