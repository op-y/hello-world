FROM nginx
RUN mkdir -p /etc/nginx/ssl
COPY server.crt /etc/nginx/ssl
COPY server.key /etc/nginx/ssl
COPY nginx.conf /etc/nginx/
CMD ["nginx", "-g", "daemon off;"]
