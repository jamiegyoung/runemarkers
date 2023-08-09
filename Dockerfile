# Development Dockerfile for Next.js and Yarn 3
FROM node:16-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

# Ls the current directory
COPY . .
COPY .yarn ./.yarn
COPY .yarnrc.yml ./

RUN yarn install --immutable

EXPOSE 3000

CMD ["yarn", "dev"]