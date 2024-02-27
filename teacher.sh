export ID_INTERVIEW="$(grep -R "L337" interviews | cut -d ':' -f 1 | rev | cut -d '-' -f 1 | rev)"

echo $ID_INTERVIEW
cat "interviews/interview-"$ID_INTERVIEW
echo $MAIN_SUSPECT