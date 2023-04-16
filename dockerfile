FROM golang:1.20

RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]