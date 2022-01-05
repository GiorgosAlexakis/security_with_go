# Data Integrity
## Definition
Ensuring data integrity requires a way to test that incoming data has not been tampered with.

## Hashing
Hashing is used as a means of testing of data tampering. This is accomplished by forming a simple number from the data's bytes. The resulting number is called a hash or hash value.

One very simple method to achieve this is to sum all the bytes in the data. However, caution is required as someone can change two bytes and still produce the same hash.

Strong security can be achieved by making it difficult to find a different sequence of bytes that produces the same hash value. 

## Go Support 
Go has support for several hashing algorithms, including MD4, MD5,RIPEMD-160, SHA1, SHA224,SHA256, SHA384, and SHA512. These can be found under the crypto package.