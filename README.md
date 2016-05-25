# UD PAC y WPAD.DAT

## Pruebas

Ejecute

 - `make`
 - `vagrant up` o `vagrant provision`

Y habrá un servidor de pruebas en:

 - http://127.0.0.1:9645/wpad.dat (probar paquete debian)
 - http://127.0.0.1:9646/wpad.dat (probar paquete centos)

Utilice estas instrucciones para configurar el proxy automáticamente en su navegador.

 - "Configuración del navegador" => "Configuración de conexión" => "URL para la configuración automática del proxy"

## Producción

Los paquetes .deb y .rpm se encuentran en https://github.com/udistrital/udpac/releases/latest.

![WPAD por DNS](http://findproxyforurl.com/wp-content/uploads/wpaddns_diagram2.png)
<sub>WPAD por DNS - &copy; http://findproxyforurl.com/</sub>

Instalar el paquete correspondiente en los servidores:

  - http://wpad.udistrital.edu.co:80/wpad.dat
  - http://wpad.udistritaloas.edu.co:80/wpad.dat

Los cuales podrían ser el mismo mismo servidor con dos registros de DNS tipo `A`.

Para algunos navegadores esto es suficiente sin embargo para ampliar la cobertura puede establecer la URL de autoconfiguración a través de:

 - Políticas de grupo (equipos en directorio activo)
 - O configuración por DHCP

![WPAD por DHCP](http://findproxyforurl.com/wp-content/uploads/wpad_diagram1.png)
<sub>WPAD por DHCP - &copy; http://findproxyforurl.com/</sub>

## Más información

 - http://findproxyforurl.com/
