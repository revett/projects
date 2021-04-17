FROM node:15.14.0-slim

RUN npm -g install aglio@2.3.0

ENTRYPOINT ["aglio", "-s", "-p", "5687", "-h", "0.0.0.0"]
