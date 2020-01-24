FROM node:alpine
WORKDIR /usr/app
# Copy everything from current dir to current working dir in image
COPY ./package.json ./
RUN npm install
COPY ./ ./
CMD ["npm", "start"]
