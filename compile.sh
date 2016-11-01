rm -rf pkg/*
gox \
  -os="linux" \
  -arch="386 amd64" \
  -output "pkg/{{.OS}}_{{.Arch}}_{{.Dir}}"
