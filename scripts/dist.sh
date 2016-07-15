#!/usr/bin/env bash
set -e

# Get the version from the command line.
VERSION=$1
if [ -z $VERSION ]; then
    echo "Please specify a version."
    exit 1
fi

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that dir because we expect that.
cd $DIR

# Generate the tag.
if [ -z $NOTAG ]; then
  echo "==> Tagging..."
  git commit --allow-empty -m "Release v$VERSION"
  git tag -a "v$VERSION" -m "Version $VERSION" master
fi

# Zip all the files
rm -rf ./pkg/dist
mkdir -p ./pkg/dist
for FILENAME in $(find ./pkg -mindepth 1 -maxdepth 1 -type f); do
  FILENAME=$(basename $FILENAME)
  cp ./pkg/${FILENAME} ./pkg/dist/health_${VERSION}_${FILENAME}
done

# Make the checksums.
pushd ./pkg/dist
shasum -a256 * > ./health_${VERSION}_SHA256SUMS
popd

exit 0
