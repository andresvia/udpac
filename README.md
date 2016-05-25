# UD PAC y WPAD.DAT

## Pruebas

Hay un servidor de pruebas en:

 - http://
 - http://

Utilice esta dirección para configurar el proxy automáticamente en su navegador. Configuración del navegador -> Configuración de conexión -> URL para la configuración automática del proxy.

## Producción

Instalar esto en los servidores:

 - http://wpad.udistrital.edu.co:80/wpad.dat
 - http://wpad.udistritaloas.edu.co:80/wpad.dat

Nota: Es posible que sea el mismo servidor con dos registros de DNS A que apuntan a la misma dirección IP.

Hay multiples maneras de configurar estos servidores de manera automatizada bien sea con políticas de grupo o con configuraciones de servidores DHCP.

## Más información

 - http://findproxyforurl.com/
