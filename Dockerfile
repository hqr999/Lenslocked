FROM node:latest AS tailwind-builder
WORKDIR /src 
COPY ./tailwind/package.json .
COPY ./tailwind/package-lock.json .
RUN npm install

COPY ./templates /templates     
COPY ./tailwind/style.css /src/style.css
RUN npx @tailwindcss/cli -i /src/style.css -o /styles.css 



FROM golang:alpine AS builder
WORKDIR /app 
COPY go.mod go.sum ./  
RUN go mod download 
COPY . . 
RUN go build -v -o ./server ./cmd/server/ 



FROM alpine
WORKDIR /app 
COPY ./assets ./assets      
COPY .env .env  
COPY --from=builder /app/server ./server 
COPY --from=tailwind-builder /styles.css /app/assets/styles.css
CMD ./server


