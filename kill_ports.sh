#!/bin/bash

# Specify the port number
PORT_NUMBER="${1:-50051}"

# Find PIDs using the specified port
PIDS=$(lsof -ti :$PORT_NUMBER)

# Check if there are any processes using the port
if [ -z "$PIDS" ]; then
    echo "No processes found using port $PORT_NUMBER"
else
    echo "Processes using port $PORT_NUMBER: $PIDS"
    
    # Kill each process
    for PID in $PIDS; do
        echo "Killing process $PID"
        kill -9 $PID
    done
    
    echo "Processes killed."
fi
