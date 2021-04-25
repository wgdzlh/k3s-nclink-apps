PREFIX="nanocpu/nclink-"
VERSION="1.0.0"
PLATFORM="linux/amd64,linux/arm64,linux/arm/v7"

docker buildx build -f ./config-distribute/Dockerfile -t ${PREFIX}configdist:${VERSION} --platform $PLATFORM --push .
docker buildx build -f ./adapter-simulator/Dockerfile -t ${PREFIX}adapter:${VERSION} --platform $PLATFORM --push .
docker buildx build -f ./model-manage-backend/Dockerfile -t ${PREFIX}modelmanage-backend:${VERSION} --platform $PLATFORM --push .

# multi-platform emulation build not supported by Node.js yet.
cd model-manage-frontend
docker buildx build -t ${PREFIX}modelmanage-frontend:${VERSION} --platform linux/amd64 --push .
