import sys
import re

import google.protobuf.descriptor_pb2 as descriptor_pb2

# Read binary data from stdin
for line in sys.stdin.buffer:
    if b"security.proto" in line:
        # we found the line we are looking for, so we can stop iterating
        break

# regex to find the string between `= { ... }`
m = re.search(r"^.*=\s*\{\s*\"(.*)\"\s*\}.*$", line.decode("utf-8"))
if m:
    data = m.group(1)
else:
    raise Exception("Could not find the string between `= { ... }`")

# fix encoding ... this is a bit hacky, but it works
data = data.encode("latin-1").decode("unicode-escape").encode("latin-1")

# create a new FileDescriptorSet
fds = descriptor_pb2.FileDescriptorSet()
fds.file.append(descriptor_pb2.FileDescriptorProto())

# convert string to bytes
fds.file[0].ParseFromString(data)

# Serialize FileDescriptorSet to bytes
serialized_data = fds.file[0].SerializeToString()

# Write the serialized data to stdout
sys.stdout.buffer.write(serialized_data)