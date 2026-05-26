# Web Interacción Persona-Ordenador

1. Instalar Git
   https://git-scm.com/

2. Instalar Go

   Windows:

   * Descargar el instalador desde:
     https://go.dev/dl/
   * Ejecutar el `.msi`
   * Reiniciar terminal después de instalar

   Mac:

   * Descargar el `.pkg` desde:
     https://go.dev/dl/
   * Instalar normalmente

3. Verificar instalación de Go
   Abrir terminal y ejecutar:

   ```bash
   go version
   ```

4. Clonar el repositorio

   ```bash
   git clone https://github.com/oliviatroitino/web-ipo.git
   ```

5. Entrar en la carpeta del proyecto

   ```bash
   cd repositorio
   ```

6. Descargar dependencias

   ```bash
   go mod tidy
   ```

7. Ejecutar el servidor

   ```bash
   go run ./cmd/server/
   ```

8. Abrir la web en el navegador
   Normalmente:

   ```text
   http://localhost:8080
   ```

   o el puerto que aparezca en la terminal.
