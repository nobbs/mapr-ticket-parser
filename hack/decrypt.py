import sys
from base64 import b64decode

from cryptography.hazmat.primitives.ciphers.aead import AESGCM


def aes_decrypt(key: bytes, cipher_text: bytes) -> bytes:
    iv = cipher_text[:16]
    aesgcm = AESGCM(key)
    plain_text = aesgcm.decrypt(iv, cipher_text[16:], None)
    return plain_text

if __name__ == "__main__":
    # read ticket from stdin
    ticket = sys.stdin.read()

    # split ticket into host and secret
    host, secret = ticket.split(" ")

    # decrypt secret
    key: bytes = ("A" * 32).encode()
    cipher_text = b64decode(secret)
    decrypted_data = aes_decrypt(key, cipher_text)

    # print decrypted secret
    print(decrypted_data)
