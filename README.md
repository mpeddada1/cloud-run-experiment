# cloud-run-experiment

An experiment that demonstrates building a Cloud Run service that accepts a string an input and 
returns the length of the string as a response. 

## Steps to build the service 

1. Call the following command to build the docker image for the service:

```
gcloud builds submit --tag us-central1-docker.pkg.dev/mpeddada-test/cloud-run-experiment/string-length-service
```

This builds and pushes the image to artifact registry in the `mpeddada-test` project. 

2. Deploy the Cloud Run service using the docker image build in Step 1:

```
gcloud run deploy string-length-service --image us-central1-docker.pkg.dev/mpeddada-test/cloud-run-experiment/string-length-service --platform managed --region us-central1
```

This returns the service url as well.

```
Service URL: https://string-length-service-931854040550.us-central1.run.app
```

## Invoke the Cloud Run service with curl

Call the following command:
```
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $(gcloud auth print-identity-token)" -d '{"input": "hello"}' https://string-length-service-931854040550.us-central1.run.app
{"length":5}
```

Not including `-H "Authorization: Bearer $(gcloud auth print-identity-token)"` will results in:

```
<html><head>
<meta http-equiv="content-type" content="text/html;charset=utf-8">
<title>403 Forbidden</title>
</head>
<body text=#000000 bgcolor=#ffffff>
<h1>Error: Forbidden</h1>
<h2>Your client does not have permission to get URL <code>/</code> from this server.</h2>
<h2></h2>
</body></html>
```