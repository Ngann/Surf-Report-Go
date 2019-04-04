package main

var testRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

var testString = `#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
2019 04 04 16 00  1.1  0.8 14.8  0.6  5.0 WSW WSW      SWELL  5.3 241
2019 04 04 15 00  1.0  0.8 14.8  0.6  5.3 WSW WSW      SWELL  5.4 243
2019 04 04 14 00  1.1  1.0 14.8  0.5  4.2   E WSW      SWELL  6.1  80
2019 04 04 13 00  1.2  1.1 14.8  0.4  3.8 SSW WSW      SWELL  5.8 211
2019 04 04 12 00  1.3  1.2 14.8  0.4  3.8  SW WSW      SWELL  5.8 218
2019 04 04 11 00  1.2  1.1  6.2  0.4  4.0 WNW WSW      STEEP  5.6 282
2019 04 04 10 00  1.3  1.2 13.8  0.4  3.4  SW  SW      SWELL  6.0 236
2019 04 04 09 00  1.3  1.3  7.1  0.4  4.0 WSW WSW      STEEP  5.9 250
2019 04 04 08 00  1.4  1.3  5.6  0.4  3.8 WSW   W VERY_STEEP  5.9 256
2019 04 04 07 00  1.4  1.4  6.2  0.4  3.4   W  SW      STEEP  5.8 267`
