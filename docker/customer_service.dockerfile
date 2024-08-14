
FROM alpine:latest

RUN mkdir /app

COPY ./CustomerService /app

CMD ["/app/CustomerService"]