jq  -r '.[] | select(.id==170) | .name ' all.json
jq -r '.[] | select(.id==170) | .powerstats.power ' all.json
jq  -r '.[] | select(.id==170) | .appearance.gender ' all.json


