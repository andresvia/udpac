package main

import (
	"crypto/tls"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sohlich/elogrus"
	"gopkg.in/olivere/elastic.v3"
	"gopkg.in/urfave/cli.v1"
	"io"
	"net/http"
	"os"
	"time"
)

var pac = `// ESTE ES EL ARCHIVO PAC/WPAD.DAT DE LA UNIVERSIDAD DISTRITAL
// FRANCISCO JOSE DE CALDAS.

// ESTE ARCHIVO REALIZA LA CONFIGURACION AUTOMATICA DE LA MAYORIA DE LOS
// NAVEGADORES PARA MAS INFORMACION VISITAR http://findproxyforurl.com/

function FindProxyForURL(url, host) {

  // LA ESTRATEGIA DE FAIL-OVER EN CASO DE QUE NO SE PUEDA ESTABLECER UNA
  // CONEXION CON EL PROXY ES HACER CONEXION DIRECTA (DIRECT)
  var defaultProxy = "PROXY 10.20.4.15:3128; DIRECT";

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
  // CONECTAR; ES DECIR ACTUARA COMO SI EL PROXY SE HUBIERA DESHABILITADO
  // MANUALMENTE. ESTA CONDICION SE BASA EN LA DIRECCION QUE SE VISITA EN LA
  // BARRA DE DIRECCIONES Y SE EVALUA ANTES DE QUE OCURRA CUALQUIER CONEXION O
  // RESOLUCION DE NOMBRE POR PARTE DEL NAVEGADOR
  // LAS EXCEPCIONES A ESTA REGLA GENERAL DEBERÁN INSERTARSE ANTES DE ESTO

  if (shExpMatch(host, "127.0.0.1") || // LOOPBACK
    shExpMatch(host, "*.local") || // UN DOMINIO MUY COMUN USADO PARA DESARROLLO
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
	app := cli.NewApp()
	app.Action = action
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "listen,l", Value: ":80", EnvVar: "UDPAC_LISTEN"},
		cli.StringFlag{Name: "es-server", Value: "https://zabbix.udistritaloas.edu.co:9443/"},
		cli.StringFlag{Name: "es-user", Value: "udnet"},
		cli.StringFlag{Name: "es-pass", EnvVar: "UDPAC_ES_PASS"},
		cli.BoolFlag{Name: "es-secure"},
		cli.StringFlag{Name: "es-hostname", Value: "udpac"},
		cli.StringFlag{Name: "es-index-prefix", Value: "udnet-"},
		cli.Uint64Flag{Name: "racy-counter-mod", Value: 1000000},
		cli.DurationFlag{Name: "es-index-duration", Value: 23*time.Hour + 59*time.Minute},
		cli.StringFlag{Name: "es-index-layout", Value: "2006.01.02"},
	}
	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	var err error
	var racyCounter uint64

	logctx := log.WithFields(log.Fields{})
	logctx.Info("INICIANDO UDPAC")

	ticker := time.NewTicker(ctx.Duration("es-index-duration"))

	config_es := func() {
		logctx.Info("CONFIGURANDO ELASTICSEARCH")
		client := &elastic.Client{}
		hook := &elogrus.ElasticHook{}
		es_index := ctx.String("es-index-prefix") + time.Now().Format(ctx.String("es-index-layout"))
		if client, err = elastic.NewClient(
			elastic.SetURL(ctx.String("es-server")),
			elastic.SetBasicAuth(ctx.String("es-user"), ctx.String("es-pass")),
			elastic.SetSniff(false),
			elastic.SetMaxRetries(10),
			elastic.SetHttpClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: !ctx.Bool("es-secure")}}})); err != nil {
			logctx.Error("ERROR OBTENIENDO CLIENTE ES: " + err.Error())
		} else if hook, err = elogrus.NewElasticHook(client, ctx.String("es-hostname"), log.DebugLevel, es_index); err != nil {
			logctx.Error("ERROR EN HOOK:" + err.Error())
		} else {
			log.StandardLogger().Hooks = make(log.LevelHooks)
			log.AddHook(hook)
		}
	}

	config_es()
	go func() {
		for {
			select {
			case <-ticker.C:
				config_es()
			}
		}
	}()

	logctx.Info("CONFIGURANDO PROMETHEUS")
	pacServeCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Help: "Número de veces que se ha servido el archivo PAC",
		Name: "pac_serve",
	})
	if err = prometheus.Register(pacServeCounter); err != nil {
		logctx.Error("ERROR CONFIGURANDO PROMETHEUS: " + err.Error())
	}

	logctx.Info("CONFIGURANDO SERVIDOR HTTP")
	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/x-ns-proxy-autoconfig")
		io.WriteString(w, pac)
		pacServeCounter.Inc()
		// esto es una mala idea
		racyCounter = racyCounter + 1
		if racyCounter%ctx.Uint64("racy-counter-mod") == 0 {
			log.WithField("count", racyCounter).Info(fmt.Sprintf("Another %d happy customers", ctx.Uint64("racy-counter-mod")))
		}
	})

	listen := ctx.String("listen")
	logctx.Info("INICIANDO SERVIDOR DE PAC/WPAD.DAT EN " + listen)
	logctx.Fatal(http.ListenAndServe(listen, nil))
}
