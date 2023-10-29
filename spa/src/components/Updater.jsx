
import React, { useRef, useState, useContext, useEffect } from "react"
import { useMapContext } from "../contexts/MapContext";

import api from '../api';

export const Updater = () => {

    const [intervalTrucksId, setTrucksInterval] = useState(null);
    const [intervalPointsId, setPointsInterval] = useState(null);
    const [intervalPositionsId, setPositionsInterval] = useState(null);

    const { points, setPoints, setTrucks, setPositions } = useMapContext();

    useEffect(() => {
      // const newintervalPointsId = setInterval(getPoints, 2000);
      // setTrucksInterval( newintervalPointsId );
      //
      // const newintervalTrucksId = setInterval(getTrucks, 2000);
      // setPointsInterval( newintervalTrucksId );
      //
      // const newintervalPositionsId = setInterval(getPositions, 1000);
      // setPositionsInterval( newintervalPositionsId );
      //
      // return () => {
      //   clearInterval(intervalPointsId);
      //   clearInterval(intervalTrucksId);
      //   clearInterval(intervalPositionsId);
      // };
   	}, []);

    const getPoints = async () => {
        await api.getPoints( ).then( points => {
            setPoints(points);
          }
        ).catch( err => {
          console.log( "getPoints exception " + JSON.stringify(err) );
          }
        );
    }

    const getTrucks = async () => {
      await api.getTrucks( ).then( trucks => {
          setTrucks(trucks);
        }
      ).catch( err => {
        console.log( "getTrucks exception " + JSON.stringify(err) );
        }
      );
    }

    const getPositions = async () => {
      await api.getPositions( ).then( positions => {
          setPositions(positions);
        }
      ).catch( err => {
        console.log( "getPositions exception " + JSON.stringify(err) );
        }
      );
    }


    return (
			null
  );
}
