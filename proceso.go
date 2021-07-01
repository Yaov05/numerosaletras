package numerosaletras

//"Convertir" convierte un número a letras
func Convertir(numero int) string {
	var respuesta string
	if numero >= 10 {
		respuesta = primerosDiez[numero%10]
	} else {
		respuesta = unidades[numero]
	}
	return respuesta
}
