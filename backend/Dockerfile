FROM golang

# aca tenes que declarar las variables de environment que usas dentro de la aplicacion, por ejemplo el user pass de la base de datos
ARG DBUSER
ARG DBPASS
ARG DBHOST

ENV DBUSER=${DBUSER}
ENV DBPASS=${DBPASS}
ENV DBHOST=${DBHOST}


WORKDIR /app

COPY go.* ./
COPY . .
# aca podes hacer 2 cosas. 
# Bien hecho: ves como en GO se toman variables de environment
    # En este caso usarias la declaracion de los ENV

# Medio mal: haces un sed haciendo replace del archivo de config de la app 
# donde pondrias los valores de user, pass, etc. Y reemplazas un placeholder con los parametros de ARG
    #   En este caso usarias la declaracion de los ARG y no harian falta las lineas de ENV.

# seria algo como
RUN sed 's/PLACEHODER_DBUSER/${DBUSER}/' file_que_tiene_la_config.conf

RUN go get ./main/
RUN go build -o bin ./main/

CMD [ "/app/bin" ]


# Este docker se buildearia como
# docker build -e DBUSER="miusuario" -e DBPASS="qwerty" -e DBHOST="localhost"
# pero como este docker no se buildea a mano, estas params pasan en el build del compose
# y el build del compose se refiere a los otros servicios por el nombre del servicio