# UPDATER

Updater es un simple paquete para copiar directorios y sus archivos.

# Instalacion

```
go get
```

# Configuracion

Las configuraciones de los mensajes se encuentran en **constants.go**

```
MessageInit   = "Actualizando Sigesoft..."
FinishError   = "¡Hubo un error!"
FinishSuccess = "¡Actualizado con exito!"
DeleteDirOld  = "Eliminando archivos antiguos..."
```

Para cambiar el origen y el destino, lo debe hacer en **configuration.json**

```json
{
  "routeFrom": "\\\\DESKTOP-QD7QM2Q\\archivos sistema_2\\Tecnologia e Informacion\\sigesoft",
  "routeTo": "C:\\drean\\sigesoft",
  "nameApp": "Sigesoft",

  "routeFrom2": "\\\\DESKTOP-QD7QM2Q\\archivos sistema_2\\Tecnologia e Informacion\\sigesoft-particular",
  "routeTo2": "C:\\drean\\sigesoft-particular",
  "nameApp2": "Sigesoft Particular"
}
```

# Compilacion

Para compilar el proyecto use:

```
go build .
```
Debera acompañar el instalador con el archivo **configuration.json**
