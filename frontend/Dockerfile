FROM node:23-alpine AS base
WORKDIR /frontend
EXPOSE 3000

FROM base AS dev
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["npm install && npm run dev"]
