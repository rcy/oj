FROM node:16 as builder
WORKDIR /work
COPY mySubscriptionPlugin.js package-lock.json package.json server.js worker.js .
ADD tasks tasks
ADD react react
RUN npm ci --only=production
CMD ["node", "server.js"]
