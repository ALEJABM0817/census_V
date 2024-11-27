

##  CENSUS VISUALIZER


## Introducción
Census Visualizer es una aplicación web diseñada para gestionar y visualizar datos de censos de manera eficiente y accesible. La aplicación permite a los usuarios cargar datos desde archivos CSV, registrarse e iniciar sesión para una experiencia personalizada. Una vez cargados, los datos pueden ser explorados y filtrados dinámicamente según diversas categorías y criterios.

## Requisitos 
Crear un archivo .env en /censusV y censusV/frontend

En /censusV agregar las siguientes variables

``` bash 
    POSTGRES_USER=me
    POSTGRES_PASSWORD=password
    POSTGRES_DB=census_db
    DB_HOST=localhost
    DB_PORT=5435
    DB_SSL=disable
    JWT_SECRET_KEY=supersecret
```

y en /censusV/frontend

VITE_URL_HOST=http://localhost:8080

## Comandos

En una ventana de la terminal ir a cd censusV y ejecutar
``` bash 
    docker-compose up --build -d
```

ir al archivo censusV\main.go y cambiar la variable var loadDataOnRun bool = false a true (var loadDataOnRun bool = true)

luego ejecutar

``` bash 
    go run .
```

luego en otro ventana ir a cd censusV/frontend
``` bash 
    yarn install
    yarn dev
```


## Despliegue

https://census-v.vercel.app/

