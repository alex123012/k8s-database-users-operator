package postgresql_test

const (
	sslCAKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA6PqSDsm1ohRJ9fFF+wbLSIC/NnmY7R5Cg2DgSUP9HK4cWeN9
n9QSWoXwTZS5TKjuPBGYEnh91xbiEo2DcFGD+DF6x6Lf1N7Lh2LIV+MVlcD1KkId
05deDWkKxBSwfd7qHo4HSnPu97fDsQ201uxvMcrn80nZBMsIsmjjwYAvE6GUZf9k
3AmHFWVvRKmWYpM94AAXVpjjgm5rfmUCl265vNsn7EV9V0ejDrs+8ehp9UmYt8ix
Fx0T/fSQHhe14w212cbdH0oeKe/1DGkTaLXpPPoe3iQMWNS7axQ1KKYoWChCynn3
Q9/Sw2ui/CrRUsVsnivuReIgiXBLXYlTGQjqLwIDAQABAoIBAD2inKVg47Z/W5m/
PN5OS9VgijWlMbn5eWs8Y9m+LOY7gbCeKIvyFPDx4kMEB4mqX2xw0yR/z/rpSOHT
omRCjIFKxcqu4jx0vK+SiKIHp5w4siN93lot/2nY0kpRlueV46Y1uOQPi7bpXNIo
aCE756bqoTaR4OINvL1GVorImKu+YDXMllrkqOg9WPGGQNfQFekYCc9rEPxjM87j
dE/Tkfi8kRZEQ2BxkdAZimp4KO87sO3FP80ZxTAZOQleYHosQPp9/Fd9zCXCaKWq
tcfdZjyThFZyydSGUg+8K9fSzZmkyG0eUz/W75/eqzSmTS6ugT0Tq1bvjiHpNm+v
24v3moECgYEA+/wFR5y4zykP49RYoCcvJMOOHRScjabJLsjpAPeLFeKdQj9kzxM7
j1JonFtgIh/yFhfdNE49XEE1vi6wCYFl6HFzze3bnuK0jJ+u9+OmfcLTR+036Gbv
zbRt0poRzWhCVK6J/stbjDkk0a2F7z+rLITqOJwXAAmWtoegkQwGBA8CgYEA7LEE
ATN1ZfZ1mOvar8z+A+QK7qkIwwLFRl0y9aziZxqNIVCLm8x+YssvqdKl4QRf7I7/
Cv+ZqsZ8AVo5ueX4jSNbQasf84vh3YqmoCH/Wo5468K4VvvYA9xMQ5vq8nsojMcs
eSH9VDr7nfa1oR9cwRgXcf/61R2rr5OyjSJ+F+ECgYEA+eKX3cdmYlGnJ1kqNlAF
aWDgaqhJBBQ1CEdHAaV6cU6UguDY+J1rABtKEFxxPgYODajGvZslMHqecCZefl6r
D9KKc9oAZFUPlTC506wXLDnrSjXNrpN+FbFrA2G0a82LkeywflNuSuVURPbejj7G
YlTA7Tilem0H36UqLw0MXjMCgYARdGhoMkRJFajMczA4YLSm0s1flkWYI/8qVjso
1OwJUHLx7v+sqKL1ZCiKrIchFfKA/naeeAT8DBEfBGlXZTc0KVRUfmsnybwJW204
R+mN4w3VzRFNENt4RWm2Xqwv35c48oM8F56X9JWTq5rvW+G6N62a8Zas2rhLhWfY
cp+74QKBgGZZDdrEnNJKeAia4AAyDoTlZaPjxOtOTxsL7zwjykMGkKtGCpbKjhej
Og4n+OuzlhY1cVv3u/Jj++XtaJfM0sM8LNxyYzeb/B4N+vNmKMogOOjVIJhHsZTR
fQKcnTU/S4jGtSfPLUePvdaK5+ezCyepCXO18mZ7TSemGZVyroBy
-----END RSA PRIVATE KEY-----
`

	sslCACert = `
-----BEGIN CERTIFICATE-----
MIIDrjCCApYCCQDu7SVJupOI5DANBgkqhkiG9w0BAQsFADCBlzELMAkGA1UEBhMC
R0UxDDAKBgNVBAgMA1VsbTEMMAoGA1UEBwwDVWxtMQ4wDAYDVQQKDAVGbGFudDEQ
MA4GA1UECwwHRm94dHJvdDEgMB4GA1UEAwwXZGF0YWJhc2UtdXNlcnMtb3BlcmF0
b3IxKDAmBgkqhkiG9w0BCQEWGWFsZXhleS5tYWtob25pbkBmbGFudC5jb20wIBcN
MjMwNjI1MTkwNjI3WhgPMzAyMjEwMjYxOTA2MjdaMIGXMQswCQYDVQQGEwJHRTEM
MAoGA1UECAwDVWxtMQwwCgYDVQQHDANVbG0xDjAMBgNVBAoMBUZsYW50MRAwDgYD
VQQLDAdGb3h0cm90MSAwHgYDVQQDDBdkYXRhYmFzZS11c2Vycy1vcGVyYXRvcjEo
MCYGCSqGSIb3DQEJARYZYWxleGV5Lm1ha2hvbmluQGZsYW50LmNvbTCCASIwDQYJ
KoZIhvcNAQEBBQADggEPADCCAQoCggEBAOj6kg7JtaIUSfXxRfsGy0iAvzZ5mO0e
QoNg4ElD/RyuHFnjfZ/UElqF8E2UuUyo7jwRmBJ4fdcW4hKNg3BRg/gxesei39Te
y4diyFfjFZXA9SpCHdOXXg1pCsQUsH3e6h6OB0pz7ve3w7ENtNbsbzHK5/NJ2QTL
CLJo48GALxOhlGX/ZNwJhxVlb0SplmKTPeAAF1aY44Jua35lApduubzbJ+xFfVdH
ow67PvHoafVJmLfIsRcdE/30kB4XteMNtdnG3R9KHinv9QxpE2i16Tz6Ht4kDFjU
u2sUNSimKFgoQsp590Pf0sNrovwq0VLFbJ4r7kXiIIlwS12JUxkI6i8CAwEAATAN
BgkqhkiG9w0BAQsFAAOCAQEAWkZlg9Iwlogeaa2fNnWcn/8+l5P0r6ntmsVlAuIm
rwHVOk3gYb6Y5f3vjerL+Hl8WMwOtNUSBvJPKfqDjSEPLBsexTI2aL1Bg2uYO7RQ
gGDPqMfCwh6fnRIohxS8x77nnol4Q1S/yrJiPAySCr5y5iZ/Cs2ddlj8Vwg2cAIW
Wjjoai1uC9JtSKYRKHJ0UnnVnoDqcfh3L/aFdp4GiWTxt7Wf5305dwMhUb+QQISS
LCY4jRRLiz7w1Ps6j/UWXPS/kEKkBI1vzRs1MIR6u7dh97p2hhBheZo91tjgAMti
cvd2vcFJnLalUzjlWlMdNhGbcamwtehiUumwwFhX3gGz0Q==
-----END CERTIFICATE-----
`

	sslJohnCert = `
-----BEGIN CERTIFICATE-----
MIIDbTCCAlWgAwIBAgIRAJP0yfwGwnS6+IPOEyG0MlQwDQYJKoZIhvcNAQELBQAw
gZcxCzAJBgNVBAYTAkdFMQwwCgYDVQQIDANVbG0xDDAKBgNVBAcMA1VsbTEOMAwG
A1UECgwFRmxhbnQxEDAOBgNVBAsMB0ZveHRyb3QxIDAeBgNVBAMMF2RhdGFiYXNl
LXVzZXJzLW9wZXJhdG9yMSgwJgYJKoZIhvcNAQkBFhlhbGV4ZXkubWFraG9uaW5A
ZmxhbnQuY29tMB4XDTIzMDYyNDE5MDg0OVoXDTI4MDYyODE5MDg0OVowIzESMBAG
A1UEChMJQ29ja3JvYWNoMQ0wCwYDVQQDEwRqb2huMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAnJcQWasALHeSiY5PdrKiF34tuS+8svhHd/BuwQ5/Et/n
YhNbearEtKbC9aVTDYl5NX1Wy/NoTTVOvhPvR7NX7KJ2dYuvtY6EEpZB2XK18Kua
Ejq6u21kKstKEPW5JGKTSWhP9VX2HA1cB4dfn3e6VXguNpSve12FYBnTDLne3MF8
Yr9qv+kHNUpEXMXWfIMP2/VVzPFtOjbPOS0pBCXLKeXzl+F8eBHmV/XEFdPmKWiW
7hrQRQIopsufVH4QvJdYyPNyrZos3c9N/4lI8TbxUXb6pH0SPVvF7Ng3YisNdJ9/
QQVli3eHTOr+1UEyUMJBvaJr/J6kQZLZcCyiH2qT5wIDAQABoycwJTAOBgNVHQ8B
Af8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggEB
AKMYhEZrnHfGZxpN7Eaul0+D3tneCnsSvfLCCdZzcDgIvv1GiRYdPnj2XT+ZLkNr
TXIPjyhf4HLwOPDWQzZ8fu25GsVY5AgCJ8cMqd895vCTaGI56BudGLDRJ/tJDup+
voFDouy2UxtlvnhY2ShrwcXLR+lRM3QsI1RcHHoN77VA6u6R2K2OWhEHm4YUNEyR
6IlgoQvE4WUWeDnlJibo7IKf06LZHnB5cj/xTMo6TH+xerzksS9Ul+0VJT3L87WC
yzQ9S6Fc31Ikp5DAhldppbkZ6Nw6wAEeE5kdVdAFFgwkmjodi39I3Z8NBsrqsoPa
QqO8Dwo4cNBvpWIyKCAcTvQ=
-----END CERTIFICATE-----
`

	sslJohnKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAnJcQWasALHeSiY5PdrKiF34tuS+8svhHd/BuwQ5/Et/nYhNb
earEtKbC9aVTDYl5NX1Wy/NoTTVOvhPvR7NX7KJ2dYuvtY6EEpZB2XK18KuaEjq6
u21kKstKEPW5JGKTSWhP9VX2HA1cB4dfn3e6VXguNpSve12FYBnTDLne3MF8Yr9q
v+kHNUpEXMXWfIMP2/VVzPFtOjbPOS0pBCXLKeXzl+F8eBHmV/XEFdPmKWiW7hrQ
RQIopsufVH4QvJdYyPNyrZos3c9N/4lI8TbxUXb6pH0SPVvF7Ng3YisNdJ9/QQVl
i3eHTOr+1UEyUMJBvaJr/J6kQZLZcCyiH2qT5wIDAQABAoIBAQCU077menAf00Wj
F27PEdidG3+5knV2ZCMJC6s4Md70wXnY7Szz5iouyJBjiE33f8GD5Sypix6GwzOj
1K6HJx1Z+s87yRenJ1y/ja/oS+5AX9h/mvH/UWjyg2RR2jmtK0NxcYMNWjYfU1M0
lKV9mv2uXhsOJSLjzW8Gd4Tvg3kdtgsvcaqUXiVvyxlY4FW3aL+alOlDIxyOK3Wf
XUNFQbgAjuu3uLkx6lYVdF41JgEq9l9vdQtJd7l9K/VEn0e0+SkjHC8UqBbER++1
orEiD3dmwyqOyQZ2YuGAfe6zzNK9XrtuaDKLlEfkCD5ByBphd6GG/A4T7WZqeO5+
b8IeY8UBAoGBAMURKPyEg4Qg1Vf9fivpFNDtwJfPAdZ3d/7U308BDEhE8An2gneH
wHzOQhSo2NtuhTHWEtyilPgZqVNDnQE9CZjPNwGezhCcurPtM3dAgLs0q1rYHmKL
pRiUkUn/TbjMkv1VatcmvXgQFKmXsRm7Ssbr2lipEZrQaUottR0Rw2H5AoGBAMtr
HXVf/QveAYtcenAmaI/M6dOoyvZefuFx1q90pSEJGYDqJe/JTvJAeBeuMkDpWVLr
oBLxILMg8BzqEAcQSSvFG2OCCyBYoKb4apI2mECK5fz4i5QDFbt4G1UUwt5hOGX8
B1FagY7oVhfpbk/IOXToJkmT4cH+z135z2/zyRzfAoGAZmJI+hDax03LqcBgye15
zCaJ1hVNriA5rqLoNgKkX/O5BmQVWoakAfOjL1qd+DtOZhsDh6/MV631Y/YP3zHY
B5U4zdW017ql4Y8OGxnfB+QQVs1L1AUbTE77wQcsWSOoBohXTtqou3UXVxkhgO3m
pryon0GPjPBUk551p1mwOfkCgYAo1aPgQBioHTTqKPJbORqcY2I9HxE/S9DkqNmT
9zJ+4zi/bEGZVSwH7XEuL8XeyfkocCx+IPGTg/UvmL0G7foCU1sgKqbZI3F8kzmx
iEwgCMIKekpquAPQ0leKSNSll5aewm0lo6mGapV9z1pZobQHB+NHuewD6YbvMoq1
ypaAuQKBgQCNh068Jes6p1zFV+7zy4ztIGwpSS9rzQY0sS3zTXsl8jPF0SS4ysaX
gVLGNgSsLTRLLcM5V+2cWS588yi4W4GoDxFaPykiq5nasQn2OTL5i2XmAjvcaLS0
4TzUo9PEr56k4G2qWjPtMaEWOU+XPT554DVMvAoVRXuJsLfrSPmzkg==
-----END RSA PRIVATE KEY-----
`
)
