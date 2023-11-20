counter=0
while [ $counter -le 1 ]
do   
  port=$((8282+$counter))   
  ./go-server -port $port & echo "Server created on port "  $port
((counter++))
done
echo "Created all servers"

