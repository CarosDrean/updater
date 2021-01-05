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
  "routeFrom": "D:\\temp",
  "routeTo": "C:\\drean\\sigesoft"
}
```
