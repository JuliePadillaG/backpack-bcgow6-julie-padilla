package web

import "strconv"

// Realizar la estructura Response con los campos: code, data y error:
// 	code tendra el codigo de retorno.
// 	data tendrá la estructura que envía la aplicación (en caso que no haya error).
// 	error tendrá el error recibido en formato texto (en caso que haya error).

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
 } 

// Desarrollar una función que reciba el code cómo entero, data como interfaz y error como string.
// La función debe retornar en base al código, si es una respuesta con el data o con el error.
// Implementar esta función en todos los retornos de los controladores, antes de enviar la respuesta al cliente la función debe generar la estructura que definimos.

func NewResponse(code int, data interface{}, err string) Response {

	if code < 300{
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
  
}



