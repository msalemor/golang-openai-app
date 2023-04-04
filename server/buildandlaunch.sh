podman build . -t github.io/am8850/goopenaiapp
#podman login github.io
podman push github.io/am8850/goopenaiapp github.io/am8850/goopenaiapp:latest
podman run --rm -p 3010:3010 --env-file .env github.io/am8850/goopenaiapp
