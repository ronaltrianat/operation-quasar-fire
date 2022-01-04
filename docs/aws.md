### Desplegar imagen local en AWS ECR 
1. Crear repositorio(ECR) imagen docker que contiene la aplicacion
```
aws ecr create-repository --repository-name meli-tl-challenge
```
2. Taggear imagen local con ruta de la imagen en ECR
```
docker tag meli-tl-challenge:latest accountid.dkr.ecr.us-east-2.amazonaws.com/meli-tl-challenge
```
3. Obtener credenciales para subir imagen en ECR
```
aws ecr get-login --no-include-email
```
4. Subir imagen local a ECR
```
docker push 797788214397.dkr.ecr.us-east-2.amazonaws.com/meli-tl-challenge
```