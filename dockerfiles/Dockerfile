FROM ubuntu

COPY ./dist ./demo 
WORKDIR /demo

EXPOSE 8080
ENTRYPOINT [ "./server" ]