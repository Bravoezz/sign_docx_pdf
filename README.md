# Proyecto de Firma de Documentos (PDF y DOCX)

Este proyecto permite la conversión y firma digital de documentos PDF y DOCX utilizando Go, LibreOffice y Docker. La aplicación se ejecuta en un entorno contenedorizado y soporta la conversión de archivos DOCX a PDF, además de la firma electrónica de documentos.

## Tecnologías utilizadas

- **Golang**: Lenguaje principal de desarrollo.
- **LibreOffice**: Utilizado para convertir archivos DOCX a PDF.
- **Docker**: Contenedoriza la aplicación para facilitar la ejecución en distintos entornos.
- **Alpine Linux**: Base del contenedor para minimizar el tamaño de la imagen.

## Instalación y uso

### 0. Agrega tus archivos y modifica program.go con los nombres de tus archivos

```
        const (
		docPath    = "docx-file-sign.docx"
		outputPath = "output-sign.docx"
		imagePath  = "sign-image.png"
		searchText = "text to search"
	)

	signerDoc := signer.NewSigner(&signer.SignOp{
		InputPath:     path.Join(basePath, "assets", docPath),
		OutputPath:    path.Join(basePath, "store", outputPath),
		SignaturePath: path.Join(basePath, "assets", imagePath),
		SearchText:    searchText,
	})

```

### 1. Construcción de la imagen Docker

Para compilar la aplicación y crear la imagen de Docker, ejecuta:

\`\`\`sh
docker build -t signer-app .
\`\`\`

### 2. Ejecución con Docker

Para ejecutar la aplicación en un contenedor:

\`\`\`sh
docker run --rm -v $(pwd)/assets:/app/assets -v $(pwd)/store:/app/store signer-app
\`\`\`

### 3. Ejecución con Docker Compose

También puedes usar **Docker Compose** para administrar el contenedor de manera más sencilla. Usa el siguiente archivo \`docker-compose.yml\`:

\`\`\`yaml
services:
signer:
build: .
command: /app/bin/signer.app
volumes:
- ./store:/app/store
- ./assets:/app/assets
\`\`\`

Para iniciar el servicio (NO OLVIDES PONER TUS ARCHIVOS A FIRMAR EN /assets):

\`\`\`sh
docker compose up --build
\`\`\`

## Conversión de archivos DOCX a PDF

La conversión se realiza mediante LibreOffice en modo headless:

\`\`\`sh
libreoffice --headless --convert-to pdf /ruta/al/archivo.docx --outdir /ruta/de/salida/
\`\`\`

## Problemas comunes y soluciones

### 1. **El texto en el PDF generado aparece ilegible (caracteres extraños)**
**Solución:** Asegurar que las fuentes adecuadas estén instaladas en el contenedor:
\`\`\`dockerfile
RUN apt-get update && apt-get install -y fonts-dejavu fonts-liberation ttf-mscorefonts-installer
\`\`\`

### 2. **No se encuentra el comando \`libreoffice\` en el contenedor**
**Solución:** Verificar que LibreOffice esté correctamente instalado:
\`\`\`sh
apk add libreoffice
\`\`\`

### 3. **El PDF no se puede leer correctamente en el código Go**
**Solución:** Asegurar que la conversión se realice con fuentes embebidas:
\`\`\`sh
libreoffice --headless --convert-to pdf --infilter="writer_pdf_Export:EmbedStandardFonts=true" archivo.docx
\`\`\`

## Licencia

Este proyecto está bajo la licencia MIT.
