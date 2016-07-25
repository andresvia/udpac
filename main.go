package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"log"
	"net/http"
	"os"
)

var pac = `// ESTE ES EL ARCHIVO PAC/WPAD.DAT DE LA UNIVERSIDAD DISTRITAL
// FRANCISCO JOSE DE CALDAS.

// ESTE ARCHIVO REALIZA LA CONFIGURACION AUTOMATICA DE LA MAYORIA DE LOS
// NAVEGADORES PARA MAS INFORMACION VISITAR http://findproxyforurl.com/

function FindProxyForURL(url, host) {

  // LA ESTRATEGIA DE FAIL-OVER EN CASO DE QUE NO SE PUEDA ESTABLECER UNA
  // CONEXION CON EL PROXY ES HACER CONEXION DIRECTA (DIRECT)
  var defaultProxy = "PROXY 10.20.4.15:3128; PROXY proxy.udistrital.edu.co:3128; DIRECT";

  // EXCEPCION OFICINA ASESORA DE SISTEMAS
  // LA OFICINA ASESORA DE SISTEMAS TIENE UN DOMINIO EL CUAL ES ATENDIDO POR FUERA
  // DE LA RED LOCAL, ESTE DEBE REDIRIGIRSE A TRAVES DEL PROXY
  if (shExpMatch(host, "*.portaloas.udistrital.edu.co")) {
    return defaultProxy;
  }

  // CUALQUIER PETICION DESDE EL NAVEGADOR A DOMINIOS QUE TERMINEN EN
  // .udistrital.edu.co O .udistritaloas.edu.co O .local[domain] NO SE ENVIARA A
  // TRAVES DEL PROXY, QUIERE DECIR QUE CONECTARAN DIRECTAMENTE A TRAVES DE LA
  // RED LOCAL, SIEMPRE Y CUANDO LOS CLIENTES PUEDAN RESOLVER ESTOS NOMBRES CON
  // CONFIGURACION ACTUAL DE DNS Y QUE EXISTA UNA RUTA EN LA RED QUE LOS PUEDA
  // CONECTAR ES; DECIR ACTUARA COMO SI EL PROXY SE HUBIERA DESHABILITADO
  // MANUALMENTE. ESTA CONDICION SE BASA EN LA DIRECCION QUE SE VISITA EN LA
  // BARRA DE DIRECCIONES Y SE EVALUA ANTES DE QUE OCURRA CUALQUIER CONEXION O
  // RESOLUCION DE NOMBRE POR PARTE DEL NAVEGADOR
  // LAS EXCEPCIONES A ESTA REGLA GENERAL DEBERÁN INSERTARSE ANTES DE ESTO

  if (shExpMatch(host, "*.local") || // UN DOMINIO MUY COMUN USADO PARA DESARROLLO
    shExpMatch(host, "*.localdomain") || // UN DOMINIO COMUN USADO POR DEFECTO
    shExpMatch(host, "*.nip.io") || // NIP ES UN CLON DE XIP
    shExpMatch(host, "*.xip.io") || // XIP UN SERVICIO QUE MAPEA DOMINIOS COMODIN
    shExpMatch(host, "*.udistrital.edu.co") || // DOMINIO OFICIAL DE LA UNIVERSIDAD
    shExpMatch(host, "*.udistritaloas.edu.co") // DOMINIO NO OFICIAL DE LA OFICINA ASESORA DE SISTEMAS
  ) {
    return "DIRECT";
  }

  // CUALQUIER PETICION QUE RESUELVA EL NOMBRE DNS A UNA IP LOCAL (ES DECIR
  // CUALQUIER IP DEL ESPACIO PRIVADO DE IPV4) NO SERA ENVIADA A TRAVES DEL PROXY
  // LAS EXCEPCIONES A ESTA REGLA GENERAL DEBERÁN INSERTARSE ANTES DE ESTO

  var resolved_ip = dnsResolve(host);
  if (isInNet(resolved_ip, "10.0.0.0", "255.0.0.0") || // UNA RED DE CLASE A
    isInNet(resolved_ip, "172.16.0.0", "255.240.0.0") || // 16 REDES DE CLASE B
    isInNet(resolved_ip, "192.168.0.0", "255.255.0.0") || // 256 REDES DE CLASE C
    isInNet(resolved_ip, "169.254.0.0", "255.255.0.0") || // LINK-LOCAL RFC3927
    isInNet(resolved_ip, "192.0.2.0", "255.255.255.0") || // TEST-NET RFC3330
    isInNet(resolved_ip, "127.0.0.0", "255.0.0.0") // LOOPBACK
  ) {
    return "DIRECT";
  }
  // CUALQUIER PETICION ENVIADA A UN HOST NO CALIFICADO (POR EJEMPLO http://www/ O
  // https://mail/) NO SERA ENVIADA A TRAVES DEL PROXY

  if (isPlainHostName(host)) {
    return "DIRECT";
  }

  // SI NINGUNA DE LAS ANTERIORES CONDICIONES ANTERIORES APLICAN SE USARA EL PROXY
  // POR DEFECTO
  return defaultProxy;

}
// FIN`

func main() {
	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/x-ns-proxy-autoconfig")
		io.WriteString(w, pac)
	})
	listen := os.Getenv("UDPAC_LISTEN")
	if listen == "" {
		listen = ":80"
	}
	log.Println("INICIANDO SERVIDOR DE PAC/WPAD.DAT EN " + listen)
	log.Println("UTILIZAR VARIABLE DE ENTORNO UDPAC_LISTEN PARA CAMBIAR EL PUERTO")
	log.Fatal(http.ListenAndServe(listen, nil))
}
