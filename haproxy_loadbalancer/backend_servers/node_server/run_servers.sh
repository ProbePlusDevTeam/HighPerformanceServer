counter=0
while [ $counter -le 0 ]
do   
  port=$((8282+$counter))   
  nodejs /home/proebeplus23/node_server/new.js $port & echo "Server created on port "  $port
((counter++))
done
echo "Created all servers"

