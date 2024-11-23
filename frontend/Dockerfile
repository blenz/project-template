FROM node:23-alpine AS base
WORKDIR /frontend
EXPOSE 3000

FROM base AS dev
ENV NODE_NO_WARNINGS 1
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["npm install && npm run dev"]
