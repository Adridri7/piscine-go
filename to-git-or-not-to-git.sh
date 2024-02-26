

curl --silent "https://zone01normandie.org/assets/superhero/all.json" | jq  -r '.[] | select(.id==170) | .name ' 
curl --silent "https://zone01normandie.org/assets/superhero/all.json" | jq -r '.[] | select(.id==170) | .powerstats.power '
curl --silent "https://zone01normandie.org/assets/superhero/all.json" | jq  -r '.[] | select(.id==170) | .appearance.gender '


