package main

var DeploymentManifest string = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
        resources:
          requests: {}
          limits:
            cpu: 50Mi
            memory: 200Mi`

var PodManifest string = `apiVersion: v1
kind: Pod
metadata:
  name: <pod-name>
  labels:
    app: <app-name>
spec:
  containers:
  - name: <container-name>
    image: <image-name>
    ports:
    - containerPort: <container-port>
    resources:
      requests: {}
      limits: # Adjust these as needed
        cpu: 50Mi
        memory: 200Mi`
