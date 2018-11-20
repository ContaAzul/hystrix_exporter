package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ContaAzul/hystrix_exporter/config"
	"github.com/ContaAzul/hystrix_exporter/hystrix"
	"github.com/ContaAzul/hystrix_exporter/metrics"
)

func init() {
	log.SetHandler(cli.Default)
}

var (
	version    = "dev"
	app        = kingpin.New("hystrix-exporter", "exports hystrix metrics in the prometheus format")
	configFile = app.Flag("config", "config file").Short('c').Default("config.yml").ExistingFile()
	listenAddr = app.Flag("listen-addr", "address to listen to").Default(":9444").String()
	debug      = app.Flag("debug", "show debug logs").Default("false").Bool()
)

var indexTmpl = `
<html>
<head><title>Hystrix Exporter</title></head>
<body>
	<h1>Hystrix Exporter</h1>
	<a href="/metrics">Metrics</a>
	<h3>Clusters being monitored:</h3>
	<ul>
	{{ range .Clusters }}
		<li><a href="{{ .URL }}">{{  .Name }}</a></li>
	{{ end }}
	<ul>
</body>
<html>
`

func main() {
	app.Version(version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	conf, err := config.Parse(*configFile)
	if err != nil {
		log.WithError(err).Fatalf("failed to parse config file: %s", *configFile)
	}

	metrics.MustRegister(prometheus.DefaultRegisterer)
	for _, cluster := range conf.Clusters {
		go read(cluster.URL, cluster.Name)
	}

	var mux = http.NewServeMux()
	var index = template.Must(template.New("index").Parse(indexTmpl))
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := index.Execute(w, &conf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Infof("starting server at %s", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, mux); err != nil {
		log.WithError(err).Fatal("failed to start server")
	}
}

func read(url, cluster string) {
	var log = log.WithField("url", url).WithField("cluster", cluster)
	log.Info("reading")
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Warn("failed to read url")
		time.Sleep(time.Second * 10)
		read(url, cluster)
		return
	}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		log.Debug("new line")
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		if strings.Contains(line, `{"type":"ping"}`) {
			continue
		}
		report(cluster, line)
	}
	log.Warn("stream stop reporting")
	time.Sleep(time.Second * 5)
	read(url, cluster)
}

func report(cluster, line string) {
	data, err := hystrix.Unmarshal(line)
	if err != nil {
		log.WithError(err).Warn("failed to umarshal hystrix json")
		return
	}
	switch data.Type {
	case "HystrixThreadPool":
		metrics.ReportThreadPool(cluster, data)
	case "HystrixCommand":
		metrics.ReportCommand(cluster, data)
	case "meta":
		return
	default:
		log.Warnf("don't know what to do with type '%s'", data.Type)
	}
}
