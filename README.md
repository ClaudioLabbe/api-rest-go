# API REST con Go y Gin

Este proyecto es una API REST construida con el framework [Gin](https://github.com/gin-gonic/gin) en Go. La API permite realizar operaciones CRUD (Crear, Leer, Actualizar y Eliminar) para gestionar álbumes.

## Características

- CRUD para Álbumes
- Arquitectura organizada en controladores y rutas
- Uso de variables globales como almacenamiento temporal de datos (sin base de datos)

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

## Ejecución:
    go run main.go

Por defecto, el servidor se ejecutará en http://localhost:8080.

## Endpoints de la API

**API de Álbumes**

Esta API permite gestionar álbumes a través de varios endpoints.

| Método | Endpoint       | Descripción                              |
|--------|----------------|------------------------------------------|
| `GET`    | /album         | Obtiene todos los álbumes                |
| `POST`   | /album         | Crea un nuevo álbum                      |
| `GET`    | /album/:id     | Obtiene un álbum por su ID               |
| `PATCH`  | /album/:id     | Actualiza un álbum por su ID             |
| `DELETE` | /album/:id     | Elimina un álbum por su ID               |

## Estructura del Proyecto
El proyecto sigue una estructura de carpetas modular:

api-rest-go/
├── src/
│   ├── controllers/    # Controladores para manejar la lógica de negocio
│   │   ├── album_controller.go
│   ├── routes/         # Definición de rutas
│   │   └── routes.go
│   └── models/         # Definición los modelos
│       └── album.go
├── .gitignore          # Archivos y directorios ignorados por Git
├── go.mod              # Módulo de Go y dependencias
├── go.sum              # Suma de comprobación de dependencias de Go
└── main.go             # Punto de entrada principal de la aplicación


