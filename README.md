# operation-quasar-fire
Challenge for Tech Lead in Mercado Libre.

## Descripcion de los Servicios

* [Top Secret](docs/topsecret.md) : `POST /topsecret`
* [Save Top Secret Split](docs/post_topsecret_split.md) : `POST /topsecret_split/{satellite_name}`
* [Get Top Secret Split](docs/get_topsecret_split.md) : `GET /topsecret_split?satellite_name=name`

## Postman con ejemplos de consumo

Con el siguiente archivo postman, puede realizar test de los servicios desplegado en AWS.

* [Export Postman Test Cloud](docs/meli-tl-challenge.postman_collection.json)
* [Export Postman Test Local](docs/meli-tl-challenge-local.postman_collection.json)

## Iniciar Aplicacion Localmente
* Prerequisitos:
    - docker
    - docker-compose
    - go v1.15

* git clone git@github.com:ronaltrianat/meli-tl-challenge.git
* ejecutar: docker-compose -f docker-compose-dynamodb-local.yaml up -d
* ejecutar: go run main.go

## Arquitectura Despliegue AWS

![Diagrama 3D](/docs/aws-3d.png)
![Diagrama 2D](/docs/aws-2d.png)

### linter con docker
```
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.31.0 golangci-lint run -v
```