interview_ID= grep -r "reports seeing the car that fled the scene."  | cut -d ":" -f1 | cut  -d "-" -f5 | sed 's/teacher.sh//' | tail -n 1 #Etape 1
echo $interview_ID #Etape 2

grep -r "reports seeing the car that fled the scene."  | cut -d ":" -f2

echo $MAIN_SUSPECT