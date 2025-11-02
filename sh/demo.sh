#!/bin/bash
scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir="${scriptdir%/*}"

cd "${basedir}"

_rcmd() {
  cmd=${@}
  out="$(echo "${cmd}" | sed "s|^r|ipex|g")"
  echo "# ${out}"
  eval ${cmd}
  echo ""
}

_rcmd r -b 255.255.166.33 22
_rcmd r -b 255.255.166.33 22 33 44
_rcmd r -b 255.255.166.33 155.22/30 99 200/30
