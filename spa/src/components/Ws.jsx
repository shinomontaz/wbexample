
import React, { useRef, useState, useContext, useEffect } from "react"
import { useMapContext } from "../contexts/MapContext";

import io from 'socket.io-client';

//import { socket } from '../socket';

import {SocketContext} from '../socket';

export default function Ws() {
    const socket = useContext(SocketContext);

    const [isConnected, setIsConnected] = useState(socket.connected);

    // const [points, setPoints] = useState([]);
    // const [trucks, setTrucks] = useState([]);
    // const [positions, setPositions] = useState([]);

    const { points, setPoints, setTrucks, setPositions } = useMapContext();

    useEffect(() => {
      socket.on('connect', () => {
        setIsConnected(true);
        console.log("connected ws");
      });
      socket.on('message', (data) => {
        console.log("message come: " + JSON.stringify(data));
//        parseMessage(data);
      });

      socket.on("disconnect", () => {
          console.log("Client disconnected");
      });
      return () => {
        socket.disconnect();
      };
   	}, []);

    // const parseMessage = (data) => {
    //   // get type
    //   switch (data.type) {
    //     case 'point':
    //       if (data.state == 'new') {
    //         setPoints([...points, data]);
    //       }
    //
    //       // if data.state == 'delete' {
    //       //   setPoints(points.filter(...))
    //       // }
    //
    //     break;
    //     case 'truck':
    //       setTrucks(data.trucks);
    //     break;
    //     case 'position':
    //       setPositions(data.positions);
    //     break;
    //   }
    // }

    return (
			<>
      isConnected: {JSON.stringify( isConnected )}
			</>
  );
}

//export default Map;
