# TopSecret

Este servicio recibe informacion de 3 satelites y a partir de la misma calcula el mensaje secreto y la posicion de la nave del imperio.

**URL** : `https://f0g0tobrb6.execute-api.us-east-1.amazonaws.com/operation-quasar-fire/api/topsecret`

**Method** : `POST`

**Solicitud de ejemplo**

```json
{
  "satellites": [
    {
      "name": "nombre", //nombre-del-satelite
      "distance": 500.0, //distancia-del-satelite-al-portacarga-imperial
      "message": [ // lista-string-con-el-mensaje-transmitido
        "",
        "es",
        "",
        "",
        "secreto"
      ],
      "position": { // posicion-x-y-del-satelite
        "x": 100,
        "y": -100
      }
    },
    {
      "name": "nombre", //nombre-del-satelite
      "distance": 500.0, //distancia-del-satelite-al-portacarga-imperial
      "message": [ // lista-string-con-el-mensaje-transmitido
        "",
        "es",
        "",
        "",
        "secreto"
      ],
      "position": { // posicion-x-y-del-satelite
        "x": 100,
        "y": -100
      }
    },
    {
      "name": "nombre", //nombre-del-satelite
      "distance": 500.0, //distancia-del-satelite-al-portacarga-imperial
      "message": [ // lista-string-con-el-mensaje-transmitido
        "",
        "es",
        "",
        "",
        "secreto"
      ],
      "position": { // posicion-x-y-del-satelite
        "x": 100,
        "y": -100
      }
    }
  ]
}
```

**Solicitud de ejemplo**

```json
{
  "satellites": [
    {
      "name": "skywalker",
      "distance": 500.0,
      "message": [
        "",
        "es",
        "",
        "",
        "secreto"
      ],
      "position": {
        "x": 100,
        "y": -100
      }
    },
    {
      "name": "sato",
      "distance": 728.0110,
      "message": [
        "este",
        "",
        "un",
        "",
        ""
      ],
      "position": {
        "x": 500,
        "y": 100
      }
    },
    {
      "name": "kenobi",
      "distance": 583.0952,
      "message": [
        "este",
        "",
        "",
        "mensaje",
        ""
      ],
      "position": {
        "x": -500,
        "y": -200
      }
    }
  ]
}
```

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