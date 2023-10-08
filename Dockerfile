# Development Dockerfile for Next.js and pnpm
FROM node:20-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY . .

RUN npm install -g pnpm

RUN pnpm install --frozen-lockfile

EXPOSE 3000

CMD ["pnpm", "dev"]

