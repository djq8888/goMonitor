sed -i "s/127.0.0.1/$1/g" js/*.js
sed -i "s/8080/$2/g" js/*.js
sed -i "s/8080/$2/g" main.go