# Golang Azure OpenAI demo app

This is a golang an electronic products content generator using Azure OpenAI.

## Running Locally

- Add the environment files
- Change folders to `cgenerator` and execute: `sh deploy.sh`
- Change forlders to `server` and execute: `go run .`
- Open the browser at: http://localhost:3010

## Running as a container

## Environment Variables

## FrontEnd - React - cgenerator - `.env.local`

```bash
APP_PORT=3000
OPENAI_ENDPOINT=/openAI
```

## Backend - Go - Server - `.env`

```bash
APP_PORT=3010
OPENAI_API_KEY=<KEY>
OPENAI_ENDPOINT=<AZURE_ENDPOINT>
OPENAI_N=1
OPENAI_TEMPERATURE=0.5
OPENAI_MAX_TOKENS=300
```
