
#
**Pre-requirements:**
#
1. Let's assume kibana is accessible on localhost:5601

Example:
> jx ns jx-staging

> kubectl  port-forward web-kibana-XXXXX 5601:5601

#

**Export dashbords:**
#

1. Generate/adjust dashbords.yaml
```
The dashboard ID is available in the dashboard URL. For example, in case the dashboard URL is app/kibana#/dashboard/7fea2930-478e-11e7-b1f0-cb29bac6bf8b?_g=()&_a=(description:'Overview%2..., the dashboard ID is 7fea2930-478e-11e7-b1f0-cb29bac6bf8b.
```

2. Run: 
> go get -v

> go run export_dashboards.go -yml dashbords.yaml
```
Results will be in ./_meta/kibana/6/dashbord/* folder.
If dashboard already exist, it will be updated.
```

#
**Import dashbord (example for export.json):**
#
2. Run: 
> go get -v

> go run export_dashboards.go -json _meta/kibana/6/dashboard/export.json

#
helpful links:
- https://www.elastic.co/guide/en/beats/devguide/6.7/export-dashboards.html
- https://github.com/elastic/kibana/pull/10858
- https://github.com/elastic/beats/issues/4409
- https://github.com/elastic/beats/tree/master/dev-tools/cmd/dashboards (or it's better to use branch: 6.4)
- https://github.com/elastic/beats/tree/master/libbeat/dashboards (or it's better to use branch: 6.4)


#
**Another method (using curl):**
#
```
1. Get id of dashbord: ID_DASH
2. curl -k -XGET "http://localhost:5601/api/kibana/dashboards/export?dashboard=${ID_DASH}" > export.json
3. curl -k -XPOST -H "Content-Type: application/json" -H "kbn-xsrf: true" -d @export.json http://localhost:5601/api/kibana/dashboards/import
```