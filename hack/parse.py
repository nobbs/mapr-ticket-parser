import sys
from base64 import b64decode

from decrypt import aes_decrypt
from security_pb2 import TicketAndKey

# read ticket from stdin
ticket = sys.stdin.read()

# split ticket into host and secret
host, secret = ticket.split(" ")

# decrypt secret
key: bytes = ("A" * 32).encode()
cipher_text = b64decode(secret)
decrypted_data = aes_decrypt(key, cipher_text)

# parse the ticket into the protobuf
ticket_and_key = TicketAndKey()
ticket_and_key.ParseFromString(decrypted_data)

# print parsed ticket
print(ticket_and_key)
