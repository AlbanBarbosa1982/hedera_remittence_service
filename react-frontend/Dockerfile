# Stage 1 - Build the Node.js application
FROM node:latest AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Stage 2 - Set up Nginx
FROM nginx:alpine

# Copy built app from the previous stage
COPY --from=builder /app/build /usr/share/nginx/html

# Copy the Nginx configuration file
COPY ./nginx.conf /etc/nginx/nginx.conf