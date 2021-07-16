# uuid-base64-converter

Installer
1.  wget https://github.com/BlackMocca/uuid-base64-converter/releases/download/v1.0.0/ub64c-darwin-amd64-1.0.0
2.  mv ub64c-darwin-amd64-1.0.0 ub64c
3.  chmod +x ub64c
4.  mv ub64c /usr/local/bin

Example Usage
```command
ub64c 64f9aec7-e890-49e5-81f6-d17ef89a4fa9
```

Will Be Show
```command
uid             :  64f9aec7-e890-49e5-81f6-d17ef89a4fa9
base64          :  ZPmux+iQSeWB9tF++JpPqQ==
mongoBinaryData :  {"$binary":{"base64":"ZPmux+iQSeWB9tF++JpPqQ==","subType":"4"}}
```

MongoBinaryData is json that can be input on database mongodb with type UUID("64f9aec7-e890-49e5-81f6-d17ef89a4fa9")