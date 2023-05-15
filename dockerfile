FROM golang:1.20

RUN mkdir /app

##copy seluruh file ke app
COPY . /app

##set direktori utama
WORKDIR /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]