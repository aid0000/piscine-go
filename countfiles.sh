
( find . -type d && find . -type f ) | wc -l | sed ':1d'
