# Guía de uso de Git para el proyecto
Guía paso a paso para que se conecten al repositorio y puedan empezar a trabajar en su propia rama.

> **Antes que nada:**
> 1. Crear una cuenta en [GitHub](https://github.com).
> 2. Avisarme de sus usuarios así los puedo agregar como colaboradores. **Si no no van a poder hacer nada**.
> 3. Instalar git en el ordenador. Si abren una terminal y ponen `git --version` y les sale un número, es que ya lo tienen instalado. Si no:
>    - **Mac:** instalar con Homebrew (`brew install git`).
>    - **Windows:** instalar [Git for Windows](https://git-scm.com/download/win) y usar la aplicación "Git Bash" que viene incluida.

## 1. Configurar git (solo la primera vez)
Git necesita saber quiénes son para firmar los commits. Tiene que ser el correo que usaron para GitHub.

```zsh
git config --global user.name "Tu Nombre"
git config --global user.email "tu-correo@ejemplo.com"
```

Esto solo hay que hacerlo una vez por ordenador, no por proyecto.

## 2. Clonar el repositorio
Esto hace que tengan una versión del repositorio en su ordenador. Primero tienen que meterse con la terminal en la carpeta donde quieran guardar el proyecto (por ejemplo, Documents):

```zsh
cd ~/Documents
git clone https://github.com/oliviatroitino/web-ipo.git
cd web-ipo
```

A partir de acá, todos los comandos de git se ejecutan dentro de esta carpeta.

## 3. Crear su propia rama

La rama `main` es el código "oficial" del proyecto. Nadie trabaja directamente sobre ella: cada uno trabaja en su propia rama y luego la fusionamos cuando esté lista. Para crear una rama nueva y cambiarse a ella a la vez:
```zsh
git switch -c nombre-apellido
```

Para comprobar en qué rama están en cualquier momento:
```zsh
git branch
```

La rama con el asterisco `*` es en la que están ahora.

## 4. Subir la rama a GitHub
Hasta ahora la rama solo existe en su ordenador. Para que aparezca en GitHub y los demás la puedan ver, hay que "empujarla" la primera vez con `-u`:

```zsh
git push -u origin nombre-apellido
```

El `-u` solo hace falta la primera vez. Las siguientes veces que quieran subir cambios basta con `git push origin nombre-apellido`.

## 5. Flujo de trabajo diario

Cada vez que vayan a trabajar:

```zsh
# 1. Asegurarse de estar en su rama
git switch nombre-apellido

# 2. Traer los últimos cambios del repositorio remoto
git pull

# 3. Trabajar en los ficheros que toquen...

# 4. Ver qué hayan cambiado
git status

# 5. Marcar los cambios para el próximo commit
git add .

# 6. Guardar el commit con un mensaje que explique el porqué
git commit -m "feat: añadir sección de técnicas de evaluación"

# 7. Subir los cambios a GitHub
git push
```

### Sobre los mensajes de commit
- Escribirlos en imperativo ("añadir", "corregir"), no en pasado.
- Formato `<tipo>: <descripción corta en imperativo>`. Los tipos más frecuentes:
    - feat: nueva funcionalidad.
    - fix: corrección de un error.
    - docs: cambios solo en documentación.
    - style: formato, espacios, comas (sin cambio de lógica).
    - refactor: reestructurar código sin cambiar lo que hace.
    - test: añadir o corregir pruebas.
    - chore: tareas de mantenimiento (dependencias, configuración).

- Cortos pero claros. Ejemplos: `feat: añadir cabecera responsive`, `fix: corregir enlace roto en nav`, `docs: actualizar README`.
- Commits pequeños y frecuentes. Mejor cinco commits pequeños que uno enorme con todo mezclado.