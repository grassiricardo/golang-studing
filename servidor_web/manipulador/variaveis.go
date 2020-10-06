package manipulador

import "html/template"

//Modelos armazena os modelos de dados
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
