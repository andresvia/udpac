# UD PAC y WPAD.DAT

**Los paquetes .deb y .rpm se encuentran en https://github.com/andresvia/udpac/releases/latest.**

## Pruebas

### En desarrollo

```
go get github.com/andresvia/udpac
UDPAC_LISTEN=:9645 $GOPATH/bin/udpac
```

### Pruebas de CI

Ejecute

 - `make`
 - `vagrant up` o `vagrant provision`

Esto creará dos servidores virtuales de pruebas en:

Debian:

 - `curl -v http://127.0.0.1:9645/`

CentOS:

 - `curl -v http://127.0.0.1:9646/`

## Configuración de navegador

Utilice estas instrucciones para configurar el proxy automáticamente en su navegador.

 - "Configuración del navegador" => "Configuración de conexión" => "URL para la configuración automática del proxy" => (Aquí configure la URL de prueba, http://127.0.0.1:9645/ o http://127.0.0.1:9646/)

## Paso a producción

Los paquetes .deb y .rpm se encuentran en https://github.com/andresvia/udpac/releases/latest.

WPAD funciona automáticamente por DNS.

<sub>WPAD por DNS - &copy; http://findproxyforurl.com/</sub>

![WPAD por DNS](http://findproxyforurl.com/wp-content/uploads/wpaddns_diagram2.png)

Instalar el paquete correspondiente (`.deb`, .`rpm` o binario `udpac_linux_amd64`) en los servidores:

  - wpad.udistrital.edu.co
   - La URL de PAC será: http://wpad.udistrital.edu.co:80/
  - wpad.udistritaloas.edu.co
   - La URL de PAC será: http://wpad.udistritaloas.edu.co:80/

Los cuales de hecho podrían ser el mismo mismo servidor con dos registros de DNS tipo `A`. Es conveniente utilizar la misma técnica de "split-brain DNS" que tiene el proxy (proxy.udistrital.edu.co) actualmente.

Este servidor responde con el archivo PAC en cualquier ruta por lo tanto:

```
http://wpad.udistrital.edu.co:80/wpad.dat == http://wpad.udistrital.edu.co:80/ == http://wpad.udistrital.edu.co:80/otra_cosa
```

Todas las rutas responden con el archivo PAC, con la excepción de la ruta especial `/metrics` la cual retorna información de desempeño del servidor.

Para algunos navegadores esto es suficiente sin embargo para ampliar la cobertura puede establecer la URL de autoconfiguración a través de:

 - Políticas de grupo (equipos en directorio activo)
 - O configuración por DHCP (equipos conectados a la red, que configuran su IP automáticamente)

<sub>WPAD por DHCP - &copy; http://findproxyforurl.com/</sub>

![WPAD por DHCP](http://findproxyforurl.com/wp-content/uploads/wpad_diagram1.png)

## Más información

 - http://findproxyforurl.com/
