# UD PAC y WPAD.DAT

## Pruebas

Ejecute

 - `vagrant up`

Habrá un servidor de pruebas en:

 - http://<tu ip>:9645/wpad.dat (debian)
 - http://<tu ip>:9646/wpad.dat (centos)

Utilice estas instrucciones para configurar el proxy automáticamente en su navegador.

 - "Configuración del navegador" => "Configuración de conexión" => "URL para la configuración automática del proxy"

## Producción

Instalar esto en los servidores:

 - http://wpad.udistrital.edu.co:80/wpad.dat
 - http://wpad.udistritaloas.edu.co:80/wpad.dat

![WPAD](http://findproxyforurl.com/wp-content/uploads/wpaddns_diagram2.png)

Los cuales podrían ser el mismo mismo servidor con dos registros de DNS "A".

Los paquetes .deb y .rpm se encuentran en https://github.com/udistrital/udpac/releases/latest.

Para algunos navegadores esto es suficiente establezca la URL de autoconfiguración a través de:

 - Políticas de grupo
 - Configuración de servidor DHCP

## Más información

 - http://findproxyforurl.com/
