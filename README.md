# Clipboard Typer

Escribe el contenido del portapapeles carácter por carácter, como si lo estuvieras tecleando. Útil para ingresar contraseñas en máquinas virtuales donde el copy-paste no funciona.

## Descargar

Ir a [Releases](../../releases/latest) y bajar el archivo para tu sistema:

| Sistema | Archivo |
|---------|---------|
| Windows | `clipboard-typer-windows.exe` |
| Mac     | `clipboard-typer-mac` |
| Linux   | `clipboard-typer-linux` |

## Uso

1. Ejecutar el archivo
2. Aparece un ícono en la barra del sistema — eso significa que está corriendo
3. Copiar la contraseña al portapapeles (`Ctrl+C` / `Cmd+C`)
4. Conectarse a la VM y hacer click en el campo donde querés escribir
5. Presionar **Ctrl+Shift+V** → escribe solo

Para cerrar: click derecho en el ícono → Salir.

## Notas por sistema

**Windows** — funciona directo, sin instalar nada ni necesitar permisos de administrador.

**Mac** — la primera vez va a pedir aprobar permisos de Accesibilidad:
1. Aparece un aviso del sistema
2. Ir a Ajustes del Sistema → Privacidad y Seguridad → Accesibilidad
3. Activar el permiso para Clipboard Typer
4. Volver a ejecutar

**Linux** — requiere `xdotool` instalado (`sudo apt install xdotool` o equivalente). Funciona en entornos con X11; en Wayland puede no funcionar.

## Configuración

Si la VM pierde caracteres (pasa en conexiones lentas), se puede ajustar la velocidad de tipeo editando el archivo `main.py` antes de compilar:

```python
DELAY_BETWEEN_CHARS = 0.04  # aumentar a 0.08 o más si hay problemas
```
