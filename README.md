# prueba-tecnica GOLANG

## Compilar y Ejecutar la Aplicación

### Localmente

La aplicación puede ser compilada con el siguiente comando:

```shell script
go build
```

Esto produce el archivo `prueba-tecnica.exe` en el directorio base del proyecto.

La aplicación tambien puede ser ejecutada desde la terminal utilizando el comando `go run main.go`.

### Google Cloud Run

Compilar la aplicación con el comando:

```shell script
go build
```

Inicie sesión en GCP utilizando el SDK

```shell script
gcloud auth login
```

Luego, use Cloud Build para compilar la imagen del proyecto, esto subirá a un depósito de Google Cloud Storage todos los
archivos de la aplicación (excepto las ignoradas por el archivo `.gcloudignore`).

```shell script
gcloud builds submit --tag gcr.io/PROJECT-ID/prueba-tecnica-golang
```

Finalmente, se utiliza Cloud Run para iniciar la aplicación.

```shell script
gcloud run deploy --image gcr.io/PROJECT-ID/prueba-tecnica-golang --platform managed
```

## Explorar el API

Este Projecto contiene una integración con OPENAPI y Swagger

### Achivo OpenApi

Para descargar el archivo openapi.yaml, haga clic en el siguiente
enlace: [OpenApi.yml](https://prueba-tecnica-golang-g6zov2ubiq-ue.a.run.app/swagger-ui/doc.json)

> **_NOTA:_** Para acceder a este archivo localmente, debe acceder a la ruta: http://localhost:8080/swagger-ui/doc.json

### Swagger UI

Para explorar la documentación de los endpoints Rest disponibles con la API de Swagger UI, haga clic en el siguiente
enlace [Go To Swagger-UI](https://prueba-tecnica-golang-g6zov2ubiq-ue.a.run.app/swagger-ui/)
> **_NOTA:_** Para explorar el API localmente, debe acceder a la ruta: http://localhost:8080/swagger-ui/