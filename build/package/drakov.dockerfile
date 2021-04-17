FROM node:15.14.0-slim

RUN npm -g install drakov@2.0.1

CMD ["-f", "/api/*.apib", "--watch"]
ENTRYPOINT ["drakov", "--public=\"true\"", "-p", "4587"]
