FROM node:14.17.0 as build

WORKDIR /build

# Cache dependencies in their own layer
COPY package.json package-lock.json ./
RUN npm install
RUN npm install -g @angular/cli

COPY . .
RUN ng build

# Execution Stage
FROM nginx:stable-alpine
#COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /build/dist/webapp /usr/share/nginx/html
