// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/url"
// 	"time"

// 	"github.com/elastic/beats/libbeat/common"
// 	"github.com/elastic/beats/libbeat/kibana"
// 	"github.com/elastic/beats/libbeat/logp"
// )

// var importAPI = "/api/kibana/dashboards/import"

// type KibanaLoader struct {
// 	client       *kibana.Client
// 	config       *Config
// 	version      common.Version
// 	hostname     string
// 	msgOutputter MessageOutputter
// }

// // NewKibanaLoader creates a new loader to load Kibana files
// func NewKibanaLoader(ctx context.Context, cfg *common.Config, dashboardsConfig *Config, hostname string, msgOutputter MessageOutputter) (*KibanaLoader, error) {

// 	if cfg == nil || !cfg.Enabled() {
// 		return nil, fmt.Errorf("Kibana is not configured or enabled")
// 	}

// 	client, err := getKibanaClient(ctx, cfg, dashboardsConfig.Retry, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error creating Kibana client: %v", err)
// 	}

// 	loader := KibanaLoader{
// 		client:       client,
// 		config:       dashboardsConfig,
// 		version:      client.GetVersion(),
// 		hostname:     hostname,
// 		msgOutputter: msgOutputter,
// 	}

// 	version := client.GetVersion()
// 	loader.statusMsg("Initialize the Kibana %s loader", version.String())

// 	return &loader, nil
// }

// func getKibanaClient(ctx context.Context, cfg *common.Config, retryCfg *Retry, retryAttempt uint) (*kibana.Client, error) {
// 	client, err := kibana.NewKibanaClient(cfg)
// 	if err != nil {
// 		if retryCfg.Enabled && (retryCfg.Maximum == 0 || retryCfg.Maximum > retryAttempt) {
// 			select {
// 			case <-ctx.Done():
// 				return nil, err
// 			case <-time.After(retryCfg.Interval):
// 				return getKibanaClient(ctx, cfg, retryCfg, retryAttempt+1)
// 			}
// 		}
// 		return nil, fmt.Errorf("Error creating Kibana client: %v", err)
// 	}
// 	return client, nil
// }

// // ImportDashboard imports the dashboard file
// func (loader KibanaLoader) ImportDashboard(file string) error {
// 	params := url.Values{}
// 	params.Set("force", "true")            //overwrite the existing dashboards
// 	params.Add("exclude", "index-pattern") //don't import the index pattern from the dashboards

// 	// read json file
// 	reader, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return fmt.Errorf("fail to read dashboard from file %s: %v", file, err)
// 	}
// 	var content common.MapStr
// 	err = json.Unmarshal(reader, &content)
// 	if err != nil {
// 		return fmt.Errorf("fail to unmarshal the dashboard content from file %s: %v", file, err)
// 	}

// 	return loader.client.ImportJSON(importAPI, params, content)
// }

// func (loader KibanaLoader) Close() error {
// 	return loader.client.Close()
// }

// func (loader KibanaLoader) statusMsg(msg string, a ...interface{}) {
// 	if loader.msgOutputter != nil {
// 		loader.msgOutputter(msg, a...)
// 	} else {
// 		logp.Debug("dashboards", msg, a...)
// 	}
// }

// // func main() {
// // 	kibanaURL := flag.String("kibana", "http://localhost:5601", "Kibana URL")
// // 	spaceID := flag.String("space-id", "", "Space ID")
// // 	dashboard := flag.String("dashboard", "", "Dashboard ID")
// // 	fileOutput := flag.String("output", "output.json", "Output file")
// // 	ymlFile := flag.String("yml", "", "Path to the module.yml file containing the dashboards")
// // 	flag.BoolVar(&indexPattern, "indexPattern", false, "include index-pattern in output")
// // 	flag.BoolVar(&quiet, "quiet", false, "be quiet")

// // 	flag.Parse()
// // 	log.SetFlags(0)

// // 	u, err := url.Parse(*kibanaURL)
// // 	if err != nil {
// // 		log.Fatalf("Error parsing Kibana URL: %v", err)
// // 	}

// // 	var user, pass string
// // 	if u.User != nil {
// // 		user = u.User.Username()
// // 		pass, _ = u.User.Password()
// // 	}
// // 	client, err := kibana.NewClientWithConfig(&kibana.ClientConfig{
// // 		Protocol: u.Scheme,
// // 		Host:     u.Host,
// // 		Username: user,
// // 		Password: pass,
// // 		Path:     u.Path,
// // 		SpaceID:  *spaceID,
// // 		Timeout:  kibanaTimeout,
// // 	})
// // 	if err != nil {
// // 		log.Fatalf("Error while connecting to Kibana: %v", err)
// // 	}
// // }
