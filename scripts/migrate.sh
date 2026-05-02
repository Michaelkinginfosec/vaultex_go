#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_DIR="$SCRIPT_DIR/../migrations"

# Load .env if it exists
if [ -f "$SCRIPT_DIR/../.env" ]; then
  export $(grep -v '^#' "$SCRIPT_DIR/../.env" | xargs)
fi

if [ -z "${DATABASE_URL:-}" ]; then
  echo "Error: DATABASE_URL is not set. Add it to your .env file or export it before running this script."
  exit 1
fi

COMMAND="${1:-up}"

case "$COMMAND" in
  up)
    echo "Applying migrations..."
    migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" up
    echo "Done."
    ;;
  down)
    read -r -p "Are you sure you want to roll back the last migration? [y/N] " confirm
    if [[ "$confirm" != "y" && "$confirm" != "Y" ]]; then
      echo "Aborted."
      exit 0
    fi
    echo "Rolling back last migration..."
    migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" down 1
    echo "Done."
    ;;
  drop)
    echo "Dropping all migrations..."
    migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" drop -f
    echo "Done."
    ;;
  version)
    migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" version
    ;;
  *)
    echo "Usage: $0 [up|down|drop|version]"
    echo "  up       - apply all pending migrations (default)"
    echo "  down     - roll back the last migration"
    echo "  drop     - drop everything (destructive)"
    echo "  version  - show current migration version"
    exit 1
    ;;
esac