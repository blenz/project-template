FROM node:23-alpine AS base
WORKDIR /frontend
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000

FROM base AS dev
CMD ["npm", "run", "dev"]
