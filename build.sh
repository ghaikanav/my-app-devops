#!/bin/bash
set -e  # Exit immediately if any command fails

# Authenticate with Docker
docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD"

# Generate short SHA tag for the Docker image
SHORT_SHA=$(git rev-parse --short HEAD)
IMAGE=kanavghai/my-app:$SHORT_SHA

# Build and push multi-arch Docker image
docker buildx build --platform linux/amd64,linux/arm64 -t "$IMAGE" --push .

# Set Git identity for committing
git config --global user.email "github-actions[bot]@users.noreply.github.com"
git config --global user.name "github-actions"

# Clone the deployment repo using the GH_TOKEN
git clone https://x-access-token:"$GH_TOKEN"@github.com/ghaikanav/my-app-config.git
cd my-app-config

# Update Helm values.yaml with new image tag
yq e '.image.tag = "'"$SHORT_SHA"'"' -i values.yaml

# Commit and push the change
git add .
git commit -m "Update image to $SHORT_SHA"
git push

# Clean up
cd ..
rm -rf my-app-config
