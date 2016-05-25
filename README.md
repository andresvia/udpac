# UD PAC y WPAD.DAT

## Pruebas

Hay un servidor de pruebas en:

 - http://
 - http://

Utilice estas direcciones para configurar el proxy automáticamente en su navegador.

 - "Configuración del navegador" => "Configuración de conexión" => "URL para la configuración automática del proxy"

## Producción

Instalar esto en los servidores:

 - http://wpad.udistrital.edu.co:80/wpad.dat
 - http://wpad.udistritaloas.edu.co:80/wpad.dat

Los cuales pueden ser el mismo mismo servidor con dos registros de DNS "A".

Los paquetes .deb y .rpm se encuentran en https://github.com/andresvia/udpac/releases/latest.

Para algunos navegadores esto es suficiente, pero para abarcar la mayor cantidad posible de clientes se pueden considerar dos opciones más.

 - Políticas de grupo
 - Configuración de servidor DHCP

## Más información

 - http://findproxyforurl.com/
