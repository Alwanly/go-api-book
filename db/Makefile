DB_URI := $(shell cat '.env.db' | grep 'DB_URI' | cut -d '=' -f2-)


# ---- db migrations
new:
	atlas migrate new ${NAME} --env local --var url=${DB_URI}

migration:
	atlas migrate diff ${NAME} --env local --var url=${DB_URI}

diff:
	atlas migrate diff --env local --format '{{ sql . "  " }}' --var url=${DB_URI}

status:
	atlas migrate status --env local --var url=${DB_URI}

up:
	atlas migrate apply --env local --allow-dirty --var url=${DB_URI}

reset:
	atlas schema clean --env local --var url=${DB_URI}
	atlas migrate apply --env local --var url=${DB_URI}

hash:
	atlas migrate hash

# ---- dependencies
install:
	curl -sSf https://atlasgo.sh | sh
	go mod download

.PHONY: migration diff status up reset hash inspect install
