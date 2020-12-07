#!/bin/bash -e

cd $(dirname $0)

MAX=100
GOTOC=../gotoc
PROTOCMP=./protocmp
PROTOBUF=$HOME/src/protobuf

go build -o protocmp protocmp.go

failures=0
for ((i=1; $i <= $MAX; i=$((i+1)))); do
  if [ ! -f $i.proto ]; then continue; fi
  echo "---[ Test $i ]---" 1>&2
  cat $i.proto | sed 's,^,	,' 1>&2

  $GOTOC --descriptor_only $i.proto > $i.actual
  $PROTOCMP $i.expected $i.actual || {
    echo "==> FAILED expected vs. actual" 1>&2
    failures=$(($failures + 1))
  }

  protoc --descriptor_set_out=_baseline.raw $i.proto
  protoc --decode=google.protobuf.FileDescriptorSet -I $PROTOBUF $PROTOBUF/src/google/protobuf/descriptor.proto \
    < _baseline.raw > $i.baseline
  $PROTOCMP $i.expected $i.baseline || {
    echo "==> BASELINE MISMATCH" 1>&2
    failures=$(($failures + 1))
  }
done

echo "----------" 1>&2
if [ $failures -eq 0 ]; then
  echo "All OK" 1>&2
else
  echo "$failures test failure(s)" 1>&2
fi
echo "----------" 1>&2
exit $failures
