# operation-quasar-fire
Challenge for Tech Lead in Mercado Libre.

## Descripcion de los Servicios

* [Top Secret](docs/topsecret.md) : `POST /topsecret`
* [Save Top Secret Split](docs/post_topsecret_split.md) : `POST /topsecret_split/{satellite_name}`
* [Get Top Secret Split](docs/get_topsecret_split.md) : `GET /topsecret_split?satellites=name`

## Postman con ejemplos de consumo

Con el siguiente archivo postman, puede realizar test de los servicios desplegado en AWS.

* [Export Postman Test Cloud](docs/meli-tl-challenge.postman_collection.json)
* [Export Postman Test Local](docs/meli-tl-challenge-local.postman_collection.json)

## Iniciar Aplicacion Localmente
* Prerequisitos:
    - docker
    - go v1.17.5
    - git

* git clone git@github.com:ronaltrianat/operation-quasar-fire.git
* cd operation-quasar-fire
* make docker-build
* make docker-run

## Diagrama de Clases

Se implementa una arquitectura hexagonal en la cual se busca encapsular la logica de negocio.

El microservicio esta diseñado para que se pueda crear una implementacion nueva de los puertos de entrada y de salida mediante nuevos adaptadores sin necesidad de afectar la logica de negocio ya construida.

Hay 2 casos de uso importantes aqui y es el tratamiento de los mensajes enviados por los satelites y la triangulacion de la posicion de la nave imperial.
Estos casos de uso fueron desarrollados de tal forma que se puedan implementar nuevas versiones ya sea por mejoramiento en el performance del algoritmo o para que diferentes mienbros del equipo se dividan la construccion de los mismos evitando colisionar o afectar el codigo del otro.

![Diagrama 3D](/docs/arch-diagram-meli.png)

## Arquitectura Basica de Despliegue AWS

Para disponibilizar de forma rapida la aplicacion se desplego directamente en instancias EC2 de tipo t2.micro(Architecture:	i386, x86_64, vCPUs: 1, Memory GiB: 1) con OS Amazon Linux de 64 bits.

Pero facilmente esta aplicacion puede ser desplegada en un cluster de ECS o EKS.

![Diagrama 3D](/docs/aws-3d.png)

## Load Test

Se realizan las pruebas de carga con la herramienta K6 [Load testing for engineering teams](https://k6.io/) configurando los siguientes stages:

```json
stages: [
    { duration: '10m', target: 100 }, // Los primeros 10 min levanta 100 usuarios virtuales 
    { duration: '10m', target: 200 }, // Durante los siguientes 10 min sube a 200 usuarios virtuales
    { duration: '10m', target: 0 }, // Durante los siguientes 10 min reduce a 0 los usuarios virtuales.
  ]
```
Cada usuario hace una peticion POST al servicio /topsecret cada segundo.

![Diagrama 3D](/docs/load-test.png)

Al final los resultados nos indican lo siguiente:

El servicio con esta prueba sin la necesidad de escalar fue capaz de procesar 88 RPS(Request per second) o 5280 RPM(Request per minute).

## Consideraciones de seguridad

Filtros de seguridad a nivel de backend:
* Autenticacion
* Autorizacion
* Cifrar/Descifrar request/response
* HMAC
* RECAPCHA

Hacer uso de patrones como API Gateway o Sidecar para evitar agregar logica adicional a los componentes core.

![Diagrama 3D](/docs/security-back-1.png)
![Diagrama 3D](/docs/security-back-2.png)


Seguridad a nivel de Infraestructura en la Nube:
* AWS WAF: Filtre el tráfico web malintencionado
* AWS Shield: Protección frente a ataques DDoS
* AWS KMS: Administración y almacenamiento clave
* AWS Secrets Manager: Alterne, administre y recupere datos confidenciales
* AWS Certificate Manager: Aprovisionamiento, administración e implementación de certificados públicos y privados SSL/TLS

![Diagrama 3D](/docs/security-infra-aws.png)

[AWS Security Products](https://aws.amazon.com/es/products/security/)

## Notificacion a los centros de comando en la tierra
![Diagrama 3D](/docs/topsecret_split_notification.png)

## Monitore y Almacenamiento de logs
Hacemos uso de un APM para validar metricas como:

* La cantidad de tiempo que toma cada solicitud.
* La cantidad de peticiones fallidas por segundo.
* El numero de peticiones procesadas por segundo.
* Visualizar la interaccion entre servicios.
* Generar alertas cuando algun servicio falle.
* Visualizar las solicitudes end-to-end que permitan identificar puntos de mejora.

![Monitoreo App](/docs/monitoreo-app.png)

## Golang Linter

[Golang Linter](https://golangci-lint.run/)

```
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.42.1 golangci-lint run -v
```