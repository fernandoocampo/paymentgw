FROM iron/base
LABEL maintainer=https://github.com/fernandoocampo
ADD bin/paymentgwd-linux-amd64 paymentgwd
EXPOSE 8287
ENTRYPOINT [ "/paymentgwd" ]
