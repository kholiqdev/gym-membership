FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf
COPY deployment/docker/dockerfiles/nginx/nginx.conf /etc/nginx/conf.d

WORKDIR /usr/src/app
