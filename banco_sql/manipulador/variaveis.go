package manipulador

import "html/template"

//ModeloOla armazena os modelos de dados
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena os modelos de dados
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
