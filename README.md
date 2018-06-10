Prerequisites:<br />
    Machine running Docker using Linux Containers

Instructions:<br />
    cd <install_directory><br />
    docker-compose up --build
    
Web site is running on port 80 of the container host.

API:

    POST /truetosize
    JSON: { size: int }


    GET /truetosize
    returns JSON: { averageTrueToSize: float32 }
