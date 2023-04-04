podman build . -t docker.io/am8850/goopenaiapp:latest
#podman login docker.io
podman push am8850/goopenaiapp:latest docker://docker.io/am8850/goopenaiapp:latest
podman run --rm -p 3010:3010 --env-file .env docker.io/am8850/goopenaiapp
