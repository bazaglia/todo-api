## Deploy

1. Install kubectl and Google Cloud SDK:

```
brew install kubernetes-cli
brew cask install google-cloud-sdk
```

2. Grant authorization to gcloud:

```
gcloud auth login
```

4. Set gcloud current project ID:

```
gcloud config set project <project-id>
```

4. Get GKE cluster credentials:

```
gcloud container clusters get-credentials backend --zone southamerica-east1-a
```

5. To create a node pool in a running cluster:

```
gcloud container node-pools create default \           
 --cluster backend \
 --machine-type n1-standard-2 \
 --num-nodes 1 \
 --zone southamerica-east1-a \
 --scopes default,storage-rw
```