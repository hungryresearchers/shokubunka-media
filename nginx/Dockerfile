FROM nginx:1.15

RUN mkdir -p /var/www/build

COPY nginx.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off;"]