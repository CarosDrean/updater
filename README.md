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

Para cambiar el origen y el destino, lo debe hacer en **configuration.json**, puede agregar mas campos al arreglo.

```json
{
  "configs": [
    {
      "_id": "1",
      "routeFrom": "\\\\DESKTOP-QD7QM2Q\\archivos sistema_2\\Tecnologia e Informacion\\sigesoft",
      "routeTo": "C:\\drean\\sigesoft",
      "nameApp": "Sigesoft"
    },
    {
      "_id": "2",
      "routeFrom": "\\\\DESKTOP-QD7QM2Q\\archivos sistema_2\\Tecnologia e Informacion\\sigesoft-particular",
      "routeTo": "C:\\drean\\sigesoft-particular",
      "nameApp": "Sigesoft Particular"
    }
  ]
}
```

# Compilacion

Para compilar el proyecto use:

```
go build .
```
Debera acompañar el instalador con el archivo **configuration.json**
