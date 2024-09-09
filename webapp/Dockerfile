FROM docker.io/library/node:alpine as builder

RUN mkdir -p /code

# Copy over application and build it
WORKDIR /code

COPY ./webapp/package.json .
RUN yarn

COPY ./.env .
COPY ./webapp/ .
RUN yarn build

FROM docker.io/library/nginx:alpine

# Copy over built files and host it.
COPY --from=builder /code/dist/* /usr/share/nginx/html
COPY --from=builder /code/nginx.conf /etc/nginx/nginx.conf
RUN mkdir -p /usr/share/nginx/html/assets && \
    mv /usr/share/nginx/html/index-* /usr/share/nginx/html/assets/

CMD ["nginx", "-g", "daemon off;"]
