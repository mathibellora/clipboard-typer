import time
import threading
import pyperclip
from pynput import keyboard
from pynput.keyboard import Controller

try:
    import pystray
    from PIL import Image, ImageDraw
    HAS_TRAY = True
except ImportError:
    HAS_TRAY = False

HOTKEY = '<ctrl>+<shift>+v'
DELAY_BETWEEN_CHARS = 0.04  # segundos entre caracteres (ajustar si la VM pierde chars)

kb = Controller()
typing = False


def make_tray_icon():
    img = Image.new('RGB', (64, 64), color='#1a1a2e')
    d = ImageDraw.Draw(img)
    d.rectangle([12, 20, 52, 44], fill='#e0e0e0')
    d.rectangle([20, 28, 44, 36], fill='#1a1a2e')
    return img


def type_clipboard():
    global typing
    if typing:
        return
    typing = True
    try:
        text = pyperclip.paste()
        if not text:
            return
        time.sleep(0.1)
        for char in text:
            kb.type(char)
            time.sleep(DELAY_BETWEEN_CHARS)
    finally:
        typing = False


def on_activate():
    threading.Thread(target=type_clipboard, daemon=True).start()


def run_hotkey_listener():
    with keyboard.GlobalHotKeys({HOTKEY: on_activate}) as listener:
        listener.join()


def main():
    listener_thread = threading.Thread(target=run_hotkey_listener, daemon=True)
    listener_thread.start()

    if HAS_TRAY:
        icon = pystray.Icon(
            'clipboard-typer',
            make_tray_icon(),
            'Clipboard Typer',
            menu=pystray.Menu(
                pystray.MenuItem('Clipboard Typer', None, enabled=False),
                pystray.MenuItem(f'Atajo: Ctrl+Shift+V', None, enabled=False),
                pystray.Menu.SEPARATOR,
                pystray.MenuItem('Salir', lambda icon, _: icon.stop()),
            )
        )
        icon.run()
    else:
        print('Clipboard Typer corriendo. Presioná Ctrl+Shift+V para escribir el portapapeles.')
        print('Ctrl+C para salir.')
        listener_thread.join()


if __name__ == '__main__':
    main()
