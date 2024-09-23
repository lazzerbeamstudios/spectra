## GCP

**commands**

    gcloud auth login
    gcloud auth revoke

    gcloud auth list
    gcloud config set account [account]

    gcloud projects list
    gcloud config set project [project]

**docker**

    gcloud auth configure-docker us-central1-docker.pkg.dev

**terraform**

    gcloud auth application-default login
    gcloud auth application-default revoke

**key-json**

    cat gcp.json | base64
