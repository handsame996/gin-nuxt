FROM node:20-slim AS base
ENV PATH="/app/node_modules/.bin:$PATH"
WORKDIR /app

# 复制包文件并安装依赖
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

# 复制应用代码
COPY . .

# 构建生产版本
FROM base AS build
ENV NODE_ENV=production
RUN yarn build

# 运行时环境
FROM node:20-slim AS runtime
WORKDIR /app

# 安装生产依赖
COPY package.json yarn.lock ./
RUN yarn install --production --frozen-lockfile

# 复制构建产物
COPY --from=build /app/.output ./.output

# 暴露端口并启动应用
EXPOSE 8080
CMD ["node", ".output/server/index.mjs"]