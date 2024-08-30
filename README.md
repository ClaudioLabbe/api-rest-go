# API REST con Go y Gin

Este proyecto es una API REST construida con el framework [Gin](https://github.com/gin-gonic/gin) en Go. La API permite realizar operaciones CRUD (Crear, Leer, Actualizar y Eliminar) para gestionar álbumes.

## Características

- CRUD para Álbumes
- Arquitectura organizada en controladores y rutas
- Configuración con conexión a PostgreSQL

## Requisitos Previos

- [Go](https://golang.org/dl/) 1.20 o superior

## Instalación

1. **Clona este repositorio:**

    ```bash
    git clone https://github.com/tu-usuario/api-rest-go.git
    cd api-rest-go
2. **Instala las dependencias:**
    ```bash
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get github.com/joho/godotenv
3. **Crea un archivo .env en la raíz del proyecto para las variables de entorno**
    ```bash
    DB_HOST=localhost
    DB_PORT=puerto
    DB_USER=tu_usuario
    DB_PASSWORD=tu_contraseña
    DB_NAME=tu_base_de_datos
## Ejecución:
    go run main.go

Por defecto, el servidor se ejecutará en http://localhost:8080.

## Endpoints de la API

**API de Álbumes**

Esta API permite gestionar álbumes a través de varios endpoints.

| Método | Endpoint       | Descripción                              |
|--------|----------------|------------------------------------------|
| `GET`    | /album         | Obtiene todos los álbumes                |
| `GET`    | /album/:id     | Obtiene un álbum por su ID               |
| `POST`   | /album         | Crea un nuevo álbum                      |
| `PATCH`  | /album/:id     | Actualiza un álbum por su ID             |
| `DELETE` | /album/:id     | Elimina un álbum por su ID               |
