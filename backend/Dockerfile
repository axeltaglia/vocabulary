FROM golang:alpine3.18

WORKDIR /app

COPY go.* ./
COPY . .

RUN go get ./main/apiApp/
RUN go build -o bin ./main/apiApp/

CMD [ "/app/bin" ]


# Este docker se buildearia como
# docker build -e DBUSER="miusuario" -e DBPASS="qwerty" -e DBHOST="localhost"
# pero como este docker no se buildea a mano, estas params pasan en el build del compose
# y el build del compose se refiere a los otros servicios por el nombre del servicio