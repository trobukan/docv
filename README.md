# Docv (Visualizador de documentación CLI)

`docv` es una herramienta ligera de linea de comandos (CLI) desarrolada en **Go** que permite explorar y visualizar la documentación local del sistema (ubicada en `/usr/share/doc`) de forma rápida a través de un servidor HTTP local y el navegador web.

## Requisitos
- **Go** (versión 1.22 o superior recomendada)
- Sistema operativo Linux (con la ruta `/usr/share/doc` disponible)
- **xdg-utils** (necesario para que `xdg-open` abra el navegador)

## Compilación
```bash
make
```
o 
```bash
go build -o docv ./cmd/docv/main.go
```

# Uso
```bash
./docv flatpak 
```
