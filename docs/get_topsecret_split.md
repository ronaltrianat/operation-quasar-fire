# Search TopSecret by satellite names 

Este servicio recibe los nombres de los satelites sobre los cuales se deben realizar los calculos de la posicion y mensaje transmitidos por el portacarga imperial. La informacion de cada satelite ya fue almacenada con anterioridad por otro servicio. Los nombres de los satelites se envian como query parameter. El nombre del parametro es: satellite_name

**URL** : `https://ukqi44w6cl.execute-api.us-east-2.amazonaws.com/dev/topsecret_split?satellite_name=nombre&satellite_name=nombre&satellite_name=nombre`

**Method** : `GET`

**Solicitud de ejemplo**

`GET` `https://ukqi44w6cl.execute-api.us-east-2.amazonaws.com/dev/topsecret_split?satellite_name=kenobi&satellite_name=skywalker&satellite_name=sato`

## Respuesta exitosa

**Condicion** : Los datos proporcionados son válidos

**Codigo** : `200 OK`

**Respuesta de ejemplo** :

```json
{
    "position": {
        "x": -200,
        "y": 300
    },
    "message": "este es un mensaje secreto",
    "success": true,
    "messageResponse": "Operación ejecutada exitosamente",
    "technicalMessageResponse": ""
}
```

## Respuesta Fallida

**Condition** : Si los datos proporcionados son invalidos

**Codigo** : `200 OK`

**Respuesta de ejemplo** :

```json
{
    "position": {
        "x": 0,
        "y": 0
    },
    "message": "",
    "success": false,
    "messageResponse": "Ha ocurrido un error internamente. Por favor intente mas tarde",
    "technicalMessageResponse": "Ha sido imposible obtener la ubicación del satélite imperial"
}
```