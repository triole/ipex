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

_rcmd r -b 127.0.0.1 22
_rcmd r -p ! -b 192.168.166.33 22 33 44
_rcmd r -b 255.255.166.33 155.22/30 99 200/30
_rcmd r -b 255.255.166.33 254+4 1-4 128/31
_rcmd r -s -u -b 255.255.166.33 99+5 99+10 90+12
