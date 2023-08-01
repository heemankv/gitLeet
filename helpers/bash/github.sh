#!/bin/sh
param1=$1 # commit, push
param2=$2 # commit time
param3=$3 # commit message

case $param1 in

  commit)
    echo -n "Add & Commit"
    git add .
    GIT_COMMITTER_DATE="$param2" git commit --m "$param3" --date "$param2"
    ;;

  push)
    echo -n "Push"
    git push
    ;;

  *)
    echo -n "Error"
    ;;
esac