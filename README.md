# Realtime Chat with Go and ReactJS

## Iniciando el Backend con GO

1. Crear una carpeta llamada "backend" y navegar a ella usando la línea de comandos:

   ```bash
   cd backend
   ```

2. Habilitar el uso de módulos Go con el siguiente comando:

   ```bash
   set GO111MODULE=on
   ```

3. Inicializar un nuevo módulo Go llamado "Realtime_Chat":

   ```bash
   go mod init Realtime_Chat
   ```

4. Dentro de la carpeta "backend," crear un archivo llamado "main.go" con el siguiente contenido:

   ```go
   package main

   import "fmt"

   func main() {
      fmt.Println("Chat App v0.01")
   }
   ```

5. Finalmente, ejecutar el programa en Go con el siguiente comando:

   ```bash
   go mod init Realtime_Chat
   ```

6. Esto imprimirá en la consola:
   ```bash
   Chat App v0.01
   ```

## Inicializando un Proyecto de ReactJS con Vite

Para comenzar, sigue estos pasos para crear un proyecto de ReactJS utilizando Vite en la carpeta `frontend`.

1. Crea la carpeta `frontend` en tu directorio de proyectos.

2. Abre la terminal en la carpeta `frontend` y ejecuta los siguientes comandos:

   ```bash
   npm init @vitejs/app .
   ```

3. Durante la configuración, selecciona "React" como template y, si lo prefieres, elige "TypeScript + SWC" como lenguaje de programación. Luego, ejecuta:

   ```bash
   npm install
   ```

4. Para comprobar que todo se ha instalado correctamente, utiliza el siguiente comando:

   ```bash
   npm run dev -- --host 0.0.0.0
   ```

Esto abrirá tu proyecto en el puerto 5173. Accede a él en tu navegador:

[http://localhost:5173/](http://localhost:5173/)

Cuando desees finalizar la ejecución, simplemente pulsa `Ctrl + C` en la terminal para detener el servidor.

## Iniciando Server

1. Para comenzar, nos ubicamos en nuestro backend, primero asegúrate de tener instalada la biblioteca `gorilla/websocket` usando el comando:

   ```bash
   go get github.com/gorilla/websocket
   ```

2. Luego, puedes probar un servidor web simple utilizando el siguiente código en Go:

   ```go
   package main

   import (
      "fmt"
      "net/http"
   )

   func setupRoutes() {
      http.HandleFunc("/", func(w http.ResponseWriter, r \*http.Request) {
         fmt.Fprintf(w, "Hola Mundo, es Server!!")
      })
   }

   func main() {
      setupRoutes()
      http.ListenAndServe(":3000", nil)
   }
   ```

   Al ejecutar este código, tu servidor estará en funcionamiento en el puerto 3000 [http://localhost:3000/](http://localhost:3000/) y mostrará el mensaje **"Hola Mundo, es Server!!"** en el navegador cuando accedas a él.
