

1. Generate/adjust dashbords.yaml
```
The dashboard ID is available in the dashboard URL. For example, in case the dashboard URL is app/kibana#/dashboard/7fea2930-478e-11e7-b1f0-cb29bac6bf8b?_g=()&_a=(description:'Overview%2..., the dashboard ID is 7fea2930-478e-11e7-b1f0-cb29bac6bf8b.
```

2. Run: 
> go get -v

> go run main.go -yml dashbords.yaml
```
Results will be in ./_meta/kibana/6/dashbord/ folder.
If dashboard already exist, it will be updated.
```


helpful links:
- https://www.elastic.co/guide/en/beats/devguide/6.7/export-dashboards.html
- https://github.com/elastic/beats/tree/master/dev-tools/cmd/dashboards