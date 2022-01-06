#!/usr/bin/env sh

# Semantic versioning details: https://semver.org/
# Constants
RELEASE_PATCH="release-patch"
RELEASE_MINOR="release-minor"
RELEASE_MAJOR="release-major"

# Command line arguments.
if [ ${#} -lt 2 ];
then
  echo "Usage: ${0} <semver_file> [ ${RELEASE_PATCH} | ${RELEASE_MINOR} | ${RELEASE_MAJOR} ]"
  exit 1
fi # Check if version file exists 
if [ ! -f ${1} ];
then
  echo "Version file $1 does not exist. Please sepcify a proper version file and try again."
  exit 1
fi

if [ ${2} != ${RELEASE_PATCH} ] && [ ${2} != ${RELEASE_MINOR} ] && [ ${2} != ${RELEASE_MAJOR} ];
then
  echo "Parameter [ ${2} ] should be any of [ ${RELEASE_PATCH} | ${RELEASE_MINOR} | ${RELEASE_MAJOR} ]."
  exit 1
fi

git pull

version_patch=$(cat VERSION | grep -Eo "[0-9]+$")
version_minor=$(cat VERSION | grep -Eo "[0-9]+\.[0-9]+$" | grep -Eo "^[0-9]+")
version_major=$(cat VERSION | grep -Eo "^[0-9]+")

function push_git() {
  # Push the version tags to git
  git tag -a "${version_major}.${version_minor}.${version_patch}" -m "${version_major}.${version_minor}.${version_patch}"
  git push origin "${version_major}.${version_minor}.${version_patch}"
}

echo "Existing version: ${version_major}.${version_minor}.${version_patch}"

if [ ${2} = ${RELEASE_PATCH} ];
then
  let "version_patch=version_patch+1"
elif [ ${2} = ${RELEASE_MINOR} ];
then
  version_patch=0
  let "version_minor=version_minor+1"
elif [ ${2} = ${RELEASE_MAJOR} ];
then
  version_patch=0
  version_minor=0
  let "version_major=version_major+1"
fi

printf "New version: "
echo "${version_major}.${version_minor}.${version_patch}" | tee ${1}

if [ ${2} = ${RELEASE_MAJOR} -o ${2} = ${RELEASE_MINOR} ];
then
  push_git
fi
