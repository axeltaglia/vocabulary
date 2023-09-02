FROM golang

WORKDIR /app

COPY go.* ./
COPY . .

RUN go get ./main/
RUN go build -o bin ./main/

CMD [ "/app/bin" ]
