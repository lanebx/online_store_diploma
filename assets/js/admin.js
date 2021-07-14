//console.log("Привет админ")
//console.log( document.cookie );
console.log( imgnews );


//Удаление авторизации
function CookiesDelete() {
	var cookies = document.cookie.split(";");
	for (var i = 0; i < cookies.length; i++) {
		var cookie = cookies[i];
		var eqPos = cookie.indexOf("=");
		var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
		document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT;";
		document.cookie = name + '=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
	}
    location.reload()
}

/**
 *  Добавление продукта
 */
var imgprod;
function changeImagesProd(){
  let preview = document.querySelector(".preview-image");
  if(!preview){
      console.warn("Елемент [.preview-image] не найден");
      return;
  }
  for (const file of this.files){
      let img = document.createElement("img");
      img.src = window.URL.createObjectURL(file);
      img.height = 100;

      preview.append(img);
      console.log(file.name)
      imgprod = file.name;
  }
}

function createModalAddProduct(){
  ckeckelemNews = document.getElementById("modal-add-news");
  if (ckeckelemNews){
      return false;
  }
  let back  = document.createElement("div"),
    modal  = document.createElement("div"),
    close  = document.createElement("div"),
    icon_close  = document.createElement("i"),
    h2  = document.createElement("h2"),
    inputName  = document.createElement("input"),
    inputPrise  = document.createElement("input"),
    inputCat = document.createElement("input"),
    textareaText  = document.createElement("textarea"),
    label  = document.createElement("label"),
    inputFile  = document.createElement("input"),
    buttonFile  = document.createElement("div"),
    previewImage  = document.createElement("div"),
    buttonCreate  = document.createElement("div")
  ;

  back.append(modal)
  modal.append(close, h2, inputName, inputPrise, inputCat, textareaText, label, previewImage, buttonCreate)
  close.append(icon_close)
  label.append(inputFile, buttonFile )

  back.className = "back";
  modal.className = "modal-add-news";
  close.className = "close";
  icon_close.className = "icon-cross";
  h2.textContent = "Додати продукт";
  inputName.type = "text";
  inputName.placeholder = "Назва продукту";
  inputPrise.type = "number";
  inputPrise.step="500";
  inputPrise.placeholder = "Ціна продукту";
  inputCat.type = "text";
  inputCat.placeholder = "Категорія продукту";
  textareaText.placeholder = "Опис продукту";
  label.className = "file";

  inputFile.type = "file";
  inputFile.accept = "image/*";
  inputFile.multiple = true;
  inputFile.onchange = changeImagesProd;

  buttonFile.className = "button";
  buttonFile.textContent = "Обрати зображення";
  previewImage.className = "preview-image";
  buttonCreate.className = "button";
  buttonCreate.textContent = "Додати";
  
  icon_close.onclick = function() {
    this.parentElement.parentElement.parentElement.remove();
  }

  buttonCreate.onclick = createProd.bind(buttonCreate, inputName, inputPrise, inputCat, textareaText, inputFile);

  document.body.append(back);
}

/**
 * продукт
 * @param {HTMLInputElement} nameProd
 * @param {HTMLInputElement} priseProd
 * @param {HTMLInputElement} catProd
 * @param {HTMLTextAreaElement} textProd 
 * @param {HTMLInputElement} imgProd 
 */
 function createProd(nameProd, priseProd, catProd, textProd, imgProd){
  let isCorrect = true;

  if(!nameProd || nameProd.value.length < 3){
    nameProd.classList.add("incorrect");
    nameProd.onclick = ClearIncorrect;
    isCorrect = false;
  }

  if(priseProd.value == ""){
    priseProd.classList.add("incorrect");
    priseProd.onclick = ClearIncorrect;
    isCorrect = false;
  }

  if(catProd.value == "Верхній одяг" || catProd.value == "Брюки" || catProd.value == "Жакети" || catProd.value == "Сукні" || catProd.value == "Трикотаж"){
    let a = 2;
  }else {
    catProd.classList.add("incorrect");
    catProd.onclick = ClearIncorrect;
    isCorrect = false;
  }

  if(!textProd || textProd.value.length < 5){
    textProd.classList.add("incorrect");
    textProd.onclick = ClearIncorrect;
    isCorrect = false;
  }

  imgProd = imgprod;
  console.log( "это имя файла:", imgProd );

  if (!isCorrect){
    return;
  }

  let xhr = new XMLHttpRequest();

  xhr.open("POST", "/addProduct");

  //переводит обьект в строку
  let data2 = JSON.stringify({
    NameProduct: nameProd.value,
    PriseProduct: priseProd.value,
    CatProduct: catProd.value,
    TextProduct: textProd.value,
    ImgProduct: imgProd,
  });
    
  console.log(data2);

  xhr.send(data2);

  this.parentElement.parentElement.remove();
  alert("Товар успішно доданий");

}


/**
 *  Добавление новости
 * 
 * 
 */
 var imgnews;

 function changeImages(){
    let preview = document.querySelector(".preview-image");
    if(!preview){
        console.warn("Елемент [.preview-image] не найден");
        return;
    }
    for (const file of this.files){
        let img = document.createElement("img");
        img.src = window.URL.createObjectURL(file);
        img.height = 100;

        preview.append(img);
        console.log(file.name)
        imgnews = file.name;
    }
}

function createModalAddNews(){
  ckeckelemNews = document.getElementById("modal-add-news");
  if (ckeckelemNews){
      return false;
  }
  let back  = document.createElement("div"),
    modal  = document.createElement("div"),
    close  = document.createElement("div"),
    icon_close  = document.createElement("i"),
    h2  = document.createElement("h2"),
    inputName  = document.createElement("input"),
    textareaText  = document.createElement("textarea"),
    label  = document.createElement("label"),
    inputFile  = document.createElement("input"),
    buttonFile  = document.createElement("div"),
    previewImage  = document.createElement("div"),
    buttonCreate  = document.createElement("div")
  ;

  back.append(modal)
  modal.append(close, h2, inputName, textareaText, label, previewImage, buttonCreate)
  close.append(icon_close)
  label.append(inputFile, buttonFile )

  back.className = "back";
  modal.className = "modal-add-news";
  close.className = "close";
  icon_close.className = "icon-cross";
  h2.textContent = "Створити новину";
  inputName.type = "text";
  inputName.placeholder = "Заголовок новини";
  textareaText.placeholder = "Текст новини";
  label.className = "file";

  inputFile.type = "file";
  inputFile.accept = "image/*";
  inputFile.multiple = true;
  inputFile.onchange = changeImages;

  buttonFile.className = "button";
  buttonFile.textContent = "Обрати зображення";
  previewImage.className = "preview-image";
  buttonCreate.className = "button";
  buttonCreate.textContent = "Створити";
  
  icon_close.onclick = function() {
    this.parentElement.parentElement.parentElement.remove();
  }

  buttonCreate.onclick = createNews.bind(buttonCreate, inputName, textareaText, inputFile);

  document.body.append(back);

}

/**
 * новость
 * @param {HTMLInputElement} nameNews 
 * @param {HTMLTextAreaElement} textNews 
 * @param {HTMLInputElement} imgNews 
 * 
 */
function createNews(nameNews, textNews, imgNews){
  let isCorrect = true;

  if(!nameNews || nameNews.value.length < 5){
    nameNews.classList.add("incorrect");
    nameNews.onclick = ClearIncorrect;
    isCorrect = false;
  }

  if(!textNews || textNews.value.length < 5){
    textNews.classList.add("incorrect");
    textNews.onclick = ClearIncorrect;
    isCorrect = false;
  }

  imgNews = imgnews;
  console.log( "это имя файла:", imgNews );

  if (!isCorrect){
    return;
  }

  let xhr = new XMLHttpRequest();

  xhr.open("POST", "/addNews");

  //переводит обьект в строку
  let data1 = JSON.stringify({
    NameNews: nameNews.value,
    TextNews: textNews.value,
    ImgNews: imgNews,
  });
    
  console.log(data1);

  xhr.send(data1);

  this.parentElement.parentElement.remove();
  alert("Новина успішно додана!");

}

//Очистка incorrect
function ClearIncorrect(){
  this.classList.remove("incorrect");
  this.value = "";
  this.onclick = undefined;
}