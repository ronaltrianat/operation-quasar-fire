# Save TopSecret by Satelite

Este servicio recibe la informacion de los satelites de forma individual y se almacena en base de datos.

**URL** : `https://f0g0tobrb6.execute-api.us-east-1.amazonaws.com/operation-quasar-fire/api/topsecret_split/{satellite_name}`

**Method** : `POST`

**Solicitud de ejemplo**

```json
{
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
```

**Solicitud de ejemplo**
`POST` `https://f0g0tobrb6.execute-api.us-east-1.amazonaws.com/operation-quasar-fire/api/topsecret_split/skywalker`

```json
{
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
}
```

## Respuesta exitosa

**Condicion** : Los datos proporcionados son válidos

**Codigo** : `200 OK`

**Respuesta de ejemplo** :

```json
{
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
    "success": false,
    "messageResponse": "Ha ocurrido un error internamente. Por favor intente mas tarde",
    "technicalMessageResponse": "algun error"
}
```