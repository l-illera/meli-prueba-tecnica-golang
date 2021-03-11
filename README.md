# prueba-tecnica GOLANG

## Correr la aplicación en local

```shell script
./gradlew quarkusDev
```

> **_NOTA:_**  Quarkus ahora tiene incluida una interfaz de usuario Dev, que está disponible solamente en el modo dev en la URL: http://localhost:8080/q/dev/.

## Compilar y Ejecutar la Aplicación

### Localmente

La aplicación puede ser compilada con el siguiente comando:

```shell script
./gradlew build
```

Esto produce el archivo `quarkus-run.jar` en el directorio `build/quarkus-app/`. al compilarse de este modo, las
liberías quedan separadas del archivo `quarkus-run.jar` en el directorio `build/quarkus-app/lib/`.

Si quiere construir un jar completo (_über-jar_), Ejecute el comando:

```shell script
./gradlew build -Dquarkus.package.type=uber-jar
```

La aplicación ahora puede ser corrida desde la terminal utilizando el
comando `java -jar build/quarkus-app/quarkus-run.jar`.

### Google Cloud Run

Compilar la aplicación con el comando:

```shell script
./gradlew build
```

Inicie sesión en GCP utilizando el SDK
```shell script
gcloud auth login
```

Luego, use Cloud Build para compilar la imagen del proyecto, esto subirá a un depósito de Google Cloud Storage todos los
archivos de la aplicación (excepto las ignoradas por el archivo `.gcloudignore`).

```shell script
gcloud builds submit --tag gcr.io/PROJECT-ID/prueba-tecnica-java
```

Finalmente, se utiliza Cloud Run para iniciar la aplicación.

```shell script
gcloud run deploy --image gcr.io/PROJECT-ID/prueba-tecnica-java --platform managed
```

## Explorar el API

Este Projecto contiene una integración con OPENAPI y Swagger

### Achivo OpenApi

Para descargar el archivo openapi.yaml, haga clic en el siguiente enlace: [OpenApi.yml](https://prueba-tecnica-g6zov2ubiq-ue.a.run.app/q/openapi)

> **_NOTA:_** Para acceder a este archivo localmente, debe acceder a la ruta: http://localhost:8080/q/openapi

### Swagger UI

Para explorar la documentación de los endpoints Rest disponibles con la API de Swagger UI, haga clic en el siguiente enlace [Go To Swagger-UI](https://prueba-tecnica-g6zov2ubiq-ue.a.run.app/q/swagger-ui)
> **_NOTA:_** Para explorar el API localmente, debe acceder a la ruta: http://localhost:8080/q/swagger-ui