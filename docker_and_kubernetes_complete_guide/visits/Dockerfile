FROM node:alpine
EXPOSE 8081
WORKDIR "/app"
COPY package.json .
RUN npm install
COPY . .
CMD ["npm", "start"]
